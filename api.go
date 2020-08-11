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

// AddBaseInfo 画家创作基本信息上链
func (o OilChainCode) AddColorInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeColorInfo(ctx, args)
}

// ReworkColorInfo 定制发布者退回信息
func (o OilChainCode) ReworkColorInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeColorInfo(ctx, args)
}

// UpdateColorInfo 创作基本信息上链
func (o OilChainCode) UpdateColorInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeColorInfo(ctx, args)
}

// ConfirmColorInfo 确认信息上链
func (o OilChainCode) ConfirmColorInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeColorInfo(ctx, args)
}

// AddBaseInfo 画家创作基本信息上链
func (o OilChainCode) AddShadowInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeShadowInfo(ctx, args)
}

// ReworkShadowInfo 定制发布者退回信息
func (o OilChainCode) ReworkShadowInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeShadowInfo(ctx, args)
}

// UpdateShadowInfo 创作基本信息上链
func (o OilChainCode) UpdateShadowInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeShadowInfo(ctx, args)
}

// ConfirmShadowInfo 确认信息上链
func (o OilChainCode) ConfirmShadowInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeShadowInfo(ctx, args)
}

// AddSketchInfo 画家创作基本信息上链
func (o OilChainCode) AddSketchInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeSketchInfo(ctx, args)
}

// ReworkSketchInfo 定制发布者退回信息
func (o OilChainCode) ReworkSketchInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeSketchInfo(ctx, args)
}

// UpdateSketchInfo 创作基本信息上链
func (o OilChainCode) UpdateSketchInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeSketchInfo(ctx, args)
}

// ConfirmSketchInfo 确认信息上链
func (o OilChainCode) ConfirmSketchInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	return InvokeSketchInfo(ctx, args)
}
