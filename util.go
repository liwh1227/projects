package oilPaintingChaincode

import (
	"encoding/json"
	"oilPaintingChaincode/contractapi"
	"reflect"
	"strconv"
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
func validator(ctx contractapi.TransactionContextInterface, baseInfo *OilBaseInfo) (bool, error) {
	// 1.get val by key
	bytes, err := Read(ctx, baseInfo.DocType, getKey(baseInfo))
	if err != nil {
		return false, err
	}
	// 2.获取链码上的oilBaseInfo
	chaincodeBaseInfo := new(OilBaseInfo)
	err = json.Unmarshal(bytes, chaincodeBaseInfo)
	if err != nil {
		return false, err
	}
	// 3.对比
	if reflect.DeepEqual(baseInfo, chaincodeBaseInfo) {
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
	case *OilShadowInfo:
	case *OilSketchInfo:
	}
	return nil
}
