package oilPaintingChaincode

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"oilPaintingChaincode/contractapi"
	"strconv"
	"time"
)

func InvokeBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilBaseInfo)
	err := getOilInfo(args, info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, info.DocType, []string{info.DocType, info.CustomId, strconv.Itoa(info.ProductionProcess)})
	if err != nil {
		return BadRequest("Read Key has error: %v\n")
	}

	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		info.CreateTime = time.Now().String()
	} else {
		// 3.3 存在，判断时[update]还是[confirm]
		if info.ConfirmStatus == CONFIRMED {
			// 3.3.1 校验
			isEqual, err := validator(ctx, info)
			if err != nil {
				return BadRequest("validator info error.\n")
			}

			if !isEqual {
				return BadRequest("Please confirm that the information is consistent with the chain code\n")
			}
			// 3.3.2 校验成功后，
			info.ConfirmTime = time.Now().String()
		} else {
			info.UpdateTime = time.Now().String()
		}
	}

	// 4.存储到db中
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}

func handle() {

}
