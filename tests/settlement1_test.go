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

func TestSettlement1(t *testing.T) {
	var (
		expectedStateHash = "0xafcb21d2384e3167008e1df50c5ce768784ea71d8e2ea01a9b6ecb84c98bb1d1"
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
		Txs: []types.Transaction{
			&types.Settlement1{
				OpType:   types.SettlementOp11,
				Token1:   1,
				Token2:   2,
				Account1: 8,
				Account2: 12,
				Rate1: types.PackedAmount{
					Mantisa: 1,
					Exp:     18,
				},
				Rate2: types.PackedAmount{
					Mantisa: 1,
					Exp:     18,
				},
				Amount1: types.PackedAmount{
					Mantisa: 2,
					Exp:     6,
				},
				Amount2: types.PackedAmount{
					Mantisa: 3,
					Exp:     6,
				},
				Fee1: types.PackedFee{
					Mantisa: 7,
					Exp:     3,
				},
				Fee2: types.PackedFee{
					Mantisa: 4,
					Exp:     2,
				},
				ValidSince1:  1600661872,
				ValidSince2:  1600661873,
				ValidPeriod1: 86400,
				ValidPeriod2: 86400,
			},
		},
	}

	_ = bc.AddMiniBlock(blk)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
