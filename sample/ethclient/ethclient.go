package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	gosdk "github.com/okex/exchain-go-sdk"
)

const (
	host     string = "https://exchaintestrpc.okex.org"
	alice    string = "0x2CF4ea7dF75b513509d95946B43062E26bD88035"
	bob      string = "0x0073F2E28ef8F117e53d858094086Defaf1837D5"
	aliceKey string = "e47a1fe74a7f9bfa44a362a3c6fbe96667242f62e6b8e138b3f61bd431c3215d"
)

func main() {

	client, err := gosdk.NewEthClient(context.Background(), host)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	fmt.Println("gasPrice", gasPrice, err)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	fmt.Println("balance:", balance, err)

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	fmt.Println("nonce", nonce, err)

	pendingNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(alice))
	fmt.Println("pendingNonce", pendingNonce, err)

	chainID, err := client.ChainID(context.Background())
	fmt.Println("chainID", chainID, err)

	privateKey, _ := crypto.HexToECDSA(aliceKey)
	unsignedTx := types.NewTransaction(pendingNonce, common.HexToAddress(bob), big.NewInt(1000000000000000000), 30000, gasPrice, []byte{})
	signedTx, _ := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)

	err = client.SendTransaction(context.Background(), signedTx)
	time.Sleep(time.Second * 4)
	fmt.Println("sendTx err:", err)
	fmt.Println("txHash", signedTx.Hash())

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	fmt.Printf("recipt %+v\n", receipt, err)

	pendingCode, err := client.PendingCodeAt(context.Background(), common.HexToAddress(alice))
	fmt.Println("pendingCode", pendingCode, err)
	code, err := client.CodeAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	fmt.Println("code", code, err)

	to := common.HexToAddress(bob)
	msg := ethereum.CallMsg{From: common.HexToAddress(alice), To: &to, GasPrice: gasPrice, Value: big.NewInt(1), Data: []byte{}}
	estimateGas, err := client.EstimateGas(context.Background(), msg)
	fmt.Println("estimateGas", estimateGas, err)

	re, err := client.CallContract(context.Background(), msg, big.NewInt(1))
	fmt.Println("callContract", re, err)

	block, err := client.BlockByNumber(context.Background(), nil)
	fmt.Println("block", block, err)

	ss, err := client.SubscribeNewHead(context.Background(), make(chan *types.Header))
	fmt.Println("subscription", ss, err)
	es, err := client.EthSubscribe(context.Background(), make(chan *types.Header), "newHeads")
	fmt.Println("ethSubscribe", es, err)

	var hex hexutil.Big
	err = client.CallContext(context.Background(), &hex, "eth_gasPrice")
	fmt.Println("CallContext", hex, err)
}
