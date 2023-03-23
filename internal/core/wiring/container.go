package wiring

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"

	"assessment/internal/core/config"
	"assessment/internal/core/logging"
	"assessment/internal/core/server"
	"assessment/internal/periodic_task_list/delivery"
	"assessment/internal/periodic_task_list/domain"
	"assessment/internal/periodic_task_list/service"
	"assessment/internal/utils/parser"
	"assessment/internal/utils/validators"
)

type Container struct {
	Server        *server.Server
	Logger        *logging.Logger
	serverAddress string
	serverPort    int
}

func NewContainer(serverAddress string, serverPort int) *Container {
	return &Container{
		serverAddress: serverAddress,
		serverPort:    serverPort,
	}
}

func (c *Container) InitializeDependencies() {
	vpr := viper.New()
	cfg := config.NewConfig(vpr)

	cfg.Load("config")

	cfg.SetServerAddress(c.serverAddress)
	cfg.SetServerPort(c.serverPort)

	logger := logging.NewLogger(cfg)
	logger.Initialize()

	irisApp := iris.New()

	structValidator := validators.NewStructValidator(validator.New())
	urlParamsParser := parser.NewUrlParamsParser(structValidator, schema.NewDecoder())

	ptListDomain := domain.NewPtListDomain()
	ptListService := service.NewPtListService(ptListDomain)
	ptListHandler := delivery.NewPtListHandler(urlParamsParser, ptListService, logger)

	c.Server = server.NewServer(irisApp, logger, cfg, ptListHandler)
}
