package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"inventory/internal/database"
	"inventory/internal/models"
	"inventory/internal/repository"
)

// ─────────────────────────────────────────────
//  СОТРУДНИКИ
// ─────────────────────────────────────────────

func GetEmployees(c *gin.Context) {
	list, err := repository.GetAllEmployees(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e, err := repository.GetEmployeeByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Сотрудник не найден"})
		return
	}
	c.JSON(http.StatusOK, e)
}

func CreateEmployee(c *gin.Context) {
	var e models.Employee
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	created, err := repository.CreateEmployee(database.DB, e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func UpdateEmployee(c *gin.Context) {
	var e models.Employee
	e.ID, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if err := repository.UpdateEmployee(database.DB, e); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Обновлено"})
}

func DeleteEmployee(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteEmployee(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}

// ─────────────────────────────────────────────
//  ИНСТРУМЕНТЫ
// ─────────────────────────────────────────────

func GetInstruments(c *gin.Context) {
	list, err := repository.GetAllInstruments(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetInstrument(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	i, err := repository.GetInstrumentByID(database.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Инструмент не найден"})
		return
	}
	c.JSON(http.StatusOK, i)
}

func CreateInstrument(c *gin.Context) {
	var i models.Instrument
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if i.Status == "" {
		i.Status = "in_stock"
	}
	created, err := repository.CreateInstrument(database.DB, i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func UpdateInstrument(c *gin.Context) {
	var i models.Instrument
	i.ID, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if err := repository.UpdateInstrument(database.DB, i); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Обновлено"})
}

func DeleteInstrument(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteInstrument(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}

// ─────────────────────────────────────────────
//  ВЫДАЧИ
// ─────────────────────────────────────────────

func GetIssues(c *gin.Context) {
	list, err := repository.GetAllIssues(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func CreateIssue(c *gin.Context) {
	var iss models.Issue
	if err := c.ShouldBindJSON(&iss); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	created, err := repository.CreateIssue(database.DB, iss)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func ReturnIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.ReturnIssue(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Возврат оформлен"})
}

func DeleteIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteIssue(database.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}
