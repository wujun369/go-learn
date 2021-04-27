package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const difficulty int = 1 //难度系数为Hash前面至少包含一个 0

/**
区块的数据模型
*/
type Block struct {
	Index      int    //数据的位置
	Timestamp  string //时间戳
	Hash       string //数据的 SHA256 标识符
	PrevHash   string //前一个数据的 SHA256 标识符
	Difficulty int    //挖矿难度
	Nonce      string //PoW 符合条件的数字
}

var BlockChain []Block    //存放区块的数据
var mutex = &sync.Mutex{} //通过锁的方式防止同一时间产生多个区块

/**
处理 GET 请求，创建区块
*/
func HandleGetBlockchain(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	mutex.Lock()
	newBlock := generateBlock(BlockChain[len(BlockChain)-1]) //创建区块
	mutex.Unlock()

	if isBlockValid(newBlock, BlockChain[len(BlockChain)-1]) { //判断区块是否合法
		BlockChain = append(BlockChain, newBlock) //添加区块(通过数组维护)
		spew.Dump(BlockChain)                     //打印区块链的详细信息
	}
	respondWithJSON(c, http.StatusCreated, newBlock) //响应客户端，显示新的区块
}

/**
处理 POST 请求
*/
func HandleWriteBlock(c *gin.Context) {
	bytes, err := json.MarshalIndent(BlockChain, "", " ") //解析成 json 数据
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(c.Writer, string(bytes)) // 向客户端打印所有区块信息
}

/**
响应服务器错误信息（数据传输的错误）
*/
func respondWithJSON(c *gin.Context, code int, payload interface{}) {

	c.Header("Content-Type", "application/json")

	response, err := json.MarshalIndent(payload, "", " ")

	if err != nil {

		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	c.Writer.WriteHeader(code)
	c.Writer.Write(response)
}

/**
生成新区块
*/
func generateBlock(oldBlock Block) Block {

	var newBlock Block //新区块
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1 //新增区块，index + 1
	newBlock.Timestamp = t.String()     //为新区块添加时间戳
	newBlock.PrevHash = oldBlock.Hash   //新区快存储上一个区块的 Hash 值
	newBlock.Difficulty = difficulty
	for i := 0; ; i++ { //通过循环改变 Nonce 值
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex //选出符合难度的 Nonce

		hash := calculateHash(newBlock)
		flag := isHashValid(hash, difficulty)

		if flag { //判断 Hash 的0的个数是否与难度系数一致
			fmt.Println(calculateHash(newBlock), "挖矿成功.....")
			newBlock.Hash = calculateHash(newBlock)
			break
		} else {
			fmt.Println(calculateHash(newBlock), "挖矿中......")
			time.Sleep(time.Second)
			continue
		}
	}
	return newBlock
}

func isHashValid(hash string, difficulty int) bool {

	prefix := strings.Repeat("0", difficulty) //复制 difficulty 个 0
	return strings.HasPrefix(hash, prefix)    //判断 hash 是否有 difficulty 个 0
}

/**
生成 Hash 序列
*/
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash + block.Nonce
	h := sha256.New() //获得一个 SHA256 校验算法的接口
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

/**
验证区块是否正确
*/

func isBlockValid(newBlock Block, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false //确认index 的增长是否正确
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false //确认 PreHash 是否与前一个块的 Hash 值相同
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false //判断Hash值是否有变更
	}
	return true
}
