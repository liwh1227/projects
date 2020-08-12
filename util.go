package oilPaintingChaincode

import (
	"encoding/json"
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
	return []string{keyInfo.CustomId, strconv.Itoa(keyInfo.ProductionProcess)}
}

// validator 验证confirm阶段的参数是否与链上一致
//func validator(ctx contractapi.TransactionContextInterface, info interface{}, docType string) (bool, error) {
//	// 1.get val by key
//	key := getKey(info)
//	fmt.Printf("[validator]  key is %v\n\n",key)
//	bytes, err := Read(ctx, docType, key)
//	if err != nil {
//		return false, err
//	}
//	fmt.Printf("[validator] len(bytes) is %v\n",len(bytes))
//	// 2.获取链码上的oilBaseInfo
//	var v interface{}
//	switch info.(type) {
//	case *OilBaseInfo:
//		v = new(OilBaseInfo)
//	}
//	fmt.Printf("[validator] v is %v\n",v)
//	err = json.Unmarshal(bytes, v)
//	if err != nil {
//		fmt.Printf("json.umarshal has error %v\n",err)
//		return false, err
//	}
//	fmt.Printf("validator is %#v\n",v)
//	// 3.对比
//	if reflect.DeepEqual(info, v) {
//		return true, nil
//	}
//	return false, nil
//}

////	handleChainCode 处理要上链码的参数信息
//func handleChainCode(ctx contractapi.TransactionContextInterface, req *ChainCodeRequest) (*TimeValue, error) {
//	// 3.1获取是否存在该key，若无,则赋值createTime
//	fmt.Printf("start Read-----%#v\n",req)
//	bytes, err := Read(ctx, req.DocType, []string{req.CustomId, strconv.Itoa(req.ProductionProcess)})
//	if err != nil {
//		fmt.Println("Read error")
//		return nil, err
//	}
//	t := new(TimeValue)
//	// 3.2 不存在,createTime 赋值
//	fmt.Printf("len(bytes) is %v\n",len(bytes))
//	if len(bytes) == 0 {
//		t.CreateTime = time.Now().String()
//	} else if req.ConfirmStatus == CONFIRMED {
//		// 3.3.1 校验
//		isEqual, err := validator(ctx, req, req.DocType)
//		if err != nil {
//			return nil, err
//		}
//		fmt.Println("validator")
//		if !isEqual {
//			fmt.Println("isEqual")
//			return nil, err
//		}
//		// 3.3.2 校验成功后，
//		t.ConfirmTime = time.Now().String()
//	} else {
//		t.UpdateTime = time.Now().String()
//	}
//	return t, nil
//}

// convertToConfirmBaseStruct 为了进行struct的比较，去掉影响的字段time、status
func convertToConfirmBaseStruct(info *OilBaseInfo) *OilConfirmBaseInfo {
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil
	}

	v := new(OilConfirmBaseInfo)
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return nil
	}
	return v
}

// convertToConfirmBaseStruct 为了进行struct的比较，去掉影响的字段time、status
func convertToConfirmColorStruct(info *OilColorInfo) *OilConfirmColorInfo {
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil
	}

	v := new(OilConfirmColorInfo)
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return nil
	}
	return v
}

// convertToConfirmBaseStruct 为了进行struct的比较，去掉影响的字段time、status
func convertToConfirmSketchStruct(info *OilSketchInfo) *OilConfirmSketchInfo {
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil
	}

	v := new(OilConfirmSketchInfo)
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return nil
	}
	return v
}

func convertToConfirmShadowStruct(info *OilShadowInfo) *OilConfirmShadowInfo {
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil
	}

	v := new(OilConfirmShadowInfo)
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return nil
	}
	return v
}
