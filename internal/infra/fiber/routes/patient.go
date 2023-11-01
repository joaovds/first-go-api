package routes

import (
  "github.com/gofiber/fiber/v2"
  patient_controllers "github.com/joaovds/first-go-api/internal/controllers"
)

func handlePatientRoutes(router fiber.Router) {
  patientRouter := router.Group("/patients")

  patientRouter.Get("/", patient_controllers.GetAllPatients)
  patientRouter.Get("/:patientId", patient_controllers.GetPatientById)
  patientRouter.Delete("/:patientId", patient_controllers.DeletePatient)
}


