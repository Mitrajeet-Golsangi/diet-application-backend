package logging

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Initialize the logging for the application
func InitializeLogger() {
	// Create a new log file
	f, err := os.OpenFile("gin.log", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	
	if err != nil {
		log.Fatalf("Failed to create the log file: %v", err)
	}

	// Set the gin log output to write in both the file and the console
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Log output to the file and the console
	log.SetOutput(gin.DefaultWriter)

	// Set the gin log format
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
}

func CloseLogger() {
	// Close the log file
	_ = gin.DefaultWriter.(*os.File).Close()
}