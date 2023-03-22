package parser

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"assessment/internal/utils/validators"
)

func TestParseUrlParams(t *testing.T) {
	type request struct {
		FirstField  string `schema:"firstField" validate:"required"`
		SecondField string `schema:"secondField" validate:"required"`
	}

	type args struct {
		httpRequest *http.Request
		parseData   interface{}
	}
	tests := []struct {
		name   string
		args   args
		field1 string
		field2 string
		err    error
	}{
		{
			name: "invalid field name",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com/test?firstInvalidField=field1",
					nil,
				),
				parseData: &request{},
			},
			err: &UrlParamParserError{ErrorMessage: `schema: invalid path "firstInvalidField"`},
		},
		{
			name: "required url param is missing",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com/test?firstField=field1",
					nil,
				),
				parseData: &request{},
			},
			err: &validators.StructValidatorError{ErrorMessage: "Key: 'request.SecondField' Error:Field validation for 'SecondField' failed on the 'required' tag"},
		},
		{
			name: "valid url params parsing",
			args: args{
				httpRequest: httptest.NewRequest(
					"POST",
					"http://test.com/test?firstField=field1&secondField=field2",
					nil,
				),
				parseData: &request{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strValidator := validators.NewStructValidator(validator.New())
			parser := NewUrlParamsParser(strValidator)

			err := parser.ParseUrlParams(tt.args.httpRequest, tt.args.parseData)

			assert.Equal(t, tt.err, err)
		})
	}
}
