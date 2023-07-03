package main

import (
    "context"
    "strconv"
    "fmt"
    "log"
     "github.com/ignite/cli/ignite/pkg/cosmosclient"
     "evmstorechain/x/evmstorechain/types"
)

func main() {
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
    var blocknum uint64 = 1
    var state uint64 = 12
    msg := &types.MsgSubmitEthereumState{
        Creator:     addr,
        Blocknumber: blocknum,
        State:       state,
    }

    //broadcast transaction from accountName "alice"
    txResp, err := client.BroadcastTx(ctx, account, msg)
    if err !=nil {
        fmt.Println("here")
        log.Fatal(err)
    }
    fmt.Println("Ethereum State Submited")
    fmt.Println(txResp)

    // Query
    queryClient := types.NewQueryClient(client.Context())
    queryResp, err := queryClient.Blockstoragestate(ctx, &types.QueryGetBlockstoragestateRequest{ Blocknumber: strconv.FormatUint(blocknum,10)})
    if err!=nil {
        log.Fatal(err)
    }
    fmt.Println(queryResp)
}