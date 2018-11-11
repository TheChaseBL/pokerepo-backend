package main

import (
	"github.com/TheChaseBL/pokerepo-backend/api"
	"github.com/TheChaseBL/pokerepo-backend/database"
)

func main() {
	database.Conn = database.NewConnection()
	api.StartAPI()
}