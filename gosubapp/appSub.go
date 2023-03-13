package main

import (
	router "godocker/gosubapp/router" // "godocker/gosuapp/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", router.AddNum)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
