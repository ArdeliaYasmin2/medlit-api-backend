package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"medlit-api-backend/api/models"
)

type MedicineInput struct {
	GenericName      string `json:"generic_name" binding:"required"`
	PhotoURL         string `json:"photo_url" binding:"required"`
	Purpose          string `json:"purpose" binding:"required"`
	SideEffects      string `json:"side_effects" binding:"required"`
	Contraindication string `json:"contraindication" binding:"required"`
	Dosage           string `json:"dosage" binding:"required"`
	Ingredients      string `json:"ingredients" binding:"required"`
}

type ControllerRepo struct {
	Repo models.RepoInterface
}

func NewController(repo models.RepoInterface) *ControllerRepo {
	return &ControllerRepo{Repo: repo}
}

func (m *ControllerRepo) AddMedicine(c *gin.Context) {
	var input MedicineInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	medicine := models.Medicine{}
	medicine.GenericName = input.GenericName
	medicine.PhotoURL = input.PhotoURL
	medicine.Purpose = input.Purpose
	medicine.SideEffects = input.SideEffects
	medicine.Contraindication = input.Contraindication
	medicine.Dosage = input.Dosage
	medicine.Ingredients = input.Ingredients

	err := m.Repo.CreateMedicine(medicine)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   "false",
		"message": "Medicine created",
	})
}

func (m *ControllerRepo) GetAllMedicine(c *gin.Context) {
	medicine := m.Repo.GetAllMedicine()

	c.JSON(http.StatusOK, gin.H{
		"error":        "false",
		"message":      "Medicine list",
		"medicineList": medicine,
	})
}

func (m *ControllerRepo) GetMedicineByID(c *gin.Context) {
	id := c.Param("id")

	medicine, err := m.Repo.GetMedicineByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":    "false",
		"message":  "Medicine found",
		"medicine": medicine,
	})
}

func (m *ControllerRepo) GetMedicineByQuery(c *gin.Context) {
	generic_name := c.Query("generic_name")

	medicine, err := m.Repo.GetMedicineByQuery(generic_name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":        "false",
		"message":      "Medicine found",
		"medicineList": medicine,
	})
}
