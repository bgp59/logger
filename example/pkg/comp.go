package pkg

// Define a component sub-logger:
var compLogger = RootLogger.NewCompLogger("comp")

// Pretend function(s) to invoke from main tpo log some messages:
func Func1() {
	compLogger.Error("Error")
	compLogger.Warn("Warn")
	compLogger.Info("Info")
	compLogger.Debug("Debug")
	compLogger.Trace("Trace")
}

func Func2() {
	compLogger.Error("Error")
	compLogger.Warn("Warn")
	compLogger.Info("Info")
	compLogger.Debug("Debug")
	compLogger.Trace("Trace")
}
