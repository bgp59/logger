// Reference code on how to use logrusx

package main

import (
	"flag"
	"os"

	"github.com/bgp59/logger/logrusx"

	// The app logger will be defined in some package:
	"github.com/bgp59/logger/example/pkg"
)

var mainLogger = pkg.RootLogger.NewCompLogger("main")

func main() {
	useJsonFlag := flag.Bool(
		"use-json",
		logrusx.LOGGER_CONFIG_USE_JSON_DEFAULT,
		"Structure the logged record in JSON",
	)

	levelFlag := flag.String(
		"level",
		logrusx.LOGGER_CONFIG_LEVEL_DEFAULT,
		"Log level name: info, warn... etc",
	)

	disableSrcFileFlag := flag.Bool(
		"disable-src-file",
		logrusx.LOGGER_CONFIG_DISBALE_SRC_FILE_DEFAULT,
		"Disable the reporting of the source file:line# info",
	)

	logFileFlag := flag.String(
		"log-file",
		logrusx.LOGGER_CONFIG_LOG_FILE_DEFAULT,
		"Log to a file or use stdout/stderr",
	)

	logFileMaxSizeMbFlag := flag.Int(
		"log-file-max-size-mb",
		logrusx.LOGGER_CONFIG_LOG_FILE_MAX_SIZE_MB_DEFAULT,
		"Log file max size, in MB, before rotation, use 0 to disable",
	)

	logFileMaxBackupNumFlag := flag.Int(
		"log-file-max-backup-num",
		logrusx.LOGGER_CONFIG_LOG_FILE_MAX_BACKUP_NUM_DEFAULT,
		"How many older log files to keep upon rotation",
	)

	flag.Parse()

	cfg := &logrusx.LoggerConfig{
		UseJson:             *useJsonFlag,
		Level:               *levelFlag,
		DisableSrcFile:      *disableSrcFileFlag,
		LogFile:             *logFileFlag,
		LogFileMaxSizeMB:    *logFileMaxSizeMbFlag,
		LogFileMaxBackupNum: *logFileMaxBackupNumFlag,
	}

	if err := pkg.RootLogger.SetLogger(cfg); err != nil {
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
