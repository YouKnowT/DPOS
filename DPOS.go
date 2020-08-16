package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Block struct {
	Index     int
	TimeStamp string
	BPM       int
	Hash      string
	PrevHash  string
	Delegate  string
}

// 产生新的区块
func generateBlock(oldBlock Block, _BMP int, address string) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.TimeStamp = t.String()
	newBlock.BPM = _BMP
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = createBlockHash(newBlock)
	newBlock.Delegate = address
	fmt.Println("NewBlock: ", newBlock)
	return newBlock, nil
}

// 生成区块的hash = sha256('当前区块的index序号' +  '时间戳' + '区块的BPM' + '上一个区块的hash').string()
func createBlockHash(block Block) string {
	record := string(block.Index) + block.TimeStamp + string(block.BPM) + block.PrevHash
	sha3 := sha256.New()
	sha3.Write([]byte(record))
	hash := sha3.Sum(nil)
	fmt.Println("NewHash: ", hex.EncodeToString(hash))
	return hex.EncodeToString(hash)
}

//检视区块是否合法
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		fmt.Println("失败！！index非法")
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		fmt.Println("失败！！PrevHash非法")
		return false
	}
	fmt.Println("合法")

	return true
}

var blockChain []Block

type Node struct {
	name string
	coin float32
}
type Trustee struct {
	name   string
	coin   float32
	age    float32
	credit float32
	super  float32
}

type trusteeList []Trustee

var _trusteeList = []Trustee{
	{"node1", 0, 30.0, 500.0, 0.0},
	{"node2", 0, 30.0, 500.0, 0.0},
	{"node3", 0, 30.0, 500.0, 0.0},
	{"node4", 0, 30.0, 500.0, 0.0},
	{"node5", 0, 30.0, 500.0, 0.0},
	{"node6", 0, 30.0, 500.0, 0.0},
	{"node7", 0, 30.0, 500.0, 0.0},
	{"node8", 0, 30.0, 500.0, 0.0},
	{"node9", 0, 30.0, 500.0, 0.0},
	{"node10", 0, 30.0, 500.0, 0.0},
	{"node11", 0, 30.0, 500.0, 0.0},
	{"node12", 0, 30.0, 500.0, 0.0},
}

func (_trusteeList trusteeList) Len() int {
	return len(_trusteeList)
}

func (_trusteeList trusteeList) Swap(i, j int) {
	_trusteeList[i], _trusteeList[j] = _trusteeList[j], _trusteeList[i]
}

func (_trusteeList trusteeList) Less(i, j int) bool {
	return _trusteeList[j].super < _trusteeList[i].super
}

func selecTrustee() []Trustee {
	var _trusteeList1 []Trustee
	for i := 0; i < len(_trusteeList); i++ {
		if _trusteeList[i].age != 0 {
			_trusteeList[i].age -= 1
		}
		if _trusteeList[i].credit > 1000 {
			_trusteeList[i].credit = 500
		}
		if _trusteeList[i].coin > 1500 {
			_trusteeList[i].coin = 1500
		}
		//_trusteeList[i].coin = _node[i].coin
		_trusteeList[i].super = 0.3*_trusteeList[i].age*float32(_trusteeList[i].coin) + 0.7*_trusteeList[i].credit
	}
	sort.Sort(trusteeList(_trusteeList))
	fmt.Println(_trusteeList)
	result := _trusteeList[:5]
	_trusteeList1 = result[1:]
	_trusteeList1 = append(_trusteeList1, result[0])
	fmt.Println("\n当前超级节点代表列表是：", _trusteeList1)
	return _trusteeList1
}

func main() {

	var selection string
	type t1 []string
	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, createBlockHash(Block{}), "", ""}
	fmt.Println("创世块block: ", genesisBlock)
	blockChain = append(blockChain, genesisBlock)
	var trustee Trustee
	rand.Seed(time.Now().UnixNano()) //random
	var _node = []Node{
		{"node1", rand.Float32() * 100},
		{"node2", rand.Float32() * 100},
		{"node3", rand.Float32() * 100},
		{"node4", rand.Float32() * 100},
		{"node5", rand.Float32() * 100},
		{"node6", rand.Float32() * 100},
		{"node7", rand.Float32() * 100},
		{"node8", rand.Float32() * 100},
		{"node9", rand.Float32() * 100},
		{"node10", rand.Float32() * 100},
		{"node11", rand.Float32() * 100},
		{"node12", rand.Float32() * 100},
	}
	for i := 0; i < len(_node); i++ {
		_trusteeList[i].coin = _node[i].coin
	}

L1:
	for i := 0; ; i++ {
		for _, trustee = range selecTrustee() {
			time.Sleep(2000000000)
			_BPM := rand.Intn(100)
			blockHeight := len(blockChain)
			oldBlock := blockChain[blockHeight-1]
			newBlock, err := generateBlock(oldBlock, _BPM, trustee.name)
			if err != nil {
				fmt.Println("新生成区块失败：", err)
				continue
			}
			if isBlockValid(newBlock, oldBlock) {
				blockChain = append(blockChain, newBlock)
				fmt.Println("当前操作区块节点为：", trustee.name)
				for j := 0; j < len(_trusteeList); j++ {
					if _trusteeList[j].name == trustee.name {
						_trusteeList[j].credit += 10
						_trusteeList[j].age = 0
						_trusteeList[j].coin += (_trusteeList[j].super / 365) * 0.04
						break
					}
				}
				fmt.Println("当前区块数：", len(blockChain))
				fmt.Println("当前区块信息：", blockChain[len(blockChain)-1])
			} else {
				fmt.Println("当前操作区块节点为：", trustee.name)
				for j := 0; j < len(_trusteeList); j++ {
					if _trusteeList[j].name == trustee.name {
						_trusteeList[j].credit -= 200
						//_trusteeList[j].age=0
						//_trusteeList[j].coin+=(_trusteeList[j].super/365)*0.04
						break
					}
				}
			}
		}
		fmt.Println("\ncontinue or not?\nA.Continue\nB.Not")
		fmt.Scanf("%s", &selection)
		if selection != "A" {
			return
		} else {
			goto L1
		}
	}

}
