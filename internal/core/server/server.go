package server

import (
	"context"
	"strconv"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

	"assessment/internal/core/config"
	"assessment/internal/core/logging"
	"assessment/internal/periodic_task_list/delivery"
)

type Server struct {
	irisApp       *iris.Application
	logger        *logging.Logger
	cfg           *config.Config
	ptListHandler delivery.PtListHandlerInt
}

func NewServer(
	irisApp *iris.Application,
	logger *logging.Logger,
	cfg *config.Config,
	ptListHandler delivery.PtListHandlerInt,
) *Server {
	return &Server{
		irisApp:       irisApp,
		logger:        logger,
		cfg:           cfg,
		ptListHandler: ptListHandler,
	}
}

func (s *Server) Start() {
	s.irisApp.Use(iris.Compression)

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{iris.MethodGet},
	})
	s.irisApp.Use(crs)

	s.irisApp.Get("/ptlist", s.ptListHandler.GetPtList)

	go func() {
		err := s.irisApp.Listen(s.cfg.Server.Address+":"+strconv.Itoa(s.cfg.Server.Port), iris.WithoutServerError(iris.ErrServerClosed))
		if err != nil {
			s.logger.Errorf("Error while starting server: %s", err.Error())
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := s.irisApp.Shutdown(ctx)
	if err != nil {
		s.logger.Errorf("Error while shutdown server: %s", err.Error())
	}
}
