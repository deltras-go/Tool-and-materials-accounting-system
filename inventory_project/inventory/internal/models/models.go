package models

import "time"

// Таблица 1: Сотрудники
type Employee struct {
	ID         int       `json:"id"`
	FullName   string    `json:"full_name"`
	Position   string    `json:"position"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}

// Таблица 2: Инструменты
type Instrument struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	InventoryNumber string    `json:"inventory_number"`
	Category        string    `json:"category"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
}

// Таблица 3: Выдачи
type Issue struct {
	ID                 int        `json:"id"`
	InstrumentID       int        `json:"instrument_id"`
	EmployeeID         int        `json:"employee_id"`
	IssueDate          string     `json:"issue_date"`
	ExpectedReturnDate string     `json:"expected_return_date"`
	ReturnDate         *string    `json:"return_date"`
	Note               string     `json:"note"`
	CreatedAt          time.Time  `json:"created_at"`
	// Для отображения (JOIN)
	InstrumentName string `json:"instrument_name,omitempty"`
	EmployeeName   string `json:"employee_name,omitempty"`
}
