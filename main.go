package main

import (
	"github.com/alpakih/simpel-api/database"
	"github.com/alpakih/simpel-api/models"
	"github.com/alpakih/simpel-api/router"
)

func main() {

	db := database.GetConnection()
	db.AutoMigrate(&models.Order{}, &models.OrderItem{})

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := router.StartRouter(db)

	r.Run(":3000")

}
