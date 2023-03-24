package bloom

import (
	"crypto/sha256"
	"hash/fnv"
	"math"
	"math/big"
)

type BloomFilter struct {
	bitfield []byte
	rounds   int
	hashFunc func([]byte) []byte
}

func CalculateOptimalParameters(expectedItemCount int, falsePositivePercentage float64) (int, int) {
	falsePositiveRate := float64(falsePositivePercentage) / 100
	bitCount := float64(-expectedItemCount) * math.Log(falsePositiveRate) / (math.Ln2 * math.Ln2)
	rounds := bitCount / float64(expectedItemCount) * math.Ln2
	return int(math.Ceil(bitCount)), int(math.Ceil(rounds))
}

func NewBloomFilter(bitfieldLength, rounds int, hashFunc func([]byte) []byte) *BloomFilter {
	byteCount := bitfieldLength / 8
	bitfield := make([]byte, byteCount)
	return &BloomFilter{
		bitfield: bitfield,
		rounds:   rounds,
		hashFunc: hashFunc,
	}
}

func (bf *BloomFilter) Add(input []byte) {
	for i := 0; i < bf.rounds; i++ {
		input = bf.hashFunc(input)
		bitPos := bf.getBitPos(input)
		bytePos, byteOffset := bitPos/8, bitPos%8
		bf.bitfield[bytePos] |= (1 << (7 - uint8(byteOffset)))
	}
}

func (bf *BloomFilter) IsMember(input []byte) bool {
	for i := 0; i < bf.rounds; i++ {
		input = bf.hashFunc(input)
		bitPos := bf.getBitPos(input)
		bytePos, byteOffset := bitPos/8, bitPos%8
		if bf.bitfield[bytePos]&(1<<(7-uint8(byteOffset))) == 0 {
			return false
		}
	}

	return true
}

func (bf *BloomFilter) getBitPos(hash []byte) uint64 {
	bigint := new(big.Int).SetBytes(hash)
	return bigint.Mod(bigint, big.NewInt(int64(len(bf.bitfield)*8))).Uint64()
}

func HashFuncFnv(input []byte) []byte {
	hasher := fnv.New64()
	hasher.Write(input)
	return hasher.Sum(nil)
}

// // Using sha256 instead of sha1 because the latter is deprecated and produces many false positives
func HashFuncSha(input []byte) []byte {
	hasher := sha256.New()
	hasher.Write(input)
	return hasher.Sum(nil)
}

// func HashFuncSha(input []byte) []byte {
// 	hasher := sha1.New()
// 	hasher.Write(input)
// 	return hasher.Sum(nil)
// }
