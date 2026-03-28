package repository

import (
	"database/sql"
	"inventory/internal/models"
)

// ─────────────────────────────────────────────
//  СОТРУДНИКИ (employees)
// ─────────────────────────────────────────────

func GetAllEmployees(db *sql.DB) ([]models.Employee, error) {
	rows, err := db.Query(`SELECT id, full_name, position, department, created_at FROM employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Employee
	for rows.Next() {
		var e models.Employee
		rows.Scan(&e.ID, &e.FullName, &e.Position, &e.Department, &e.CreatedAt)
		list = append(list, e)
	}
	return list, nil
}

func GetEmployeeByID(db *sql.DB, id int) (models.Employee, error) {
	var e models.Employee
	err := db.QueryRow(`SELECT id, full_name, position, department, created_at FROM employees WHERE id=$1`, id).
		Scan(&e.ID, &e.FullName, &e.Position, &e.Department, &e.CreatedAt)
	return e, err
}

func CreateEmployee(db *sql.DB, e models.Employee) (models.Employee, error) {
	err := db.QueryRow(
		`INSERT INTO employees (full_name, position, department) VALUES ($1,$2,$3) RETURNING id, created_at`,
		e.FullName, e.Position, e.Department,
	).Scan(&e.ID, &e.CreatedAt)
	return e, err
}

func UpdateEmployee(db *sql.DB, e models.Employee) error {
	_, err := db.Exec(
		`UPDATE employees SET full_name=$1, position=$2, department=$3 WHERE id=$4`,
		e.FullName, e.Position, e.Department, e.ID,
	)
	return err
}

func DeleteEmployee(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM employees WHERE id=$1`, id)
	return err
}

// ─────────────────────────────────────────────
//  ИНСТРУМЕНТЫ (instruments)
// ─────────────────────────────────────────────

func GetAllInstruments(db *sql.DB) ([]models.Instrument, error) {
	rows, err := db.Query(`SELECT id, name, inventory_number, category, status, COALESCE(description,''), created_at FROM instruments ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Instrument
	for rows.Next() {
		var i models.Instrument
		rows.Scan(&i.ID, &i.Name, &i.InventoryNumber, &i.Category, &i.Status, &i.Description, &i.CreatedAt)
		list = append(list, i)
	}
	return list, nil
}

func GetInstrumentByID(db *sql.DB, id int) (models.Instrument, error) {
	var i models.Instrument
	err := db.QueryRow(
		`SELECT id, name, inventory_number, category, status, COALESCE(description,''), created_at FROM instruments WHERE id=$1`, id,
	).Scan(&i.ID, &i.Name, &i.InventoryNumber, &i.Category, &i.Status, &i.Description, &i.CreatedAt)
	return i, err
}

func CreateInstrument(db *sql.DB, i models.Instrument) (models.Instrument, error) {
	err := db.QueryRow(
		`INSERT INTO instruments (name, inventory_number, category, status, description) VALUES ($1,$2,$3,$4,$5) RETURNING id, created_at`,
		i.Name, i.InventoryNumber, i.Category, i.Status, i.Description,
	).Scan(&i.ID, &i.CreatedAt)
	return i, err
}

func UpdateInstrument(db *sql.DB, i models.Instrument) error {
	_, err := db.Exec(
		`UPDATE instruments SET name=$1, inventory_number=$2, category=$3, status=$4, description=$5 WHERE id=$6`,
		i.Name, i.InventoryNumber, i.Category, i.Status, i.Description, i.ID,
	)
	return err
}

func DeleteInstrument(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM instruments WHERE id=$1`, id)
	return err
}

// ─────────────────────────────────────────────
//  ВЫДАЧИ (issues)
// ─────────────────────────────────────────────

func GetAllIssues(db *sql.DB) ([]models.Issue, error) {
	rows, err := db.Query(`
		SELECT iss.id, iss.instrument_id, iss.employee_id,
		       iss.issue_date, COALESCE(iss.expected_return_date::text,''),
		       iss.return_date, COALESCE(iss.note,''), iss.created_at,
		       ins.name, emp.full_name
		FROM issues iss
		JOIN instruments ins ON ins.id = iss.instrument_id
		JOIN employees   emp ON emp.id = iss.employee_id
		ORDER BY iss.id DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []models.Issue
	for rows.Next() {
		var iss models.Issue
		rows.Scan(
			&iss.ID, &iss.InstrumentID, &iss.EmployeeID,
			&iss.IssueDate, &iss.ExpectedReturnDate,
			&iss.ReturnDate, &iss.Note, &iss.CreatedAt,
			&iss.InstrumentName, &iss.EmployeeName,
		)
		list = append(list, iss)
	}
	return list, nil
}

func CreateIssue(db *sql.DB, iss models.Issue) (models.Issue, error) {
	err := db.QueryRow(
		`INSERT INTO issues (instrument_id, employee_id, issue_date, expected_return_date, note)
		 VALUES ($1,$2,$3,NULLIF($4,'')::date,$5) RETURNING id, created_at`,
		iss.InstrumentID, iss.EmployeeID, iss.IssueDate, iss.ExpectedReturnDate, iss.Note,
	).Scan(&iss.ID, &iss.CreatedAt)
	if err != nil {
		return iss, err
	}
	// Меняем статус инструмента на "выдан"
	db.Exec(`UPDATE instruments SET status='issued' WHERE id=$1`, iss.InstrumentID)
	return iss, nil
}

func ReturnIssue(db *sql.DB, id int) error {
	var instrID int
	err := db.QueryRow(
		`UPDATE issues SET return_date=CURRENT_DATE WHERE id=$1 AND return_date IS NULL RETURNING instrument_id`, id,
	).Scan(&instrID)
	if err != nil {
		return err
	}
	// Возвращаем статус инструмента
	_, err = db.Exec(`UPDATE instruments SET status='in_stock' WHERE id=$1`, instrID)
	return err
}

func DeleteIssue(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM issues WHERE id=$1`, id)
	return err
}
