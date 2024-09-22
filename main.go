package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Word struct {
	Meaning string
	Basic   string
	Id      int32
}

func main() {
	router := gin.Default()

	r := router.Group("/pusha")
	r.GET(
		"/wotd",
		func(c *gin.Context) {
			word := Word{"Mentally disposed; willing m\u025Bk\u037B", "ready", 123}
			c.JSON(http.StatusOK, word)
		},
	)
	r.GET(
		"/q", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ready": c.Query("w")})
		},
	)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
