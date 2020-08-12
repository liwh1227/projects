package oilPaintingChaincode

import (
	"encoding/json"
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"oilPaintingChaincode/contractapi"
	"reflect"
	"strconv"
	"time"
)

func InvokeBaseInfo(ctx contractapi.TransactionContextInterface, args []string) pb.Response {
	// 1.处理参数
	info := new(OilBaseInfo)
	err := json.Unmarshal([]byte(args[0]), info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}

	// 2.存放链码前进行逻辑判断,这里面涉及到结构体deepEqual
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, info.DocType, []string{info.CustomId, strconv.Itoa(info.ProductionProcess)})
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		info.CreateTime = time.Now().String()
	} else if info.ConfirmStatus == CONFIRMED {
		// 3.3.1 校验
		// 执行confirm操作时的info结构体
		baseInfo := convertToConfirmBaseStruct(info)
		bytes, err := Read(ctx, info.DocType, getKey(info))
		if err != nil {
			return BadRequest("Read from chainCode has some error!\n")
		}
		v := new(OilBaseInfo)
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return BadRequest(err.Error())
		}
		// 链上的数据
		baseInfoOnChainCode := convertToConfirmBaseStruct(info)
		if !reflect.DeepEqual(baseInfo, baseInfoOnChainCode) {
			return BadRequest("Please confrim data keep with chainCode")
		}
		// 3.3.2 校验成功后，
		info.ConfirmTime = time.Now().String()
	} else {
		info.UpdateTime = time.Now().String()
	}

	fmt.Printf("\n2.-------------------%#v------------------\n", info)

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
	err := json.Unmarshal([]byte(args[0]), info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	fmt.Printf("\n1.-------------------%#v------------------\n", info)

	// 2.存放链码前进行逻辑判断,这里面涉及到结构体deepEqual
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, info.DocType, []string{info.CustomId, strconv.Itoa(info.ProductionProcess)})
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		info.CreateTime = time.Now().String()
	} else if info.ConfirmStatus == CONFIRMED {
		// 3.3.1 校验
		// 执行confirm操作时的info结构体
		baseInfo := convertToConfirmColorStruct(info)
		bytes, err := Read(ctx, info.DocType, getKey(info))
		if err != nil {
			return BadRequest("Read from chainCode has some error!\n")
		}
		v := new(OilBaseInfo)
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return BadRequest(err.Error())
		}
		// 链上的数据
		baseInfoOnChainCode := convertToConfirmColorStruct(info)
		if !reflect.DeepEqual(baseInfo, baseInfoOnChainCode) {
			return BadRequest("Please confrim data keep with chainCode")
		}
		// 3.3.2 校验成功后，
		info.ConfirmTime = time.Now().String()
	} else {
		info.UpdateTime = time.Now().String()
	}

	fmt.Printf("\n2.-------------------%#v------------------\n", info)

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
	err := json.Unmarshal([]byte(args[0]), info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	fmt.Printf("\n1.-------------------%#v------------------\n", info)

	// 2.存放链码前进行逻辑判断,这里面涉及到结构体deepEqual
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, info.DocType, []string{info.CustomId, strconv.Itoa(info.ProductionProcess)})
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		info.CreateTime = time.Now().String()
	} else if info.ConfirmStatus == CONFIRMED {
		// 3.3.1 校验
		// 执行confirm操作时的info结构体
		baseInfo := convertToConfirmSketchStruct(info)
		bytes, err := Read(ctx, info.DocType, getKey(info))
		if err != nil {
			return BadRequest("Read from chainCode has some error!\n")
		}
		v := new(OilBaseInfo)
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return BadRequest(err.Error())
		}
		// 链上的数据
		baseInfoOnChainCode := convertToConfirmSketchStruct(info)
		if !reflect.DeepEqual(baseInfo, baseInfoOnChainCode) {
			return BadRequest("Please confrim data keep with chainCode")
		}
		// 3.3.2 校验成功后，
		info.ConfirmTime = time.Now().String()
	} else {
		info.UpdateTime = time.Now().String()
	}

	fmt.Printf("\n2.-------------------%#v------------------\n", info)

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
	err := json.Unmarshal([]byte(args[0]), info)
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 2.存放链码前进行逻辑判断,这里面涉及到结构体deepEqual
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, info.DocType, []string{info.CustomId, strconv.Itoa(info.ProductionProcess)})
	if err != nil {
		return BadRequest(fmt.Sprintf("Not valid file structure: " + err.Error()))
	}
	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		info.CreateTime = time.Now().String()
	} else if info.ConfirmStatus == CONFIRMED {
		// 3.3.1 校验
		// 执行confirm操作时的info结构体
		baseInfo := convertToConfirmShadowStruct(info)
		bytes, err := Read(ctx, info.DocType, getKey(info))
		if err != nil {
			return BadRequest("Read from chainCode has some error!\n")
		}
		v := new(OilBaseInfo)
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return BadRequest(err.Error())
		}
		// 链上的数据
		baseInfoOnChainCode := convertToConfirmShadowStruct(info)
		if !reflect.DeepEqual(baseInfo, baseInfoOnChainCode) {
			return BadRequest("Please confrim data keep with chainCode")
		}
		// 3.3.2 校验成功后，
		info.ConfirmTime = time.Now().String()
	} else {
		info.UpdateTime = time.Now().String()
	}

	fmt.Printf("\n2.-------------------%#v------------------\n", info)

	// 3.上链操作
	_, err = Save(ctx, info.DocType, getKey(info), info)
	if err != nil {
		return BadRequest("Put Key has some error.\n")
	}
	return response(ctx, ctx.GetStub().GetTxID())
}
