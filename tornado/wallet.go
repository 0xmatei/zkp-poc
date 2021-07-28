package tornado

import (
	"fmt"

	"github.com/mateigraura/zkp-poc/mock"
	"github.com/mateigraura/zkp-poc/zkp"
)

type Wallet struct {
	Owner     string
	Addresses map[string]string
}

func (w *Wallet) NewWallet(owner string) {
	w.Addresses = mock.Commitments()
	w.Owner = owner
}

func (w *Wallet) MixerDeposit(receiverAddress string, client Client, ScAddress Mixer) {
	fmt.Printf("Signing TX for x coins for owner: %s\n", w.Owner)
	fmt.Printf("Receiver address: %s\n", receiverAddress)

	var txPayload zkp.HashCircuit
	receiver := w.Addresses[receiverAddress]
	commitment := txPayload.Commit([]byte(receiver))

	client.SendFunds(txPayload, commitment, ScAddress)
}

func (w *Wallet) MixerWithdraw(client Client, ScAddress Mixer) {
	fmt.Printf("Initiating claim TX for x coins for owner: %s\n", w.Owner)

	var txPayload zkp.HashCircuit
	walletAddress := w.Addresses[w.Owner]
	commitment := txPayload.Commit([]byte(walletAddress))

	client.Withdraw(txPayload, commitment, ScAddress)
}
