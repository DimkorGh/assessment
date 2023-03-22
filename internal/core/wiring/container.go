package wiring

import (
	"github.com/go-playground/validator/v10"
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
	Server *server.Server
	Logger *logging.Logger
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) InitializeDependencies(address string, port int) {
	vpr := viper.New()
	cfg := config.NewConfig(vpr)

	cfg.Load("config")

	cfg.SetServerAddress(address)
	cfg.SetServerPort(port)

	logger := logging.NewLogger(cfg)
	logger.Initialize()

	irisApp := iris.New()

	structValidator := validators.NewStructValidator(validator.New())
	urlParamsParser := parser.NewUrlParamsParser(structValidator)

	ptListDomain := domain.NewPtListDomain()
	ptListService := service.NewPtListService(ptListDomain)
	ptListHandler := delivery.NewPtListHandler(urlParamsParser, ptListService, logger)

	c.Server = server.NewServer(irisApp, logger, cfg, ptListHandler)
}
