package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	internalerrors "projeto/internal/internal-errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_Return_InternalError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInteral
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.ErrInteral.Error())
}
func Test_HandlerError_Return_DomainError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 400, errors.New("Domain")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}
func Test_HandlerError_Return_Obj_Status(t *testing.T) {
	assert := assert.New(t)
	type BodyForTest struct {
		Id int
	}
	expectedObject := BodyForTest{Id: 1}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return expectedObject, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	returnedObject := BodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &returnedObject)
	assert.Equal(expectedObject, returnedObject)
}
