package validators

import (
	"bytes"
	"fmt"
	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
	"github.com/zilinjak/oas-proxy/internal/logging"
	"io"
	"net/http"
	"os"
)

type OAS3Validator struct {
	Validator validator.Validator
}

func NewOAS3Validator(oasPath string) *OAS3Validator {
	oasFile, err := os.ReadFile(oasPath)
	if err != nil {
		panic("Failed to read file " + oasPath + "error: " + err.Error())
	}
	document, err := libopenapi.NewDocument(oasFile)
	if err != nil {
		panic("Failed to parse OAS file: " + err.Error())
	}
	highLevelValidator, validatorErrs := validator.NewValidator(document)
	if len(validatorErrs) > 0 {
		for _, err := range validatorErrs {
			logging.Logger.Error(err.Error())
		}
		panic("Failed to validate OAS file")
	}
	logging.Logger.Info("OAS File loaded, validator created")
	return &OAS3Validator{
		Validator: highLevelValidator,
	}
}

func (v *OAS3Validator) Validate(request *http.Request, requestData []byte, response *http.Response) {
	oasResponse := &http.Response{
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       response.Body,
	}
	oasRequest := &http.Request{
		Method: request.Method,
		URL:    request.URL,
		Header: request.Header,
		Body:   io.NopCloser(bytes.NewBuffer(requestData)),
	}

	fmt.Println("Validating request")
	requestValid, errs := v.Validator.ValidateHttpRequestResponse(oasRequest, oasResponse)
	if requestValid {
		logging.Logger.Info("Request is valid")
	} else {
		logging.Logger.Info("Request is invalid")
		for _, err := range errs {
			logging.Logger.Info(err.Error())
		}
	}
}
