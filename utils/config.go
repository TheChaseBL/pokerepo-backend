package utils

import (
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct {
	MongoUrl string `json:"mongo-url"`
}

func GetConfig() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
