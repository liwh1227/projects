package oilPaintingChaincode

import (
	"encoding/json"
	"oilPaintingChaincode/contractapi"
	"reflect"
	"strconv"
	"time"
)

// getKey 获取请求的Key值
func getKey(in interface{}) []string {
	//1.将入参转为[]byte
	bytes, err := json.Marshal(in)
	if err != nil {
		return nil
	}

	//2.处理参数
	keyInfo := new(Key)

	//3.将string参数赋给具体的info结构体
	err = json.Unmarshal(bytes, keyInfo)
	if err != nil {
		return nil
	}
	return []string{keyInfo.DocType, keyInfo.CustomId, strconv.Itoa(keyInfo.ProductionProcess)}
}

// validator 验证confirm阶段的参数是否与链上一致
func validator(ctx contractapi.TransactionContextInterface, info interface{}, docType string) (bool, error) {
	// 1.get val by key
	bytes, err := Read(ctx, docType, getKey(info))
	if err != nil {
		return false, err
	}
	// 2.获取链码上的oilBaseInfo
	var v interface{}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return false, err
	}
	// 3.对比
	if reflect.DeepEqual(info, v) {
		return true, nil
	}
	return false, nil
}

func getOilInfo(args []string, in interface{}) error {
	// 1.校验参数
	if len(args) != 1 {
		return nil
	}
	switch in.(type) {
	case *OilBaseInfo:
		in = new(OilBaseInfo)
		err := json.Unmarshal([]byte(args[0]), in)
		if err != nil {
			return err
		}
	case *OilColorInfo:
		in = new(OilColorInfo)
		err := json.Unmarshal([]byte(args[0]), in)
		if err != nil {
			return err
		}
	case *OilShadowInfo:
		in = new(OilShadowInfo)
		err := json.Unmarshal([]byte(args[0]), in)
		if err != nil {
			return err
		}
	case *OilSketchInfo:
		in = new(OilSketchInfo)
		err := json.Unmarshal([]byte(args[0]), in)
		if err != nil {
			return err
		}
	}
	return nil
}

//
func handleChainCode(ctx contractapi.TransactionContextInterface, req *ChainCodeRequest) (*TimeValue, error) {
	// 3.1获取是否存在该key，若无,则赋值createTime
	bytes, err := Read(ctx, req.DocType, []string{req.DocType, req.CustomId, strconv.Itoa(req.ProductionProcess)})
	if err != nil {
		return nil, err
	}
	t := new(TimeValue)
	// 3.2 不存在,createTime 赋值
	if len(bytes) == 0 {
		t.CreateTime = time.Now().String()
	} else if req.ConfirmStatus == CONFIRMED {
		// 3.3.1 校验
		isEqual, err := validator(ctx, req, req.DocType)
		if err != nil {
			return nil, err
		}

		if !isEqual {
			return nil, err
		}
		// 3.3.2 校验成功后，
		t.ConfirmTime = time.Now().String()
	} else {
		t.UpdateTime = time.Now().String()
	}
	return t, nil
}
