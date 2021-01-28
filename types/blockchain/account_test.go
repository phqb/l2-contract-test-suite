package blockchain

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestAccountBalanceRoot(t *testing.T) {

	var (
		pubKey1, _ = hexutil.Decode("0x0cfb7e0cf1380065477345a42aa821aa1c68e7d9eb213eee1e8f00cb707458a4")
		ethAddr    = common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea")
		acc        = NewAccount(pubKey1, ethAddr)
	)
	acc.Update(0, big.NewInt(30_000))
	acc.Update(1, big.NewInt(2_000_000))
	balanceRoot := acc.tree.RootHash()

	log.Printf("root tree is %s", balanceRoot.String())

}
