package main

import (
	"log"
"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
"jwtFiber/models"
	"jwtFiber/routes"
)


func main(){
  app := fiber.New()
  dsn :=""
  db , err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})
  if err != nil {
    log.Fatal(err)
  }
 if err =  db.AutoMigrate(&models.Users12{}); err != nil {
    log.Fatal(err)
  }
  routes.UrlPath(app , db)
  app.Listen(":8080")









}
