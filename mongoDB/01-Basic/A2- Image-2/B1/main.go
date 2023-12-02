/* Download:
	
	go get github.com/gin-gonic/gin
	go get go.mongodb.org/mongo-driver/mongo

*/

package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World",
        })
    })

    r.Run()
}