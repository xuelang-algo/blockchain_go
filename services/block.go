package services

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/xuelang-algo/blockchain_go/utils"
)

// Block represents a block in the blockchain
//type Block struct {
//	Timestamp     int64
//	Transactions  []*Transaction
//	PrevBlockHash []byte
//	Hash          []byte
//	Nonce         int
//	Height        int
//}


// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := utils.NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// Serialize serializes the block
// use protobuf serialize function
func (b *Block) Serialize() []byte {
	//var result bytes.Buffer
	//encoder := gob.NewEncoder(&result)
	//err := encoder.Encode(b)
	result , err := proto.Marshal(b)

	if err != nil {
		log.Panic(err)
	}

	//return result.Bytes()
	return result
}

// DeserializeBlock deserializes a block
// use protobuf unserialize function
func DeserializeBlock(d []byte) *Block {
	var block Block
	//decoder := gob.NewDecoder(bytes.NewReader(d))
	//err := decoder.Decode(&block)
	err := proto.Unmarshal(d, &block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
