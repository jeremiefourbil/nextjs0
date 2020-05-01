package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual, "Expected response code %d. Got %d\n", expected, actual)
}

func TestGetToto(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/toto", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	body := response.Body.String()
	assert.Equal(t, "toto",body,"Expected toto. Got %s", body)
}

func TestNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/foobar", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestSum(t *testing.T) {
	total := Sum(4, 4)
	assert.Equal(t, 8, total, "Sum was incorrect, got: %d, want: %d.", total, 8)
}