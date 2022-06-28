/* Imports: External */
import { BigNumber, ethers, constants } from 'ethers'
import { MerkleTree } from 'merkletreejs'
import { Logger } from '@eth-optimism/common-ts'
import { getContractFactory } from '@metis.io/contracts'
import {
  fromHexString,
  toHexString,
  toRpcHexString,
  EventArgsSequencerBatchAppended,
  MinioClient,
  MinioConfig,
} from '@metis.io/core-utils'

/* Imports: Internal */
import {
  DecodedSequencerBatchTransaction,
  SequencerBatchAppendedExtraData,
  SequencerBatchAppendedParsedEvent,
  TransactionBatchEntry,
  TransactionEntry,
  EventHandlerSet,
} from '../../../types'
import {
  SEQUENCER_ENTRYPOINT_ADDRESS,
  SEQUENCER_GAS_LIMIT,
  parseSignatureVParam,
} from '../../../utils'
import { MissingElementError } from './errors'

export const handleEventsSequencerBatchAppended: EventHandlerSet<
  EventArgsSequencerBatchAppended,
  SequencerBatchAppendedExtraData,
  SequencerBatchAppendedParsedEvent
> = {
  getExtraData: async (event, l1RpcProvider) => {
    const l1Transaction = await event.getTransaction()
    const eventBlock = await event.getBlock()

    // TODO: We need to update our events so that we actually have enough information to parse this
    // batch without having to pull out this extra event. For the meantime, we need to find this
    // "TransactonBatchAppended" event to get the rest of the data.
    const CanonicalTransactionChain = getContractFactory(
      'CanonicalTransactionChain'
    )
      .attach(event.address)
      .connect(l1RpcProvider)

    const batchSubmissionEvent = (
      await CanonicalTransactionChain.queryFilter(
        CanonicalTransactionChain.filters.TransactionBatchAppended(),
        eventBlock.number,
        eventBlock.number
      )
    ).find((foundEvent: ethers.Event) => {
      // We might have more than one event in this block, so we specifically want to find a
      // "TransactonBatchAppended" event emitted immediately before the event in question.
      return (
        foundEvent.transactionHash === event.transactionHash &&
        foundEvent.logIndex === event.logIndex - 1
      )
    })

    if (!batchSubmissionEvent) {
      throw new Error(
        `Well, this really shouldn't happen. A SequencerBatchAppended event doesn't have a corresponding TransactionBatchAppended event.`
      )
    }

    return {
      timestamp: eventBlock.timestamp,
      blockNumber: eventBlock.number,
      submitter: l1Transaction.from,
      l1TransactionHash: l1Transaction.hash,
      l1TransactionData: l1Transaction.data,
      gasLimit: `${SEQUENCER_GAS_LIMIT}`,

      prevTotalElements: batchSubmissionEvent.args._prevTotalElements,
      batchIndex: batchSubmissionEvent.args._batchIndex,
      batchSize: batchSubmissionEvent.args._batchSize,
      batchRoot: batchSubmissionEvent.args._batchRoot,
      batchExtraData: batchSubmissionEvent.args._extraData,
    }
  },
  parseEvent: async (event, extraData, l2ChainId, options) => {
    const transactionEntries: TransactionEntry[] = []

    // It's easier to deal with this data if it's a Buffer.
    let calldata = fromHexString(extraData.l1TransactionData)
    let minioClient: MinioClient = null
    if (options.minioBucket && options.minioAccessKey &&
      options.minioSecretKey && options.minioEndpoint &&
      options.minioPort) {
        const minioConfig: MinioConfig = {
          options: {
            endPoint: options.minioEndpoint,
            port: options.minioPort,
            useSSL: options.minioUseSsl,
            accessKey: options.minioAccessKey,
            secretKey: options.minioSecretKey,
          },
          l2ChainId: l2ChainId,
          bucket: options.minioBucket
        }
        minioClient = new MinioClient(minioConfig)
      }

    // chainid + 32, so not [12, 15]
    if (calldata.length < 44) {
      throw new Error(
        `Block ${extraData.blockNumber} transaction data is invalid for decoding: ${extraData.l1TransactionData} , ` +
          `converted buffer length is < 44.`
      )
    }
    const numContexts = BigNumber.from(calldata.slice(44, 47)).toNumber()
    let transactionIndex = 0
    let enqueuedCount = 0
    let nextTxPointer = 47 + 16 * numContexts
    const leafs = []
    let rootFromCalldata = ''
    let fromStorage = false
    for (let i = 0; i < numContexts; i++) {
      const contextPointer = 47 + 16 * i
      const context = parseSequencerBatchContext(calldata, contextPointer)

      if (
        i === 0 &&
        context.blockNumber === 0 &&
        context.timestamp === 0 &&
        context.numSubsequentQueueTransactions === 0 &&
        context.numSequencedTransactions === 0
      ) {
        // calldata = timestamp[13] + zero[1]{0} + sizeOfTxData[8]{00000000} + markleRoot[64]
        const storageObject = calldata.slice(nextTxPointer).toString('hex').slice(0, 86)
        rootFromCalldata = storageObject.slice(22, 86)
        fromStorage = true
        // console.info('calc storage object name', storageObject)
        const txData = await minioClient.readObject(storageObject, 2)
        // const verified = await minioClient.verifyObject(storageObject, txData, 2)
        // if (!verified) {
        //   throw new Error(`verified calldata from storage error, storage object ${storageObject}`)
        // }
        if (!txData) {
          throw new Error(`got calldata from storage error, storage object ${storageObject}`)
        }
        console.info('got storage data', storageObject)
        calldata = Buffer.concat([
          calldata.slice(0, nextTxPointer),
          Buffer.from(txData, 'hex'),
        ])
      }

      for (let j = 0; j < context.numSequencedTransactions; j++) {
        const sequencerTransaction = parseSequencerBatchTransaction(
          calldata,
          nextTxPointer
        )

        const decoded = decodeSequencerBatchTransaction(
          sequencerTransaction,
          event.args._chainId
        )

        transactionEntries.push({
          index: extraData.prevTotalElements
            .add(BigNumber.from(transactionIndex))
            .toNumber(),
          batchIndex: extraData.batchIndex.toNumber(),
          blockNumber: BigNumber.from(context.blockNumber).toNumber(),
          timestamp: BigNumber.from(context.timestamp).toNumber(),
          gasLimit: BigNumber.from(0).toString(),
          target: constants.AddressZero,
          origin: null,
          data: toHexString(sequencerTransaction),
          queueOrigin: 'sequencer',
          value: decoded.value,
          queueIndex: null,
          decoded,
          confirmed: true,
        })
        // block number = index + 1
        leafs.push(ethers.utils.keccak256(
          ethers.utils.solidityPack(['uint256', 'bytes'],
          [
            extraData.prevTotalElements.add(BigNumber.from(transactionIndex)).add(BigNumber.from(1)).toNumber(),
            parseMerkleLeafFromSequencerBatchTransaction(calldata, nextTxPointer)
          ])))
        nextTxPointer += 3 + sequencerTransaction.length
        transactionIndex++
      }

      for (let j = 0; j < context.numSubsequentQueueTransactions; j++) {
        const queueIndex = event.args._startingQueueIndex.add(
          BigNumber.from(enqueuedCount)
        )

        // Okay, so. Since events are processed in parallel, we don't know if the Enqueue
        // event associated with this queue element has already been processed. So we'll ask
        // the api to fetch that data for itself later on and we use fake values for some
        // fields. The real TODO here is to make sure we fix this data structure to avoid ugly
        // "dummy" fields. EXCEPT timestamp, which is set to local time when enqueued. this timestamp
        // was submitted to CTC as part of the context. use the timestamp in the context otherwise
        // the batch timestamp will be inconsistent with the main node.
        transactionEntries.push({
          index: extraData.prevTotalElements
            .add(BigNumber.from(transactionIndex))
            .toNumber(),
          batchIndex: extraData.batchIndex.toNumber(),
          blockNumber: BigNumber.from(0).toNumber(),
          timestamp: extraData.prevTotalElements
            .add(BigNumber.from(transactionIndex))
            .toNumber()<=2287472 ? BigNumber.from(0).toNumber() : BigNumber.from(context.timestamp).toNumber(),  //timestamp needs to be consistent
          gasLimit: BigNumber.from(0).toString(),
          target: constants.AddressZero,
          origin: constants.AddressZero,
          data: '0x',
          queueOrigin: 'l1',
          value: '0x0',
          queueIndex: queueIndex.toNumber(),
          decoded: null,
          confirmed: true,
        })

        enqueuedCount++
        transactionIndex++
      }
    }

    const hash = (el: Buffer | string): Buffer => {
      return Buffer.from(ethers.utils.keccak256(el).slice(2), 'hex')
    }
    const tree = new MerkleTree(leafs, hash)
    let merkleRoot = tree.getHexRoot()
    if (merkleRoot.startsWith('0x')) {
      merkleRoot = merkleRoot.slice(2)
    }
    console.info(`root from batch: ${rootFromCalldata}, re-calculate root: ${merkleRoot}, equals: ${rootFromCalldata == merkleRoot}`)
    if (fromStorage && rootFromCalldata != merkleRoot) {
      throw new Error(`verified calldata from storage error, batch index is ${extraData.batchIndex.toNumber()}`)
    }

    const transactionBatchEntry: TransactionBatchEntry = {
      index: extraData.batchIndex.toNumber(),
      root: extraData.batchRoot,
      size: extraData.batchSize.toNumber(),
      prevTotalElements: extraData.prevTotalElements.toNumber(),
      extraData: extraData.batchExtraData,
      blockNumber: BigNumber.from(extraData.blockNumber).toNumber(),
      timestamp: BigNumber.from(extraData.timestamp).toNumber(),
      submitter: extraData.submitter,
      l1TransactionHash: extraData.l1TransactionHash,
    }

    return {
      transactionBatchEntry,
      transactionEntries,
    }
  },
  storeEvent: async (entry, db) => {
    // Defend against situations where we missed an event because the RPC provider
    // (infura/alchemy/whatever) is missing an event.
    if (entry.transactionBatchEntry.index > 0) {
      const prevTransactionBatchEntry = await db.getTransactionBatchByIndex(
        entry.transactionBatchEntry.index - 1
      )

      // We should *always* have a previous transaction batch here.
      if (prevTransactionBatchEntry === null) {
        throw new MissingElementError('SequencerBatchAppended')
      }
    }

    await db.putTransactionBatchEntries([entry.transactionBatchEntry])
    await db.putTransactionEntries(entry.transactionEntries)

    // Add an additional field to the enqueued transactions in the database
    // if they have already been confirmed
    for (const transactionEntry of entry.transactionEntries) {
      if (transactionEntry.queueOrigin === 'l1') {
        await db.putTransactionIndexByQueueIndex(
          transactionEntry.queueIndex,
          transactionEntry.index
        )
      }
    }
  },
}

