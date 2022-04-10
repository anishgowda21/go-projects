package main

import (
	"fmt"

	"github.com/anishgowda21/go-projects/go-fiber-crm/database"
	"github.com/anishgowda21/go-projects/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead/",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead/",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3","leads.db")

	if err != nil{
		panic("Cannot connect to databse..")
	}

	fmt.Println("Succesfully connected to databse...")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated...")
}

func main() {
	app:= fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(5000)
	defer database.DBConn.Close()
}