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

func TestSettlement2(t *testing.T) {
	expectedStateHash := "0xd237077d0cd33a50b2ead2ee899eeec1bccc1f7cb28d3609a6d5f7559a2dec89"

	bc := blockchain.NewBlockchain(&blockchain.Genesis{
		AccountAlloc: map[uint32]blockchain.GenesisAccount{
			0: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(2000000),
				},
				Pubkey:  hexutil.MustDecode("0x0cfb7e0cf1380065477345a42aa821aa1c68e7d9eb213eee1e8f00cb707458a4"),
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			17: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(70000),
					1: big.NewInt(6000000),
				},
				Pubkey:  hexutil.MustDecode("0x0c9bc9ed7b58277d5f9036c85e47958c65bc81104718a9364a294d96b4d277da"),
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			123: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(1000000),
					2: big.NewInt(5000000),
				},
				Pubkey:  hexutil.MustDecode("0x01820b7899ad2a62a1c4aacf320b1a528c8c98aa558ee777e60110be62626e42"),
				Address: common.HexToAddress("0x052f46feb45822e7f117536386c51b6bd3125157"),
			},
		},
		AccountMax: 1000,
		LooMax:     289,
		LooAlloc: map[uint64]*types.LeftOverOrder{
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
			&types.Settlement2{
				OpType:     types.SettlementOp21,
				LooID1:     243,
				AccountID2: 123,
				Rate2: types.PackedAmount{
					Mantisa: 1,
					Exp:     18,
				},
				Amount2: types.PackedAmount{
					Mantisa: 3,
					Exp:     6,
				},
				Fee2: types.PackedFee{
					Mantisa: 4,
					Exp:     2,
				},
				ValidSince2:  1600661873,
				ValidPeriod2: 8640000,
			},
		},
	}

	_ = bc.AddMiniBlock(blk)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
