package tests

import (
	"fmt"
	"github.com/KyberNetwork/l2-contract-test-suite/types"
	"github.com/KyberNetwork/l2-contract-test-suite/types/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDeposit(t *testing.T) {
	var (
		expectedStateHash = "0x8b4ccabbaf16af8d8c0ef5f894c4bd5cb50057848ed27c9f410059538f09d425"
		pubKey1, _        = hexutil.Decode("0x0cfb7e0cf1380065477345a42aa821aa1c68e7d9eb213eee1e8f00cb707458a4")
		pubKey2, _        = hexutil.Decode("0x0c9bc9ed7b58277d5f9036c85e47958c65bc81104718a9364a294d96b4d277da")
		pubKey3, _        = hexutil.Decode("0x01820b7899ad2a62a1c4aacf320b1a528c8c98aa558ee777e60110be62626e42")
	)

	bc := blockchain.NewBlockchain(&blockchain.Genesis{
		AccountAlloc: map[uint32]blockchain.GenesisAccount{
			0: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(2000000),
				},
				Pubkey:  pubKey1,
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			8: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(50000),
					1: big.NewInt(6000000),
				},
				Pubkey:  pubKey2,
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			12: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(1000000),
					2: big.NewInt(5000000),
				},
				Pubkey:  pubKey3,
				Address: common.HexToAddress("0x052f46feb45822e7f117536386c51b6bd3125157"),
			},
		},
		AccountMax: 18,
		LooMax:     0,
	})

	blk := &types.MiniBlock{
		Txs: []types.Transaction{&types.DepositOp{
			AccountID: 8,
			TokenID:   2,
			Amount:    big.NewInt(45242000),
		}},
	}

	_ = bc.AddMiniBlock(blk)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
