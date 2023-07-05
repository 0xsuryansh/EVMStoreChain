# EVMStoreChain
EVMStoreChain is a Cosmos SDK-based blockchain that allows validators to read the state from Ethereum, post the Ethereum state to the Cosmos blockchain as a vote, and query the agreed-upon state. This application interacts with the Compound protocol to fetch the borrowIndex of the cUSDT token of the latest block and post it as a vote to the Cosmos blockchain.

### Problem Statement / Requirement
- Utilize the Cosmos SDK to build a prototype blockchain that can read the state from Ethereum, verify its validity, and store it on your chain.
- The chain should be able to agree upon some Ethereum state. There should be a mechanism that facilitates this agreement. The state itself should be a single storage slot of some address.
- Once some storage value is agreed upon, it should be stored on your blockchain along with the input parameters that were used to query it. Users should be able to read data from your blockchain.

### Approach / Solution
#### Validators report what they observe: 
Validators of the Cosmos blockchain should run a script or service that observes the Ethereum blockchain and submits `SubmitEthereumState` transactions to the Cosmos blockchain with what they observe. They submit the state observed and the block number the state observed for 
Each validator should be responsible for fetching the Ethereum state independently and then submitting a proposal with that state. Even if the Ethereum state changes between the time different validators fetch it, this will not introduce non-determinism into the Cosmos chain. This is because the Cosmos chain does not directly query the Ethereum state during transaction processing. Instead, it relies on the data provided in the validators' proposals, which is deterministic and consistent across all validators.

For validators to be honest and report the correct state , slashing module is used, the EVMStoreChain slash a validator when they submit a vote that does not match the majority decision.



## Prerequisites
- Go (version 1.19 or higher)
- Cosmos-sdk
- Infura account

## Installation
Install the Ignite CLI:
```
curl https://get.ignite.com/cli! | bash
```

## Clone this repository:
```
git clone https://github.com/0xsuryansh/EVMStoreChain
```

## Repository Structure
There are 2 directories in this repository:

1. EVMStateVoteClient - A Go client to query Ethereum for a state (borrowIndex of Compound's cUSDT token contract), post Ethereum state to the Cosmos blockchain, and query the state based on the block number.

2. EVMStoreChain - Cosmos SDK-based blockchain code.

## Usage

### Starting the Blockchain
Run the following command to start the blockchain:
```
cd EVMStoreChain
ignite chain serve
```

### Interacting with Ethereum and Cosmos Chains

Navigate to the EVMStateVoteClient directory:
```
cd EVMStateVoteClient
go mod tidy
```

### Set up your Infura URL for connecting to the Ethereum node:
```
export INFURA_URL="YOUR_INFURA_URL"
```

### Run the client:
```
go run main.go
```

This client fetches the borrowIndex of the Compound's cUSDT token of the latest block and posts it as a vote to the Cosmos blockchain using the SubmitEthereumState function. You can then query the state with the highest vote.

### Using CLI
```
For validator voting,
EVMStoreChaind tx evmstorechain submit-ethereum-state [blocknumber] [state] --from alice

For querying 
EVMStoreChaind q evmstorechain show-blockstoragestate [blocknumber]
```




