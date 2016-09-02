package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/martinlindhe/winservice"
	"golang.org/x/sys/windows/svc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	winServiceCommand = kingpin.Arg("command", "Install, remove, start, stop.").Required().String()
	serviceName       = kingpin.Arg("service-name", "Service name.").Required().String()
	serviceDesc       = kingpin.Flag("description", "Service description (required for 'install').").Short('d').String()
	exePath           = kingpin.Flag("exe", "Executable (required for 'install').").String()
)

func main() {
	var err error

	kingpin.Parse()

	switch *winServiceCommand {
	case "install":
		cleanExePath := filepath.Clean(*exePath)
		if !exists(cleanExePath) {
			log.Fatalf("exe not found: %v", err)
		}
		err = winservice.Install(*serviceName, *serviceDesc, cleanExePath)
	case "remove":
		err = winservice.Remove(*serviceName)
	case "start":
		err = winservice.Start(*serviceName)
	case "stop":
		err = winservice.Control(*serviceName, svc.Stop, svc.Stopped)
	default:
		log.Fatalf("unknown svc command: %v", *winServiceCommand)
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", *winServiceCommand, *serviceName, err)
	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
