package model

import "time"

type Task struct {
	TaskName        string    `json:"taskname"`
	TaskDescription string    `json:"taskdescription"`
	TaskDeadline    time.Time `json:"taskdeadline"`
	TaskStatus      string    `json:"taskstatus"`
}

type TaskResponse struct {
	ID              uint      `json:"id"`
	TaskName        string    `json:"taskname"`
	TaskDescription string    `json:"taskdescription"`
	TaskDeadline    time.Time `json:"taskdeadline"`
	TaskStatus      string    `json:"taskstatus"`
	UserID          uint      `json:"user_id"`
}

type TaskStatusResp struct {
	ID         uint   `json:"id"`
	TasKStatus string `json:"taskStatus"`
}
