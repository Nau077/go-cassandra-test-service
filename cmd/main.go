package main

import (
	"fmt"
	"net/http"

	routes "github.com/Nau077/cassandra-golang-sv/internal/presentation/routes"

	"github.com/gin-gonic/gin"
	// "os"
	// "regexp"
	// "github.com/gin-contrib/logger"
	// "github.com/gin-gonic/gin"
	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

type routesM struct {
	router *gin.Engine
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	port := "8090"

	routes.NewPostHTTPHandler(r)
	r.Run(port)
	if errHTTP := http.ListenAndServe(":"+port, r); errHTTP != nil {
		// log.Error().Msg(errHTTP.Error())
		fmt.Print(errHTTP.Error())
	}
}
