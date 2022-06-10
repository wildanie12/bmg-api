package main

import (
	_config "gin-bmg-restful/config"
	_utils "gin-bmg-restful/utils"
)

func main() {

	config := _config.New()
	db := _utils.NewPostgresConnection(config)
	_utils.Migrate(db)


	// r := gin.Default()
	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON("")
	// })
}