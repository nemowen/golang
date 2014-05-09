package utils

import (
	"encoding/json"
	"gotest/rbg/config"
	"io/ioutil"
	"log"
	"os"
	//"sync"
)

//var once sync.Once

func init() {
	loadConfig()
}
