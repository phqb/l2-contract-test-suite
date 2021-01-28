package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func getRootHashFromBatchProof(keys []uint64, values []common.Hash, siblings []common.Hash, deep int) common.Hash {
	var (
		tmpKeys   []uint64
		tmpValues []common.Hash
	)
	tmpKeys = append(tmpKeys, keys...)
	tmpValues = append(tmpValues, values...)
	siblingsIndex := 0
	numAccounts := len(keys)

	for currentDeep := 0; currentDeep < deep; currentDeep++ {
		accountIndex := 0
		for i := 0; i < numAccounts; {
			if i != numAccounts-1 && tmpKeys[i]/2 == tmpKeys[i+1]/2 {
				tmpKeys[accountIndex] = tmpKeys[i] / 2
				tmpValues[accountIndex] = GetRoot(tmpValues[i], tmpValues[i+1])
				accountIndex++
				i += 2
				continue
			}

			if (tmpKeys[i] & 1) == 0 { // right siblings
				tmpValues[accountIndex] = GetRoot(tmpValues[i], siblings[siblingsIndex])
			} else { // left siblings
				tmpValues[accountIndex] = GetRoot(siblings[siblingsIndex], tmpValues[i])
			}
			tmpKeys[accountIndex] = tmpKeys[i] / 2
			accountIndex++
			siblingsIndex++
			i++
		}
		numAccounts = accountIndex
	}
	if numAccounts != 1 {
		panic("bug")
	}
	return tmpValues[0]
}

func TestMerkleTree_GetProofBatch(t *testing.T) {
	var (
		v1 = common.HexToHash("0x45")
		v2 = common.HexToHash("0x678")
	)

	tr := NewTree(5)
	tr.Update(3, v1)
	tr.Update(5, v2)
	tr.Update(1, common.HexToHash("0xabcd"))

	keys := []uint64{3, 5}

	values, siblings := tr.GetProofBatch(keys)
	require.Equal(t, values[0], v1)
	require.Equal(t, values[1], v2)

	recoverRootHash := getRootHashFromBatchProof(keys, values, siblings, 4)
	require.Equal(t, recoverRootHash.Hex(), tr.RootHash().Hex())
}

func TestMerkleTree_RootHash(t *testing.T) {
	fmt.Println("[")

	for i := 0; i < 1024; i ++ {
		height := uint(math.Ceil(math.Log2(float64(i + 1))))

		tr := NewTree(height + 1)

		for e := 0; e <= i; e ++ {
			h := crypto.Keccak256Hash([]byte(strconv.Itoa(e + 1)))
			tr.Update(uint64(e), common.BytesToHash(h[:]))
		}

		rootHash := tr.RootHash()
		fmt.Print("    \"")
		fmt.Print(rootHash.String()[2:])
		fmt.Println("\",")
	}

	fmt.Println("];")
}