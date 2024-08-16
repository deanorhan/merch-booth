package main

import (
	"log"
	"net/http"

	"github.com/deanorhan/merch-booth/libs/db-package"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type Merch struct {
	Status int8    `json:"status"`
	Title  string  `json:"title"`
	Price  float32 `json:"price"`
}

func GetMerch(c *gin.Context) {
	dab := GetDBConnection()

	queries := db.New(dab)

	merch, err := queries.ListMerch(c)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, merch)
}

func PostMerch(c *gin.Context) {
	var merch Merch

	if err := c.BindJSON(&merch); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	dab := GetDBConnection()

	queries := db.New(dab)

	m, err := queries.CreateMerch(c, db.CreateMerchParams{Vendor: 0, Status: 0, Title: merch.Title, Price: float64(merch.Price)})
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, Merch{Status: int8(m.Status), Title: m.Title, Price: float32(m.Price)})
}
