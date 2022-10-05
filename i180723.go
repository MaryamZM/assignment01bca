package assignment01bca

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type block struct {
	x           int
	hash        string
	prev_hash   string
	transaction string
}
type BlockChain struct {
	list []*block
}

func NewBlock(transaction string, nonce int, previousHash string, blockchain *BlockChain) *block {
	block1 := new(block)
	block1.transaction = transaction
	block1.x = nonce
	block1.prev_hash = previousHash
	block1.hash = CalculateHash(block1.transaction + strconv.Itoa(block1.x) + block1.prev_hash)
	blockchain.list = append(blockchain.list, block1)
	return block1
}

func DisplayBlocks(blockchain *BlockChain) {
	for i, a := range blockchain.list {
		fmt.Printf("%s BLOCK %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Printf(" TRANSACTION: %s \n NONCE VALUE: %d \n HASH OF PREVIOUS BLOCK : %s \n HASH OF CURRENT BLOCK %s \n \n ", a.transaction, a.x, a.prev_hash, a.hash)
	}

}

func ChangeBlock(blockchain *BlockChain) {

	var index int
	fmt.Println("enter index of block you want to edit:")
	fmt.Scan(&index)
	var chainLength int
	chainLength = len(blockchain.list)
	if index < chainLength {

		fmt.Println("Your current transaction is as follows \n")
		fmt.Printf("%s", blockchain.list[index].transaction)
		scan := bufio.NewScanner(os.Stdin)
		fmt.Println("enter new transaction: \n")
		scan.Scan()
		text := scan.Text()
		blockchain.list[index].transaction = text
		fmt.Println("changes have been made!")

	}
}

func VerifyChain(blockchain *BlockChain) {

	var verify = false
	for _, num := range blockchain.list {

		Hash := CalculateHash(num.transaction + strconv.Itoa(num.x) + num.prev_hash)
		if Hash != num.hash {

			verify = true
			break
		}
	}
	if verify == false {

		fmt.Println("verification complete, no changes detected")
	} else {
		fmt.Println("change detected in a block")
	}
}

func CalculateHash(stringToHash string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}
