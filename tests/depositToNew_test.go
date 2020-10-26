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

func TestDepositToNew(t *testing.T) {
	expectedStateHash := "0xf157a013ca602cf88ef45ec74ce78a361f4d592ddbbc3dde4b28fed1ef513809"

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
			8: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(50000),
					1: big.NewInt(6000000),
				},
				Pubkey:  hexutil.MustDecode("0x0c9bc9ed7b58277d5f9036c85e47958c65bc81104718a9364a294d96b4d277da"),
				Address: common.HexToAddress("0xdc70a72abf352a0e3f75d737430eb896ba9bf9ea"),
			},
			12: {
				Tokens: map[uint16]*big.Int{
					0: big.NewInt(30000),
					1: big.NewInt(1000000),
					2: big.NewInt(5000000),
				},
				Pubkey:  hexutil.MustDecode("0x01820b7899ad2a62a1c4aacf320b1a528c8c98aa558ee777e60110be62626e42"),
				Address: common.HexToAddress("0x052f46feb45822e7f117536386c51b6bd3125157"),
			},
		},
		AccountMax: 18,
		LooMax:     0,
	})

	blk := &types.MiniBlock{
		Txs: []types.Transaction{&types.DepositToNewOp{
			PubKey:     hexutil.MustDecode("0x04b038a79fa8f5105ac00b983c1263562dc1efdf459535e43e007162f11ca3d7"),
			WithdrawTo: common.HexToAddress("0x91f4d9ea5c1ee0fc778524b3d57fd8cf700996cf"),
			TokenID:    2,
			Amount:     big.NewInt(45242000),
		}},
	}

	bc.AddMiniBlock(blk)

	actualStateHash := bc.GetStateData().Hash().String()
	fmt.Println(actualStateHash)
	require.Equal(t, expectedStateHash, actualStateHash)
}
