package oilPaintingChaincode

import (
	pb "github.com/hyperledger/fabric/protos/peer"
	"oilPaintingChaincode/contractapi"
)

// 油画信息链码
type OilChainCode struct {
	contractapi.Contract
}

// AddBaseInfo 画家创作基本信息上链
func (o OilChainCode) AddBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeBaseInfo(ctx, args)
}

// ReworkBaseInfo 定制发布者退回信息
func (o OilChainCode) ReworkBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeBaseInfo(ctx, args)
}

// UpdateBaseInfo 创作基本信息上链
func (o OilChainCode) UpdateBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeBaseInfo(ctx, args)
}

// ConfirmBaseInfo 确认信息上链
func (o OilChainCode) ConfirmBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeBaseInfo(ctx, args)
}
