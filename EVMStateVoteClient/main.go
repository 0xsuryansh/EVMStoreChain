package main

import (
    "context"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum"
    "math/big"
    "io/ioutil"
    "log"
    "os"
    "fmt"
    "strconv"
    "strings"
    "github.com/ignite/cli/ignite/pkg/cosmosclient"
    "evmstorechain/x/evmstorechain/types"
)

type EthState struct {
    Blocknumber *big.Int
    BorrowIndex *big.Int
}

const contractAddress = "0xf650C3d88D12dB855b8bf7D11Be6C55A4e07dCC9"

func main() {
    ethState := getStateFromEthereum()
    PostEthereumStateToCosmos(ethState.Blocknumber, ethState.BorrowIndex)
    QueryEthereumStateFromCosmos(ethState.Blocknumber)
}

func getStateFromEthereum() *EthState {
    abiBytes, err := ioutil.ReadFile("abi.json")
    if err != nil {
        log.Fatal(err)
    }

    // Convert []byte to string
    contractAbi := string(abiBytes)

    fmt.Println("Querying borrowIndex from Ethereum")

    // Get Infura URL from environment variable
    infuraURL := os.Getenv("INFURA_URL")

    client, err := ethclient.Dial(infuraURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    parsedABI, err := abi.JSON(strings.NewReader(contractAbi))
    if err != nil {
        log.Fatalf("Failed to parse contract ABI: %v", err)
    }

    contractAddress := common.HexToAddress(contractAddress)

    borrowIndex := parsedABI.Methods["borrowIndex"]

    callMsg := ethereum.CallMsg{
        To:   &contractAddress,
        Data: borrowIndex.ID,
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatalf("Failed to get header: %v", err)
    }

    res, err := client.CallContract(context.Background(), callMsg, header.Number)
    if err != nil {
        log.Fatalf("Failed to execute function call: %v", err)
    }

    borrowIndexRes := new(big.Int)
    borrowIndexRes.SetBytes(res)

    fmt.Printf("BorrowIndex: %s\n", borrowIndexRes.String()) // Print the result

    // Print the borrowIndex
    log.Printf("Block number: %s, Borrow Index: %s", header.Number.String(), borrowIndexRes.String())

    return &EthState{
        Blocknumber: header.Number,
        BorrowIndex: borrowIndexRes,
    }
}

func PostEthereumStateToCosmos(blockNum, borrowIndex *big.Int) {
    ctx := context.Background()
    addressPrefix := "cosmosvaloper"

    //cosmos client instance
    client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
    if err!= nil {
        log.Fatal(err)
    }

    accountName := "alice" //posting as alice (initialised during ignite chain serve)

    account, err := client.Account(accountName)
    if err!= nil {
        log.Fatal(err)
    }

    addr, err := account.Address(addressPrefix)
    if err != nil {
        log.Fatal(err)
    }

    // msg to submit a state
    msg := &types.MsgSubmitEthereumState{
        Creator:     addr,
        Blocknumber: blockNum.Uint64(),
        State:       borrowIndex.Uint64(),
    }

    //broadcast transaction from accountName "alice"
    txResp, err := client.BroadcastTx(ctx, account, msg)
    if err !=nil {
        fmt.Println("here")
        log.Fatal(err)
    }
    fmt.Println("Ethereum State Submited")
    fmt.Println(txResp)
}

func QueryEthereumStateFromCosmos(blockNum *big.Int) {
    ctx := context.Background()
    client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("cosmos"))
    if err != nil {
        log.Fatalf("Failed to connect to the Cosmos client: %v", err)
    }
    queryClient := types.NewQueryClient(client.Context())
    queryResp, err := queryClient.Blockstoragestate(ctx, &types.QueryGetBlockstoragestateRequest{ Blocknumber: strconv.FormatUint(blockNum.Uint64(),10)})
    if err!=nil {
        log.Fatal(err)
    }
    fmt.Println("Query Result")
    fmt.Println(queryResp)
}