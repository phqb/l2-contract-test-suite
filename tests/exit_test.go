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

func TestExit(t *testing.T) {
	expectedStateHash := "0xb6b87a880f587de7f674927c3bd7d89eefc322be05692e5b702971d9f113d87b"

	bc := blockchain.NewBlockchain(&blockchain.Genesis{
		AccountAlloc: map[uint32]blockchain.GenesisAccount{
			0: {
				Tokens:  map[uint16]*big.Int{},
				Pubkey:  hexutil.MustDecode("0x06ee0abfc4fb2ce0dc84a7726b79cc354484205b509a0c0745d394503f513726"),
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			36: {
				Tokens: map[uint16]*big.Int{
					5: big.NewInt(500),
				},
				Pubkey:  hexutil.MustDecode("0x0359781e83863e6945f7704baeb504c36609b87fab674eec37439c87aea435a1"),
				Address: common.HexToAddress("0xc783df8a850f42e7f7e57013759c285caa701eb6"),
			},
			44: {
				Tokens:  map[uint16]*big.Int{},
				Pubkey:  hexutil.MustDecode("0x064d9b32b5812260f75c844aa11e79fe0bf986b202d31496360f3d5304ba6de0"),
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

	blk1 := &types.MiniBlock{
		Txs: []types.Transaction{&types.DepositOp{
			AccountID: 36,
			TokenID:   2,
			Amount:    big.NewInt(45242000),
		}, &types.DepositOp{
			AccountID: 36,
			TokenID:   4,
			Amount:    big.NewInt(135000),
		}},
	}

	bc.AddMiniBlock(blk1)

	blk2 := &types.MiniBlock{
		Txs: []types.Transaction{&types.ExitOp{
			AccountID: 36,
		}},
	}

	bc.AddMiniBlock(blk2)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
