package service

import (
	"board/app/interface/response"
	"net/http"

	"github.com/gin-gonic/gin/binding"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Validator struct{}

func (v Validator) Validate(obj any) error {
	// バリデーション対象がからの場合にはリターン
	if obj == nil {
		return nil
	}

	//validation.Validatableインターフェースにバインド
	//Validateメソッドを構造体のメソッドに持っておく必要がある。
	val, err := obj.(validation.Validatable)
	if !err {
		return nil
	}
	//バリデーションを実行する
	if err := val.Validate(); err != nil {
		if verr, ok := err.(validation.Errors); ok {
			var params []response.InvalidParams
			for key, val := range verr {
				params = append(params, response.InvalidParams{Name: key, Reason: val.Error()})
			}
			return v.newValidationError(params, err)
		}
		return v.newServerError(err)
	}
	return nil
}

// newValidationError creates an error that wraps a Internal Error.
func (v Validator) newServerError(err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "SERVER_ERROR", Title: "unexpected errors"},
		status:   http.StatusInternalServerError,
		err:      err,
	}
}

// newValidationError creates an error that wraps a Validation Error.
func (v Validator) newValidationError(params []response.InvalidParams, err error) error {
	return ValidationError{
		response: response.ValidationError{Type: "VALIDATION_ERROR", Title: "Your request parameters didn't validate.", Pramas: params},
		status:   http.StatusBadRequest,
		err:      err,
	}
}

// validation/error.go
// ValidationError represents validation error.
type ValidationError struct {
	response response.ValidationError
	status   int
	err      error
}

// Error returns the error string of Errors.
func (v ValidationError) Error() string {
	return v.err.Error()
}

// Response returns the Response.
func (v ValidationError) Response() response.ValidationError {
	return v.response
}

// Status returns the Status Code.
func (v ValidationError) Status() int {
	return v.status
}

type ozzoValidator struct {
	validator *Validator
}

func NewOzzoValidator() binding.StructValidator {
	return &ozzoValidator{validator: &Validator{}}
}

func (v *ozzoValidator) ValidateStruct(obj any) error {
	return v.validator.Validate(obj)
}

func (v *ozzoValidator) Engine() any {
	return v.validator
}
