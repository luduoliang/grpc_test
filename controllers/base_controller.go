package controllers

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

// 业务实现方法的容器
type Server struct{}

//转换时间
func TransTime(in *time.Time) *timestamp.Timestamp {
	out := new(timestamp.Timestamp)
	inTime := in.Unix()
	out.Seconds = inTime
	return out
}

//转换时间
func UnTransTime(in *timestamp.Timestamp) *time.Time {
	out := time.Unix(int64(in.Seconds), 0)
	return &out
}
