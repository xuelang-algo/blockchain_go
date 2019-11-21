package practice

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/xuelang-algo/blockchain_go/utils"
)

var (
	maxNonce = math.MaxInt64
	intOne = big.NewInt(1)
)

const targetBits = 24
const threads = 32

type ProofOfWork struct {
	block  *Block
	prepare []byte
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, []byte{}, target,}

	pow.prepare = bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
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

func (pow *ProofOfWork) prepareData2(nonce int32) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() ([]byte, []byte) {
	var done *big.Int
	done = nil
	nonce := big.NewInt(0)
	var hash [32]byte

	pool := utils.NewPool(threads)

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for done == nil {
		for i := 0; i < threads; i++ {
			//log.Println(nonce)
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
				//time.Sleep(1*time.Second)
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

func (pow *ProofOfWork) Run2() ([]byte, []byte) {
	var done *big.Int
	var hash [32]byte
	done = nil
	nonce := big.NewInt(0)
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for done == nil {
		for i := 0; i < threads; i++ {
			testNonce := new(big.Int)
			testNonce.Add(nonce, testNonce)
			var hashInt big.Int
			data := pow.prepareData(testNonce.Bytes())
			hash = sha256.Sum256(data)
			hashInt.SetBytes(hash[:])

			if hashInt.Cmp(pow.target) == -1 {
				done = testNonce
			}
			//time.Sleep(1*time.Second)
			nonce.Add(nonce, intOne)
		}
	}
	data := pow.prepareData(done.Bytes())
	hash = sha256.Sum256(data)
	nonce = done
	fmt.Print("\n\n")
	return nonce.Bytes(), hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func (pow *ProofOfWork) ValidateHash() [sha256.Size]byte {
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	return hash
}