// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

type Tag struct {
	Title string
	Color string
}

type Task struct {
	ID       string
	Title    string
	Time     int
	Complete bool
	Tags     []Tag
	BlockID  string
}

type TaskService interface {
	Create(task *Task) (*Task, error)
}

type taskService struct {
}

func NewTaskService() TaskService {
	return &taskService{}
}

func (s *taskService) Create(task *Task) (*Task, error) {
	return nil, nil
}