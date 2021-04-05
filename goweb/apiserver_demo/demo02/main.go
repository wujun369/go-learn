package main

import (
	"errors"
	"github.com/spf13/pflag"
	"goweb/apiserver_demo/demo01/router"
	"goweb/apiserver_demo/demo02/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
	/*
		pflag.StringP：从命令行中输入值给 cfg
		name：该值的名称
		shorthand：名称缩写
		value：具体输入的值，最后赋值给 cfg
		usage：解释
	 */

)

func main() {
	pflag.Parse() //解析命令行输入值


	// 解析配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 设置 gin 的运行模型，从 config.yaml 配置文件中获取，类似java的properties
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{} //获取所有中间件

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
