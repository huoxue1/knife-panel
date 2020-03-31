package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"html/template"
	"knife-panel/webtty"
	"log"
	"net/http"
	noesctmpl "text/template"
	"time"

	"github.com/gorilla/websocket"
)

// Server provides a webtty HTTP endpoint.
type Server struct {
	factory Factory
	options *Options

	upgrader      *websocket.Upgrader
	indexTemplate *template.Template
	titleTemplate *noesctmpl.Template
}

// New creates a new instance of Server.
// Server will use the New() of the factory provided to handle each request.
func New(factory Factory, options *Options) (*Server, error) {

	return &Server{
		factory: factory,
		options: options,

		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    webtty.Protocols,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}, nil
}

// Run starts the main process of the Server.
// The cancelation of ctx will shutdown the server immediately with aborting
// existing connections. Use WithGracefullContext() to support gracefull shutdown.
func (server *Server) Run(ctx context.Context, ginContext *gin.Context, options ...RunOption) error {

	counter := newCounter(time.Duration(server.options.Timeout) * time.Second)

	server.generateHandleWS(ctx, counter).ServeHTTP(ginContext.Writer, ginContext.Request)

	conn := counter.count()
	if conn > 0 {
		log.Printf("Waiting for %d connections to be closed", conn)
	}
	counter.wait()

	return nil
}
