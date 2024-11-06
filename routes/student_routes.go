package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prithvi009/fealtyx/handlers"
)

func Routers(app *fiber.App) {
	app.Get("/students", handlers.GetAllStudents)
	app.Post("/students", handlers.CreateStudent)
	app.Get("/students/:id", handlers.GetStudentById)
	app.Put("/students/:id", handlers.UpdateStudentById)
	app.Delete("/students/:id", handlers.DeleteStudentById)
	app.Get("/students/:id/summary", handlers.GetStudentSummary)
}
