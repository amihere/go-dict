// entry point for app server
// serveup data from db
package main

import (
	"net/http"
	"os"
	"time"

	dbUtils "gogeta.io/fante/db"
	"gogeta.io/fante/dictionary"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

  // db config
  _, cancel:= dbUtils.NewClient()
  defer cancel()

	r := router.Group(os.Getenv("DICT_API"))
	r.GET(
		"/wotd",
    dictionary.GetWotd,
	)
	r.GET(
		"/q",
    dictionary.Query,
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

