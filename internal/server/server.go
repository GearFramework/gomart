package server

import (
	"github.com/GearFramework/gomart/internal/pkg/alog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type HttpServer struct {
	HTTP   *http.Server
	Router *gin.Engine
	Logger *zap.SugaredLogger
	Config *Config
}

type MiddlewareFunc func() gin.HandlerFunc

func NewServer(conf *Config) *HttpServer {
	gin.SetMode(gin.ReleaseMode)
	return &HttpServer{
		Config: conf,
		Logger: alog.NewLogger("Server " + conf.Addr),
		Router: gin.New(),
	}
}

func (serv *HttpServer) SetMiddleware(mw MiddlewareFunc) *HttpServer {
	serv.Router.Use(mw())
	return serv
}

func (serv *HttpServer) Init(initRoutes func()) error {
	initRoutes()
	return nil
}

func (serv *HttpServer) Up() error {
	serv.HTTP = &http.Server{
		Addr:    serv.Config.Addr,
		Handler: serv.Router,
	}
	serv.Logger.Infof("start server at the: %s", serv.Config.Addr)
	err := serv.HTTP.ListenAndServe()
	if err != nil {
		serv.Logger.Errorf("failed: %s", err.Error())
		return err
	}
	return nil
}
