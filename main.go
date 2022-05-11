package main

//#include "break.c"
import "C"

import (
	"github.com/bugsnag/bugsnag-go/v2"
	"time"
)

func main() {

	// Configure bugsnag to send to our "Test Go" project locally
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          "ea455bb65bcb31772a8ace38e9fdd999", // other
		ReleaseStage:    "testing",
		ProjectPackages: []string{"main"},
		Endpoints:       bugsnag.Endpoints{Notify: "http://localhost:8000/", Sessions: "http://localhost:10000/"},
	})

	time.Sleep(2 * time.Second)

	C.Fault()

	time.Sleep(2 * time.Second)
}
