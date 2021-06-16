package routes_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/labstack/echo"
	"github.com/mstzn/modanisa_backend/server"
	"github.com/stretchr/testify/suite"
)

type toDoTestingSuite struct {
	suite.Suite
	port int
}

func TestToDoTestingSuite(t *testing.T) {
	suite.Run(t, &toDoTestingSuite{})
}

func (s *toDoTestingSuite) SetupSuite() {

	serverReady := make(chan bool)

	server := server.Server{
		Port:        3000,
		ServerReady: serverReady,
	}

	go server.Start()
	<-serverReady
}

func (s *toDoTestingSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}
func (s *toDoTestingSuite) Test_EndToEnd_Todo() {
	reqStr := `{"title":"sample"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/todos", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"status":200,"message":"Success","data":{"id":1}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}
