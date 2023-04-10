package libs

import (
	"net/http"

	routes "github.com/Nau077/cassandra-golang-sv/internal/presentation/routes"
	"github.com/Nau077/cassandra-golang-sv/internal/services/post"
	"github.com/gin-gonic/gin"
)

type RouterClient struct {
	router      *gin.Engine
	port        string
	postService *post.Service
}

func NewRouterClient(port string, s *post.Service) *RouterClient {
	return &RouterClient{
		router: gin.Default(),
		port:   port,
	}
}

func (r *RouterClient) Init() error {
	gin.SetMode(gin.ReleaseMode)
	port := r.port

	routes.NewPostHTTPHandler(r.router, r.postService)
	r.router.Run(port)
	if errHTTP := http.ListenAndServe(":"+port, r.router); errHTTP != nil {
		return errHTTP
	}

	return nil
}
