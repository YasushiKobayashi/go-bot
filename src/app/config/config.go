package config

import (
	"os"
)

// server config
const PORT = "7000"

var HOST = "0.0.0.0:"

// setting project
var GOPATH = os.Getenv("GOPATH")
var PROJECT_PATH = GOPATH + "/src/app/"