interface SequencerBatchContext {
  numSequencedTransactions: number
  numSubsequentQueueTransactions: number
  timestamp: number
  blockNumber: number
}

const parseSequencerBatchContext = (
  calldata: Buffer,
  offset: number
): SequencerBatchContext => {
  return {
    numSequencedTransactions: BigNumber.from(
      calldata.slice(offset, offset + 3)
    ).toNumber(),
    numSubsequentQueueTransactions: BigNumber.from(
      calldata.slice(offset + 3, offset + 6)
    ).toNumber(),
    timestamp: BigNumber.from(
      calldata.slice(offset + 6, offset + 11)
    ).toNumber(),
    blockNumber: BigNumber.from(
      calldata.slice(offset + 11, offset + 16)
    ).toNumber(),
  }
}

const parseMerkleLeafFromSequencerBatchTransaction = (
  calldata: Buffer,
  offset: number
): Buffer => {
  const transactionLength = BigNumber.from(
    calldata.slice(offset, offset + 3)
  ).toNumber()

  return calldata.slice(offset, offset + 3 + transactionLength)
}

const parseSequencerBatchTransaction = (
  calldata: Buffer,
  offset: number
): Buffer => {
  const transactionLength = BigNumber.from(
    calldata.slice(offset, offset + 3)
  ).toNumber()

  return calldata.slice(offset + 3, offset + 3 + transactionLength)
}

const decodeSequencerBatchTransaction = (
  transaction: Buffer,
  l2ChainId: number
): DecodedSequencerBatchTransaction => {
  const decodedTx = ethers.utils.parseTransaction(transaction)

  return {
    nonce: BigNumber.from(decodedTx.nonce).toString(),
    gasPrice: BigNumber.from(decodedTx.gasPrice).toString(),
    gasLimit: BigNumber.from(decodedTx.gasLimit).toString(),
    value: toRpcHexString(decodedTx.value),
    target: decodedTx.to ? toHexString(decodedTx.to) : null,
    data: toHexString(decodedTx.data),
    sig: {
      v: parseSignatureVParam(decodedTx.v, l2ChainId),
      r: toHexString(decodedTx.r),
      s: toHexString(decodedTx.s),
    },
  }
}
