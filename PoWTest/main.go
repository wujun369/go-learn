package main

import (
	"PoWTest/router"
	"PoWTest/router/handler"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex = &sync.Mutex{} //通过锁的方式防止同一时间产生多个区块

func main() {

	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	router.Load(g, middlewares...) //添加路由

	go func() {
		genesisBlock := handler.Block{} //创世区块
		spew.Dump(genesisBlock) //打印创世区块
		mutex.Lock()
		handler.BlockChain = append(handler.BlockChain, genesisBlock) //保证同一时间内只能生成一个区块
		mutex.Unlock()
	}()

	log.Printf("开始监听端口：8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

func pingServer() error {

	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://127.0.0.1:8080")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Printf("尝试重新连接(%d)", i)
		time.Sleep(time.Second)
	}
	return errors.New("连接服务器失败")

	return nil
}