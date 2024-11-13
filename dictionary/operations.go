package dictionary

import (
	"github.com/gin-gonic/gin"
	dbUtils "gogeta.io/fante/db"
	"net/http"
)

func GetWotd(c *gin.Context) {
	db, _ := dbUtils.NewClient()

	findWotd := `query exactMatch($a: Definition) {
    me(func: eq(name, $a)) {
      name
      dgraph.type
    }
  }`
	txn := db.NewReadOnlyTxn()

	variables := make(map[string]string)
	variables["$a"] = "word"

	txn.QueryWithVars(c, findWotd, variables)

	word := DefinitionModel{2014, "Mentally disposed; willing m\u025Bk\u037B", "egq", "ready"}
	c.JSON(http.StatusOK, word)
}

func Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ready": c.Query("w")})
}
