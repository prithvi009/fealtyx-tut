package handlers

import (
	"strconv"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/prithvi009/fealtyx/models"
	"github.com/prithvi009/fealtyx/utils"
)

func GetAllStudents(c *fiber.Ctx) error {
	var students []models.Student
	for _, student := range models.Students {
		students = append(students, student)
	}

	return c.JSON(students)
}

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)

	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Input"})
	}

	newId := len(models.Students) + 1

	student.Id = newId
	models.Students[newId] = *student

	return c.Status(fiber.StatusCreated).JSON(student)

}

func GetStudentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	student, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	return c.JSON(student)
}

func UpdateStudentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	student, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	updatedStudent := new(models.Student)
	if err := c.BodyParser(updatedStudent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	student.Name = updatedStudent.Name
	student.Age = updatedStudent.Age
	student.Email = updatedStudent.Email
	models.Students[id] = student

	return c.JSON(student)
}

func DeleteStudentById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	_, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	delete(models.Students, id)
	return c.SendStatus(fiber.StatusNoContent)
}

func GetStudentSummary(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	student, exists := models.Students[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	prompt := fmt.Sprintf("Provide a brief summary for the following student:\nName: %s\nAge: %d\nEmail: %s", student.Name, student.Age, student.Email)

	summary, err := utils.LlamaAPI(prompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate summary"})
	}

	return c.JSON(fiber.Map{"summary": summary})
}
