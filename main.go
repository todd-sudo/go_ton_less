package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func main() {
	client := liteclient.NewConnectionPool()

	configUrl := "https://ton-blockchain.github.io/testnet-global.config.json"
	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		log.Println(err)
	}
	api := ton.NewAPIClient(client)
	words := strings.Split("property fetch original vague cloud pulp light hub chief bid spoil swamp wealth long topic leisure pelican assume awkward yard brand virtual doctor jazz", " ")

	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		log.Println(err)
	}

	block, err := api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		log.Println(err)
	}

	balance, err := w.GetBalance(context.Background(), block)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(balance)
	fmt.Println(w.Address())

	addr := address.MustParseAddr("EQCD39VS5jcptHL8vMjEXrzGaRcCVYto7HUn4bpAOg8xqB2N")
	err = w.Transfer(context.Background(), addr, tlb.MustFromTON("0.003"), "Hey bro, happy birthday!")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ok ok ok")
}
