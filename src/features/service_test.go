package features

import (
	"fmt"
	"github.com/cucumber/godog"
	"net/http"
	"os/exec"
	"syscall"
	"time"
)

var product interface{}
var response *http.Response
var err error

func iHaveNoProductCriteria() error {
	if product != nil {
		return fmt.Errorf("product should be nil")
	}

	return nil
}


func iCallTheProductEndpoint() error {
	return godog.ErrPending
}

func iShouldReceiveABadRequestMessage() error {
	return godog.ErrPending
}

func iHaveAValidProductCriteria() error {
	return godog.ErrPending
}

func iShouldReceiveAJsonResponseWithACorrespondingMessage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have no product criteria$`, iHaveNoProductCriteria)
	s.Step(`^I call the product endpoint$`, iCallTheProductEndpoint)
	s.Step(`^I should receive a bad request message$`, iShouldReceiveABadRequestMessage)
	s.Step(`^I have a valid product criteria$`, iHaveAValidProductCriteria)
	s.Step(`^I should receive a json response with a corresponding message$`, iShouldReceiveAJsonResponseWithACorrespondingMessage)

	s.BeforeScenario(func(interface{}) {
		startServer()
	})

	s.AfterScenario(func(interface{}, error) {
		server.Process.Signal(syscall.SIGINT)
	})

}

var server *exec.Cmd


func startServer() {
	server = exec.Command("go", "build", "../main.go")
	server.Run()

	server = exec.Command("./main")
	go server.Run()

	time.Sleep(3 * time.Second)
	fmt.Printf("Server running with pid: %v", server.Process.Pid)
}