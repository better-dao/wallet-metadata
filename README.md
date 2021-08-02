# wallet-metadata
crypto wallet metadata


## feature:

- generate crypto wallet token contract metadata:
    - eth
    - bsc
- source: `data`    
- dist: `dist`


## metadata:

- [dist](dist)

### ETH Tokens: 

- ETH Tokens: [dist/eth_contract_map.json](./dist/eth_contract_map.json)
- token metadata: https://github.com/MetaMask/contract-metadata/blob/master/contract-map.json
    - this file miss some fields, not completed


### BSC Tokens:

- BSC Tokens: [dist/bsc_contract_map.json](dist/bsc_contract_map.json)
    - support bsc `mainnet/testnet`
- token icon: https://github.com/binance-chain/tokens-info
- token metadata: 
    - https://dex.binance.org/api/v1/tokens?limit=1000000000
    - https://testnet-dex.binance.org/api/v1/tokens?limit=1000000000



## quickstart:


- install:

```bash
make install
```

- init: (auto gen `.env` file)

```bash
make init 

# or:

task init

```

- see `.env` file: [.env](.env.local) (make init will copy `.env.local` to `.env`)
    - change `env vars` to gen `eth` or `bsc` metadata

- run:


```bash
make gen 

# or

task gen

```

## ref: 

- https://taskfile.dev/#/usage?id=env-files
- https://github.com/better-go/pkg
- https://gobyexample.com/writing-files
