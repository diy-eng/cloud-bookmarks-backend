package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Bookmark struct {
	Id   int
	Name string
	Url  string
}

type Bookmarks struct {
}

func main() {
	bookmarks_file, err := ioutil.ReadFile("./bookmarks.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var bookmarks []Bookmark
	err = json.Unmarshal([]byte(bookmarks_file), &bookmarks)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	log.Printf("Unmarshaled:  %+v", bookmarks)

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
		idS, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Error during ID Conversion(): ", err)
		}

		for bookmark := range bookmarks {
			log.Printf("%T %v; %T %v; %T %v", bookmark, bookmark, bookmarks[bookmark].Id, bookmarks[bookmark].Id, idS, idS)
			if bookmarks[bookmark].Id == idS {
				log.Print("same")
				c.JSON(200, bookmarks[bookmark])
				return
			}
		}
		c.String(404, "Error: Id not found")
	})

	r.POST("/bookmark/:id", func(c *gin.Context) {
		id := c.Param("id")
		idS, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Error during ID Conversion(): ", err)
		}

		// Determine if id already in use
		for bookmark := range bookmarks {
			if bookmarks[bookmark].Id == idS {
				c.String(500, "Error: Id already in use")
				return
			}
		}

		bookmarks := append(bookmarks, Bookmark{Id: idS, Name: "Test", Url: "Test"})
		c.String(200, "Sucess: added")

		log.Println(bookmarks)
	})

	r.Run()
}
