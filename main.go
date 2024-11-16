// entry point for app server
// serveup data from db
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gogeta.io/fante/dictionary"

	"github.com/gin-gonic/gin"
)

func main() {
  initDBFlag := flag.Bool("init", true, "Set to true to init the database")
  flag.Parse()

  dictionary.SetupDatabase(*initDBFlag)

	s := setupServer()

	s.ListenAndServe()
}

func setupServer() *http.Server {
	router := gin.Default()
	r := router.Group(os.Getenv("DICT_API"))
	r.GET(
		"/wotd",
		getWotd,
	)
	r.GET(
		"/search",
		query,
	)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func getWotd(c *gin.Context) {
	word := dictionary.DefinitionModel{Name: "ready", Description: "Mentally disposed; willing m\u025Bk\u037B", Phonetic: "re_a_dyia"}
	c.JSON(http.StatusOK, word)
}

func query(c *gin.Context) {
	def, _ := dictionary.GetDefinition(c.Query("q"))
  fmt.Printf("%v" ,def)
	c.JSON(http.StatusOK, def)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
