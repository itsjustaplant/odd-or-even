package main

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"strconv"
)

func main() {
	env := GetEnv()
	ConnectDb(env["POSTGRES_PASSWORD"])
	// CreateTable()
	r := gin.Default()

	r.GET("/:number", func(c *gin.Context) {
		number := c.Param("number")
		parsedNumber, err := strconv.Atoi(number)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Provided param should be a number",
			})
		} else {
			foundValue, statusCode := GetNumber(parsedNumber)
			if statusCode == http.StatusOK {
				c.JSON(statusCode, gin.H{
					"number": foundValue.Number,
					"isOdd":  foundValue.IsOdd,
				})
			} else {
				c.JSON(statusCode, gin.H{
					"error": "Requested number is not found",
				})
			}
		}
	})

	r.Run(":3007") // listen and serve on 0.0.0.0:3007
}
