package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"
  "gogeta.io/fante/dictionary"
  dbUtils "gogeta.io/fante/db"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

  // db config
  db, err := sql.Open("sqlite3", "./foo.db")
  dbUtils.CheckError(err)
  defer db.Close()

	r := router.Group(os.Getenv("DICT_API"))
	r.GET(
		"/wotd",
    dictionary.GetWotd,
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

