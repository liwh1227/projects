package oilPaintingChaincode

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"oilPaintingChaincode/contractapi"
)

func InvokeBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilBaseInfo)
	err := getOilInfo(args, info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	req := &ChainCodeRequest{
		DocType:           info.DocType,
		CustomId:          info.CustomId,
		ProductionProcess: info.ProductionProcess,
		ConfirmStatus:     info.ConfirmStatus,
	}
	// 2.存放链码前进行逻辑判断
	t, err := handleChainCode(ctx, req)
	if err != nil {
		return BadRequest("handle ChainCode error")
	}
	// 2.1 赋值时间
	info.CreateTime = t.CreateTime
	info.UpdateTime = t.UpdateTime
	info.ConfirmTime = t.ConfirmTime

	// 3.上链操作
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}

func InvokeColorInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilColorInfo)
	err := getOilInfo(args, info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	req := &ChainCodeRequest{
		DocType:           info.DocType,
		CustomId:          info.CustomId,
		ProductionProcess: info.ProductionProcess,
		ConfirmStatus:     info.ConfirmStatus,
	}
	// 2.存放链码前进行逻辑判断
	t, err := handleChainCode(ctx, req)
	if err != nil {
		return BadRequest("handle ChainCode error")
	}
	// 2.1 赋值时间
	info.CreateTime = t.CreateTime
	info.UpdateTime = t.UpdateTime
	info.ConfirmTime = t.ConfirmTime

	// 3.上链操作
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}

func InvokeShadowInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilShadowInfo)
	err := getOilInfo(args, info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	req := &ChainCodeRequest{
		DocType:           info.DocType,
		CustomId:          info.CustomId,
		ProductionProcess: info.ProductionProcess,
		ConfirmStatus:     info.ConfirmStatus,
	}
	// 2.存放链码前进行逻辑判断
	t, err := handleChainCode(ctx, req)
	if err != nil {
		return BadRequest("handle ChainCode error")
	}
	// 2.1 赋值时间
	info.CreateTime = t.CreateTime
	info.UpdateTime = t.UpdateTime
	info.ConfirmTime = t.ConfirmTime

	// 3.上链操作
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}

func InvokeSketchInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilSketchInfo)
	err := getOilInfo(args, info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	req := &ChainCodeRequest{
		DocType:           info.DocType,
		CustomId:          info.CustomId,
		ProductionProcess: info.ProductionProcess,
		ConfirmStatus:     info.ConfirmStatus,
	}
	// 2.存放链码前进行逻辑判断
	t, err := handleChainCode(ctx, req)
	if err != nil {
		return BadRequest("handle ChainCode error")
	}
	// 2.1 赋值时间
	info.CreateTime = t.CreateTime
	info.UpdateTime = t.UpdateTime
	info.ConfirmTime = t.ConfirmTime

	// 3.上链操作
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}
