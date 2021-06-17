package routes_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/mstzn/modanisa_backend/server"
)

func SetupSuite() {

	serverReady := make(chan bool)

	server := server.Server{
		Port:        10000,
		ServerReady: serverReady,
	}

	go server.Start()
	<-serverReady
}

func TestCreateTodo(t *testing.T) {

	SetupSuite()

	reqStr := `{"title":"sample"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/todos", 10000), strings.NewReader(reqStr))
	if err != nil {
		t.Error("Request instance failed")
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Error("Request failed")
	}
	if http.StatusOK != response.StatusCode {
		t.Error("Status codes not match")
	}

	byteBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error("Read body failed")
	}

	if !strings.Contains(strings.Trim(string(byteBody), "\n"), `"title":"sample"`) {
		t.Error("Response does not match")
	}

	response.Body.Close()
}
