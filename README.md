# EVMStoreChain
EVMStoreChain is a Cosmos SDK-based blockchain that allows users to read state from Ethereum, post the Ethereum state to the Cosmos blockchain, and query the agreed-upon state. This application interacts with the Compound protocol to fetch the borrowIndex of the cUSDT token of the latest block and post it as a vote to the Cosmos blockchain.

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





