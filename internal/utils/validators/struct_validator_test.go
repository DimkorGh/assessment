package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	type args struct {
		checkStruct interface{}
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "two required fields none missing",
			args: args{
				checkStruct: struct {
					Field1 string `json:"field1" validate:"required"`
					Field2 string `json:"field2" validate:"required"`
				}{
					Field1: "XXX1111",
					Field2: "1",
				},
			},
			want: nil,
		},
		{
			name: "required field missing",
			args: args{
				checkStruct: struct {
					Field1 string `json:"field1" validate:"required"`
					Field2 string `json:"field2" validate:"required"`
				}{
					Field1: "",
					Field2: "1",
				},
			},
			want: &StructValidatorError{
				ErrorMessage: "Key: 'Field1' Error:Field validation for 'Field1' failed on the 'required' tag",
			},
		},
		{
			name: "no required fields",
			args: args{
				checkStruct: struct {
					Field1 string `json:"field1"`
					Field2 string `json:"field2"`
				}{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strValidator := NewStructValidator(validator.New())
			err := strValidator.Validate(tt.args.checkStruct)

			assert.Equal(t, tt.want, err)
		})
	}
}

func TestStructValidatorError(t *testing.T) {
	tests := []struct {
		name                 string
		structValidatorError StructValidatorError
		want                 string
	}{
		{
			name:                 "test custom error function",
			structValidatorError: StructValidatorError{"i am an error message"},
			want:                 "i am an error message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.structValidatorError.Error()
			assert.Equal(t, tt.want, err)
		})
	}
}
