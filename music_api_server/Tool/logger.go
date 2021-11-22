package Tool

import (
	logs "github.com/sirupsen/logrus"
	"os"
)

var Logger *logs.Logger

func init(){
	Logger = logs.StandardLogger()
	Logger.SetLevel(logs.DebugLevel)
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logs.JSONFormatter{})
}