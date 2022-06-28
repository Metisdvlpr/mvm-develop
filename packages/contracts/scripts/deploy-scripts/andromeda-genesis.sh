#!/bin/bash
URL=https://metis-us-east-2-mainnet-json.s3.us-east-2.amazonaws.com/addresses.json
ADDRESSES=$(curl --fail --show-error --silent --retry-connrefused --retry 2 --retry-delay 5 $URL)

function envSet() {
    VAR=$1
    export $VAR=$(echo $ADDRESSES | jq -r ".$2")
}

# set the address to the proxy gateway if possible
envSet L1_STANDARD_BRIDGE_ADDRESS Proxy__OVM_L1StandardBridge

envSet L1_CROSS_DOMAIN_MESSENGER_ADDRESS Proxy__OVM_L1CrossDomainMessenger

envSet L1_METIS_MANAGER_ADDRESS Proxy__MVM_ChainManager

export L2_BLOCK_GAS_LIMIT=1100000000
export L2_CHAIN_ID=1088
export BLOCK_SIGNER_ADDRESS=0x00000398232E2064F896018496b4b44b3D62751F
export L1_FEE_WALLET_ADDRESS=0xDD6FFC7D9a4Fb420b637747edc6456340d12d377
export WHITELIST_OWNER=0xDD6FFC7D9a4Fb420b637747edc6456340d12d377
export GAS_PRICE_ORACLE_OWNER=0xDD6FFC7D9a4Fb420b637747edc6456340d12d377

export GAS_PRICE_ORACLE_OVERHEAD=2750
export GAS_PRICE_ORACLE_SCALAR=40000000
export GAS_PRICE_ORACLE_L1_BASE_FEE=150000000000
export GAS_PRICE_ORACLE_GAS_PRICE=40000000000
export GAS_PRICE_ORACLE_DECIMALS=6

export METIS_ADDRESS=0x9E32b13ce7f2E80A01932B42553652E053D6ed8e
#export MIN_L1_ERC20_BRIDGE_COST=400000000000000000
export MIN_L1_ERC20_BRIDGE_COST=1000000

yarn build:dump

cd ./dist/dumps
exec python -c \
            'import BaseHTTPServer as bhs, SimpleHTTPServer as shs; bhs.HTTPServer(("0.0.0.0", 8081), shs.SimpleHTTPRequestHandler).serve_forever()'
