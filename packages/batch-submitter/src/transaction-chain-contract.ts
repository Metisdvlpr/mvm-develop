/* External Imports */
import { Contract, ethers, BigNumber } from 'ethers'
import {
  TransactionResponse,
  TransactionRequest,
} from '@ethersproject/abstract-provider'
import { keccak256 } from 'ethers/lib/utils'
import {
  //AppendSequencerBatchParams,
  BatchContext,
  encodeAppendSequencerBatch,
  remove0x,
  EncodeSequencerBatchOptions
} from '@metis.io/core-utils'

interface AppendSequencerBatchParams {
    chainId: number;
    shouldStartAtElement: number;
    totalElementsToAppend: number;
    contexts: BatchContext[];
    transactions: string[];
    blockNumbers: number[];
}

export { encodeAppendSequencerBatch, BatchContext, AppendSequencerBatchParams }

/*
 * OVM_CanonicalTransactionChainContract is a wrapper around a normal Ethers contract
 * where the `appendSequencerBatchByChainId(...)` function uses a specialized encoding for improved efficiency.
 */
export class CanonicalTransactionChainContract extends Contract {
  public customPopulateTransaction = {
    appendSequencerBatch: async (
      batch: AppendSequencerBatchParams,
      opts?: EncodeSequencerBatchOptions
    ): Promise<ethers.PopulatedTransaction> => {
      const nonce = await this.signer.getTransactionCount()
      const to = this.address
      const data = await getEncodedCalldata(batch, opts)
      const gasLimit = await this.signer.provider.estimateGas({
        to,
        from: await this.signer.getAddress(),
        data,
      })

      return {
        nonce,
        to,
        data,
        gasLimit,
      }
    },
  }
  public async appendSequencerBatch(
    batch: AppendSequencerBatchParams,
    options?: TransactionRequest,
    opts?: EncodeSequencerBatchOptions
  ): Promise<TransactionResponse> {
    return appendSequencerBatch(this, batch, options, opts)
  }
}

/**********************
 * Internal Functions *
 *********************/

const APPEND_SEQUENCER_BATCH_METHOD_ID = keccak256(
  Buffer.from('appendSequencerBatchByChainId()')
).slice(2, 10)

const appendSequencerBatch = async (
  CanonicalTransactionChain: Contract,
  batch: AppendSequencerBatchParams,
  options?: TransactionRequest,
  opts?: EncodeSequencerBatchOptions
): Promise<TransactionResponse> => {
  return CanonicalTransactionChain.signer.sendTransaction({
    to: CanonicalTransactionChain.address,
    data: await getEncodedCalldata(batch, opts),
    ...options,
  })
}
const encodeHex = (val: any, len: number) =>
  remove0x(BigNumber.from(val).toHexString()).padStart(len, '0')
const getEncodedCalldata = async (batch: AppendSequencerBatchParams, opts?: EncodeSequencerBatchOptions): Promise<string> => {
  const methodId = APPEND_SEQUENCER_BATCH_METHOD_ID
  const calldata = await encodeAppendSequencerBatch(batch, opts)
  return '0x' + remove0x(methodId) + encodeHex(batch.chainId, 64) + remove0x(calldata)
}
