package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	env := GetEnv()
	ConnectDb(env["POSTGRES_PASSWORD"])
	CreateTable()
	r := gin.Default()

	r.GET("/:number", GetOddOrEven)

	r.Run(":3007") // listen and serve on 0.0.0.0:3007
}
