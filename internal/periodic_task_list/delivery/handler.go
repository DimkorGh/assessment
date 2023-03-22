package delivery

import (
	"github.com/kataras/iris/v12"

	"assessment/internal/core/logging"
	"assessment/internal/periodic_task_list/service"
	"assessment/internal/utils/parser"
	"assessment/internal/utils/validators"
)

type PtListHandlerInt interface {
	GetPtList(ctx iris.Context)
}

type PtListHandler struct {
	urlParamsParser parser.UrlParamsParserInt
	service         service.PtListServiceInt
	logger          logging.LoggerInt
}

func NewPtListHandler(
	urlParamsParser parser.UrlParamsParserInt,
	service service.PtListServiceInt,
	logger logging.LoggerInt,
) *PtListHandler {
	return &PtListHandler{
		urlParamsParser: urlParamsParser,
		service:         service,
		logger:          logger,
	}
}

func (th *PtListHandler) GetPtList(ctx iris.Context) {
	var (
		getPtListReq getPtListRequest
		ptList       []string
	)

	err := th.urlParamsParser.ParseUrlParams(ctx.Request(), &getPtListReq)
	if err != nil {
		th.generateResponse(ctx, ptList, err)

		return
	}

	ptList, err = th.service.GetTimestampsList(
		getPtListReq.Tz,
		getPtListReq.Period,
		getPtListReq.T1,
		getPtListReq.T2,
	)

	th.generateResponse(ctx, ptList, err)
}

func (th *PtListHandler) generateResponse(ctx iris.Context, ptList []string, err error) {
	var respErr error

	ctx.StatusCode(th.getResponseStatusCode(err))

	if err != nil {
		th.logError(err)

		respErr = ctx.JSON(th.getErrorResponseBody(err))
	} else {
		respErr = ctx.JSON(ptList)
	}

	if respErr != nil {
		th.logger.Errorf("Error while generating JSON response: %s: ", err.Error())
	}
}

func (th *PtListHandler) getResponseStatusCode(err error) int {
	if err == nil {
		return iris.StatusOK
	}

	if th.isInternalError(err) {
		return iris.StatusInternalServerError
	}

	return iris.StatusBadRequest
}

func (th *PtListHandler) logError(err error) {
	if err == nil {
		return
	}

	if th.isInternalError(err) {
		th.logger.Error(err.Error())
	}
}

func (th *PtListHandler) getErrorResponseBody(err error) getPtListErrorResponse {
	var resp getPtListErrorResponse

	resp.Status = "error"
	resp.Description = err.Error()

	return resp
}

func (th *PtListHandler) isInternalError(err error) bool {
	switch err.(type) {
	case *parser.UrlParamParserError,
		*validators.StructValidatorError,
		*service.ConvertingInputDataError:
		return false
	default:
		return true
	}
}
