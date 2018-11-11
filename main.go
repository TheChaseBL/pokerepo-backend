package main

import (
	"context"
	"fmt"
	"github.com/TheChaseBL/pokerepo-backend/database"
	"github.com/TheChaseBL/pokerepo-backend/entity"
)

func main() {
	db := database.NewConnection()
	charizard := entity.Pokemon{"Charizard"}
	db.AddEntry(charizard)
	fmt.Println(db.ReadAll())
	db.Client.Disconnect(context.Background())
}