package repositories

import (
	"database/sql"
	"example.com/student-api/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil { return nil, err }
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}
	return students, nil
}

func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	var s models.Student
	err := r.DB.QueryRow("SELECT id, name, major, gpa FROM students WHERE id = ?", id).Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	return &s, err
}

func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.DB.Exec("INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)", s.Id, s.Name, s.Major, s.GPA)
	return err
}