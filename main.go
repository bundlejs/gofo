package main

import (
	"./component/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	config := conf.Config()

	//router := gin.Default()
	//
	//router.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%s", config.Server.Port),
	//	Handler:        router,
	//	ReadTimeout:    config.Server.ReadTimeout,
	//	WriteTimeout:   config.Server.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//if err := s.ListenAndServe(); err != nil {
	//	fmt.Println(err)
	//}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(fmt.Sprintf(":%s", config.Server.Port)); err != nil {
		fmt.Println(err)
	}
}
