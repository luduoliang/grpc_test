package controllers

import (
	"github.com/astaxie/beego/logs"
	"golang.org/x/net/context"
	"grpc_test/models"
	"grpc_test/proto"
)

//转换详情
func TransPddSessions(info *models.PddSessions) *proto.PddSessions {
	out := new(proto.PddSessions)
	out.Id = int32(info.ID)
	out.TaokeID = int32(info.TaokeID)
	out.ScreenName = info.ScreenName
	out.OpenId = info.OpenId
	out.Token = info.Token
	out.RefreshToken = info.RefreshToken
	out.IsDefault = int32(info.IsDefault)
	out.ExpiredAt = TransTime(info.ExpiredAt)
	out.RefreshExpiredAt = TransTime(info.RefreshExpiredAt)
	out.CreatedAt = TransTime(info.CreatedAt)
	out.UpdatedAt = TransTime(info.UpdatedAt)
	return out
}

//转换详情
func UnTransPddSessions(info *proto.PddSessions) *models.PddSessions {
	out := new(models.PddSessions)
	out.ID = uint(info.Id)
	out.TaokeID = uint(info.TaokeID)
	out.ScreenName = info.ScreenName
	out.OpenId = info.OpenId
	out.Token = info.Token
	out.RefreshToken = info.RefreshToken
	out.IsDefault = uint8(info.IsDefault)
	out.ExpiredAt = UnTransTime(info.ExpiredAt)
	out.RefreshExpiredAt = UnTransTime(info.RefreshExpiredAt)
	out.CreatedAt = UnTransTime(info.CreatedAt)
	out.UpdatedAt = UnTransTime(info.UpdatedAt)
	return out
}

//添加
func (s *Server) AddPddSessions(ctx context.Context, in *proto.RequestAddPddSessions) (*proto.ResponseAddPddSessions, error) {
	logs.Info("Server.AddPddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.CreatePddSessions(info)
	if err != nil {
		return nil, err
	}
	return &proto.ResponseAddPddSessions{PddSessions: TransPddSessions(info)}, nil
}

//更新
func (s *Server) UpdatePddSessions(ctx context.Context, in *proto.RequestUpdatePddSessions) (*proto.ResponseUpdatePddSessions, error) {
	logs.Info("Server.UpdatePddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.UpdatePddSessions(info)
	if err != nil {
		return nil, err
	}
	return &proto.ResponseUpdatePddSessions{PddSessions: TransPddSessions(info)}, nil
}

//删除
func (s *Server) DeletePddSessions(ctx context.Context, in *proto.RequestDeletePddSessions) (*proto.ResponseDeletePddSessions, error) {
	logs.Info("Server.DeletePddSessions")
	err := models.DeletePddSessions(uint(in.Id))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseDeletePddSessions{}, nil
}

//详情
func (s *Server) GetPddSessionsInfo(ctx context.Context, in *proto.RequestGetPddSessionsInfo) (*proto.ResponseGetPddSessionsInfo, error) {
	logs.Info("Server.GetPddSessionsInfo")
	info := models.GetPddSessionsInfo(uint(in.Id))
	returnInfo := TransPddSessions(info)
	return &proto.ResponseGetPddSessionsInfo{PddSessions: returnInfo}, nil
}

//列表
func (s *Server) GetPddSessionsList(ctx context.Context, in *proto.RequestGetPddSessionsList) (*proto.ResponseGetPddSessionsList, error) {
	logs.Info("Server.GetPddSessionsList")
	list, total := models.GetPddSessionsList(int(in.Page), int(in.PerPage))
	returnList := []*proto.PddSessions{}
	for _, val := range list {
		returnList = append(returnList, TransPddSessions(val))
	}
	return &proto.ResponseGetPddSessionsList{PddSessions: returnList, Total: int32(total)}, nil
}
