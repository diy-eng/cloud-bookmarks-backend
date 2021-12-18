package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	bookmarks := [][3]string{
		{"1", "Google", "https://google.com"},
		{"2", "Example", "https://example.com"},
	}

	// REST API
	r := gin.Default()

	// Serve React App at root
	r.Use(static.Serve("/", static.LocalFile("./react-app", true)))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/bookmarks", func(c *gin.Context) {
		c.JSON(200, bookmarks)
	})

	r.GET("/bookmark/:id", func(c *gin.Context) {
		id := c.Param("id")

		for bookmark := range bookmarks {
			if bookmarks[bookmark][0] == id {
				c.JSON(200, bookmarks[bookmark])
			}
		}
	})

	r.POST("/bookmark/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Determine if id already in use
		for bookmark := range bookmarks {
			if bookmarks[bookmark][0] == id {
				c.String(500, "Error: Id already in use")
			}
		}

		bookmarks := append(bookmarks, [3]string{"100", "Example", "https://example.com"})

		fmt.Println(bookmarks)
	})

	r.Run()
}
