package demo

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func MakeLog() {
	// logrus.JSONFormatter
	log.SetFormatter(&log.TextFormatter{ForceColors: true, DisableColors: false, FullTimestamp: true})
	// Only log the debug severity or above. Default to info
	log.SetLevel(log.DebugLevel)
	// Level from low to high
	log.Debugln("I'm a debug log ...")
	log.Infoln("I'm an info log ...")
	log.Warnln("I'm a warning log ...")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	var logForError = log.New()
	logForError.Out = os.Stderr
	logForError.Errorln("I'm an error log ...")

	// Calls panic() after logging
	log.Panicln("I'm a panic log ...")
	// Calls os.Exit(1) after logging
	log.Fatalln("I'm a fatal log ...")
}
