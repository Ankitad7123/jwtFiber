package routes

import ( "gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
	"jwtFiber/controllers"
  


)

func UrlPath(app *fiber.App , db *gorm.DB){
  app.Post("/" , func(c *fiber.Ctx) error{
    return controllers.Login(c , db)
})

  app.Post("/create" , func(c *fiber.Ctx) error{
    return controllers.Create(c , db)
  })
  app.Get("/protected" , controllers.MiddleWare() , controllers.Protected)

}
