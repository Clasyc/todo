package model

import "time"

// Task represents single task
type Task struct {
	ID    int
	Title string
	Done  bool
	CreatedAt *time.Time
}

func NewTask(title string, done bool) *Task {
	return &Task{
		Title: title,
		Done:  done,
	}
}

func (t *Task) Create() error {
	return svc.DB.
		Model(&Task{}).Create(t).Error
}

func (t *Task) Toggle(b bool) error {
	return svc.DB.Model(&Task{}). Where("id = ?", t.ID).Updates(map[string]interface{}{"done": b}).Error
}

func (t *Task) Delete() error {
	return svc.DB.Model(&Task{}).Delete(t).Error
}

func GetTasks() ([]*Task, error) {
	tasks := make([]*Task, 0)
	err := svc.DB.Model(&Task{}).Find(&tasks).Error
	return tasks, err
}

func GetTask(id int) (*Task, error) {
	task := &Task{}
	err := svc.DB.Model(task).First(&task, "id = ?", id).Error
	return task, err
}