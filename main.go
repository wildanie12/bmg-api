package main

import (
	"fmt"
	_config "gin-bmg-restful/config"
)

func main() {

	config := _config.New()
	fmt.Println(config)

	// r := gin.Default()
	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON("")
	// })
}