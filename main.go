package main

import (
	"fmt"
	manager "mybio_server/fileManager"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//get query params
	router.GET("/backend/details", func(c *gin.Context) {
		id := c.Query("id")
		mType := c.Query("type")
		// c.JSON(200, gin.H{
		// 	"id":  id,
		// 	"age": mType,
		// })
		res := manager.GetFileByType(mType, id)
		fmt.Print(res)
		c.JSON(200, res)

	})

	router.Run(":3000")

}
