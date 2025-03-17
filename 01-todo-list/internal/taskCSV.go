package internal

import (
	"encoding/csv"
	"os"
	"fmt"
	"time"
	"todo-list/models"
	"strconv"
	"github.com/google/uuid"
)

type TaskRepositoryCSV struct {
	filename string
}

func NewTaskRepositoryCSV(filename string) *TaskRepositoryCSV {
	return &TaskRepositoryCSV{filename: filename}
}

func (r *TaskRepositoryCSV) AddTask(description string) error {
	file, err := os.OpenFile(r.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// New ID
	newID := uuid.New().String()

	task := models.Task{
		ID: newID,
		Description: description,
		CreatedAt: time.Now(),
		IsComplete: false,
	}

	// Write to CSV
	record := []string{
		task.ID,
		task.Description,
		task.CreatedAt.Format(time.RFC3339),
		strconv.FormatBool(task.IsComplete),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	writer.Flush()
	return nil
}

func (r *TaskRepositoryCSV) ListTasks(showAll bool) ([]models.Task, error) {
	// Đọc tất cả records hiện tại
	records, err := r.readAllRecords()
	if err != nil {
		return nil, err
	}
	
	var tasks []models.Task
	for _, record := range records {
		if len(record) < 4 {
			continue
		}
		
		createdAt, _ := time.Parse(time.RFC3339, record[2])
		isComplete, _ := strconv.ParseBool(record[3])

		tasks = append(tasks, models.Task{
			ID: record[0],
			Description: record[1],
			CreatedAt: createdAt,
			IsComplete: isComplete,
		})
	}

	return tasks, nil
}

func (r *TaskRepositoryCSV) CompleteTask(id int) error {
	// Đọc tất cả records hiện tại
	records, err := r.readAllRecords()
	if err != nil {
		return err
	}

	// Cập nhật record cần thay đổi
	found := false
	for i, record := range records {
		if len(record) < 4 {
			continue
		}

		if record[0] == strconv.Itoa(id) {
			record[3] = strconv.FormatBool(true)
			records[i] = record
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with id %d not found", id)
	}

	// Ghi lại toàn bộ file
	return r.writeAllRecords(records)
}

func (r *TaskRepositoryCSV) readAllRecords() ([][]string, error) {
	file, err := os.OpenFile(r.filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func (r *TaskRepositoryCSV) writeAllRecords(records [][]string) error {
	file, err := os.OpenFile(r.filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	return writer.WriteAll(records)
}

