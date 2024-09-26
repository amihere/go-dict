package dictionary

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWotd(c *gin.Context) {
	word := DefinitionModel{2014, "Mentally disposed; willing m\u025Bk\u037B", "egq", "ready"}
	c.JSON(http.StatusOK, word)
}

func Query(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ready": c.Query("w")})
}
