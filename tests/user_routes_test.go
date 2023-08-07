package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWhenCreateUser(t *testing.T) {
	requestUrl := "http://localhost:8080/api/user"
	requestBodyJson := `{
		"name": "test",
		"email": "test@gmail.com",
		"date_of_birth": "1990-01-01",
		"admin": false,
		"password": "123456"
	}`

	requestBodyIo := io.Reader(strings.NewReader(requestBodyJson))

	clientRequest, err := http.NewRequest(http.MethodPost, requestUrl, requestBodyIo)

	if err != nil {
		t.Error(err)
	}
	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	serverResponse, err := http.DefaultClient.Do(clientRequest)

	if err != nil {
		t.Error(err)
	}

	serverData, err := io.ReadAll(serverResponse.Body)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(serverData))

	assert.Equal(t, 200, serverResponse.StatusCode)

}
