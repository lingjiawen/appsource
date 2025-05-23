// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"mangosmithy/api/v1/system"
	"mangosmithy/internal/app/system/model/do"
)

type (
	ISysJobLog interface {
		Add(ctx context.Context, data *do.SysJobLog) (err error)
		List(ctx context.Context, req *system.SysJobLogListReq) (listRes *system.SysJobLogListRes, err error)
		Delete(ctx context.Context, logIds []uint64) (err error)
	}
)

var (
	localSysJobLog ISysJobLog
)

func SysJobLog() ISysJobLog {
	if localSysJobLog == nil {
		panic("implement not found for interface ISysJobLog, forgot register?")
	}
	return localSysJobLog
}

func RegisterSysJobLog(i ISysJobLog) {
	localSysJobLog = i
}
