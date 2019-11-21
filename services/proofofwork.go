package services

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/xuelang-algo/blockchain_go/utils"
)

var (
	maxNonce = math.MaxInt64
	intOne = big.NewInt(1)
)

const threads = 32
const targetBits = 16

// ProofOfWork represents a proof-of-work
type ProofOfWork struct {
	block  *Block
	prepare []byte
	target *big.Int
}

// NewProofOfWork builds and returns a ProofOfWork
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, []byte{}, target}
	pow.prepare = bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.HashTransactions(),
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
		},
		[]byte{},
	)
	return pow
}

func (pow *ProofOfWork) prepareData(nonce []byte) []byte {
	data := bytes.Join(
		[][]byte{
			pow.prepare,
			nonce,
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
func (pow *ProofOfWork) Run() ([]byte, []byte) {
	var done *big.Int
	done = nil
	nonce := big.NewInt(0)
	var hash [32]byte

	fmt.Printf("Mining a new block")

	pool := utils.NewPool(threads)
	for done == nil {
		for i := 0; i < threads; i++ {
			testNonce := new(big.Int)
			testNonce.Add(nonce, testNonce)
			pool.Add(1)
			go func(){
				defer pool.Done()
				var hashInt big.Int
				var hash [32]byte

				data := pow.prepareData(testNonce.Bytes())
				hash = sha256.Sum256(data)
				hashInt.SetBytes(hash[:])

				if hashInt.Cmp(pow.target) == -1 {
					done = testNonce
				}
				time.Sleep(1*time.Second)
			}()

			nonce.Add(nonce, intOne)
		}

	}
	pool.Wait()

	data := pow.prepareData(done.Bytes())
	hash = sha256.Sum256(data)
	nonce = done
	fmt.Print("\n\n")

	return nonce.Bytes(), hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
