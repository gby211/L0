package http

import (
	"L0/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	server.initRouter()
	return server
}

func (server *Server) initRouter() {
	server.router = gin.Default()
	server.router.Use(Middleware())
	server.router.GET("/order", server.getOrder)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (server *Server) Run(addr ...string) {
	err := server.router.Run(addr...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *Server) getOrder(c *gin.Context) {
	orderUid := c.Query("order_uid")
	c.JSON(http.StatusOK, service.GetById(orderUid))
}
