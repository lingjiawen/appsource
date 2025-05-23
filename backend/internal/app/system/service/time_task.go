// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"mangosmithy/internal/app/system/model"
)

type (
	ITaskList interface {
		AddTask(task *model.TimeTask)
		GetByName(funcName string) *model.TimeTask
		EditParams(funcName string, params []string)
	}
)

var (
	localTaskList ITaskList
)

func TaskList() ITaskList {
	if localTaskList == nil {
		panic("implement not found for interface ITaskList, forgot register?")
	}
	return localTaskList
}

func RegisterTaskList(i ITaskList) {
	localTaskList = i
}
