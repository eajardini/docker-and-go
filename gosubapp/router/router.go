package router

import (
	godocker "godocker/golibs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNum(c *gin.Context) {
	numAdded := godocker.MostraNum(9)
	// fmt.Println("Valor de num adicionado:", numAdded)
	c.JSON(http.StatusOK, gin.H{
		"message": numAdded,
	})
}
