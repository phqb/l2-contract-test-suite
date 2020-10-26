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

func TestSettlement3(t *testing.T) {
	var (
		expectedStateHash = "0xb173c4502945ae6243acc1837bc8be5a8187956d43278ac4561c42c0cd8f34d2"
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

			17: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(70000),
					1: big.NewInt(6000000),
				},
				Pubkey:  pubKey2,
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			30: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(1000000),
					2: big.NewInt(5000000),
				},
				Pubkey:  pubKey3,
				Address: common.HexToAddress("0x052f46feb45822e7f117536386c51b6bd3125157"),
			},
		},
		AccountMax: 1000,
		LooMax:     289,
		LooAlloc: map[uint64]*types.LeftOverOrder{
			56: {
				AccountID:   30,
				SrcToken:    2,
				DestToken:   1,
				Amount:      big.NewInt(4321),
				Fee:         big.NewInt(600),
				Rate:        types.PackedAmount{Mantisa: 4, Exp: 18}.Big(),
				ValidSince:  1601436626,
				ValidPeriod: 823000,
			},
			243: {
				AccountID:   17,
				SrcToken:    1,
				DestToken:   2,
				Amount:      big.NewInt(34500),
				Fee:         big.NewInt(67432),
				Rate:        types.PackedAmount{Mantisa: 2, Exp: 17}.Big(),
				ValidSince:  1601436627,
				ValidPeriod: 823000,
			},
		},
	})

	blk := &types.MiniBlock{
		Txs: []types.Transaction{
			&types.Settlement3{
				LooID1: 243,
				LooID2: 56,
			},
		},
	}

	_ = bc.AddMiniBlock(blk)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
