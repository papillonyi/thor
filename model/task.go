package model

import "time"

type Task struct {
	BaseMode
	TaskName   string
	Status     string
	FinishDate *time.Time
}

func StartTask(taskName string) Task {
	task := Task{TaskName: taskName, Status: "running", FinishDate: nil}
	db.Create(&task)
	return task
}

func (t *Task) Finish() {
	now := time.Now()
	db.Model(&t).Updates(Task{Status: "finished", FinishDate: &now})
}

func (t *Task) Fail() {
	now := time.Now()
	db.Model(&t).Updates(Task{Status: "failed", FinishDate: &now})
}
