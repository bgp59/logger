// Reference code on how to use logger

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/bgp59/logger"

	"github.com/bgp59/logger/example/pkg"
)

func loadConfig(configFile string) (*logger.LoggerConfig, error) {
	cfg := logger.DefaultLoggerConfig()
	if configFile == "" {
		return cfg, nil
	}

	f, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("file: %q: %v", configFile, err)
	}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, fmt.Errorf("file: %q: %v", configFile, err)
	}

	return cfg, nil
}

var mainLogger = pkg.RootLogger.NewCompLogger("main")

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		pkg.RootLogger.Errorf("%v\n", err)
		os.Exit(1)
	}

	err = pkg.RootLogger.SetLogger(cfg)
	if err != nil {
		mainLogger.Errorf("%v\n", err)
		os.Exit(1)
	}

	// Log some messages:
	mainLogger.Error("Error")
	mainLogger.Warn("Warn")
	mainLogger.Info("Info")
	mainLogger.Debug("Debug")
	mainLogger.Trace("Trace")

	// Invoke functions that use the logger:
	pkg.Func1()
	pkg.Func2()
}
