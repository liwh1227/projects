package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"oilPaintingChaincode"
	"oilPaintingChaincode/contractapi"
	"testing"
)

func checkInvokeAdd(t *testing.T, stub *shim.MockStub, args []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("AddBaseInfo"), args})
	if res.Status != shim.OK {
		fmt.Println("AddBaseInfo", "failed", res.Message)
		t.FailNow()
	}

	if res.Payload == nil {
		fmt.Println("AddBaseInfo", "failed to insert value")
		t.FailNow()
	}
	fmt.Println("The txId is ", res.Payload)
}

func checkInvokeConfirm(t *testing.T, stub *shim.MockStub, args []byte) {
	res := stub.MockInvoke("1", [][]byte{[]byte("ConfirmBaseInfo"), args})
	if res.Status != shim.OK {
		fmt.Println("AddBaseInfo", "failed", res.Message)
		t.FailNow()
	}

	if res.Payload == nil {
		fmt.Println("AddBaseInfo", "failed to insert value")
		t.FailNow()
	}
	fmt.Println("The txId is ", res.Payload)
}

//func checkState(t *testing.T, stub *shim.MockStub, orderId string, value []byte){
//	bytes := stub.State[orderId]
//	if bytes == nil {
//		fmt.Println("State", orderId, "failed to get value")
//		t.FailNow()
//	}
//
//	getOrderEvaluation   := &oilPaintingChaincode.OilBaseInfo{}
//	imputOrderEvaluation := &oilPaintingChaincode.OilBaseInfo{}
//
//	json.Unmarshal(bytes,getOrderEvaluation)
//	json.Unmarshal(value,imputOrderEvaluation)
//
//	if ! ifOrderEvaluationEqual(getOrderEvaluation,imputOrderEvaluation) {
//		t.FailNow()
//	}
//}
//

const (
	Step1 = 0
	Step2 = 1
	Step3 = 2
	Step4 = 3
)

func TestOilChainCode_AddBaseInfo(t *testing.T) {
	scc, err := contractapi.NewChaincode(new(OilChainCode))
	if err != nil {
		fmt.Printf("Error create OilPaintingChainCode: %s", err.Error())
		return
	}
	stub := shim.NewMockStub("oilPainting", scc)

	// 测试1：step0 confirmed
	order1 := &oilPaintingChaincode.OilBaseInfo{
		DocType:           "oil",
		CustomId:          "1231321",
		WorksName:         "test name",
		ProductionProcess: Step1,
		ConfirmStatus:     oilPaintingChaincode.TCB,
		CreateTime:        "",
		UpdateTime:        "",
		ConfirmTime:       "",
		Painter:           "test",
		ManufactureDate:   "test",
		Copyright:         "test",
		Canvas:            "test",
		Frame:             "test",
		Pigment:           "test",
		CanvasImage:       "test",
		CanvasImageHash:   "test",
		FrameImage:        "test",
		FrameImageHash:    "test",
		PigmentImage:      "test",
		PigmentImageHash:  "test",
	}
	order, err := json.Marshal(order1)
	if err != nil {
		fmt.Println("incorrect order1 convert to []byte ")
	}
	checkInvokeAdd(t, stub, order)
	fmt.Println("----------------checkInvoke Updated---------------------")
	order0 := &oilPaintingChaincode.OilBaseInfo{
		DocType:           "oil",
		CustomId:          "1231321",
		WorksName:         "test name",
		ProductionProcess: Step1,
		ConfirmStatus:     oilPaintingChaincode.TCB,
		CreateTime:        "",
		UpdateTime:        "",
		ConfirmTime:       "",
		Painter:           "test",
		ManufactureDate:   "test",
		Copyright:         "test",
		Canvas:            "test",
		Frame:             "test",
		Pigment:           "test",
		CanvasImage:       "test",
		CanvasImageHash:   "test",
		FrameImage:        "test",
		FrameImageHash:    "test",
		PigmentImage:      "test",
		PigmentImageHash:  "test",
	}
	order4, err := json.Marshal(order0)
	if err != nil {
		fmt.Println("incorrect order0 convert to []byte ")
	}
	checkInvokeConfirm(t, stub, order4)
	fmt.Println("----------------checkInvoke Confirmed---------------------")
	order2 := &oilPaintingChaincode.OilBaseInfo{
		DocType:           "oil",
		CustomId:          "1231321",
		WorksName:         "test name",
		ProductionProcess: Step1,
		ConfirmStatus:     oilPaintingChaincode.CONFIRMED,
		CreateTime:        "",
		UpdateTime:        "",
		ConfirmTime:       "",
		Painter:           "test",
		ManufactureDate:   "test",
		Copyright:         "test",
		Canvas:            "test",
		Frame:             "test",
		Pigment:           "test",
		CanvasImage:       "test",
		CanvasImageHash:   "test",
		FrameImage:        "test",
		FrameImageHash:    "test",
		PigmentImage:      "test",
		PigmentImageHash:  "test",
	}
	order3, err := json.Marshal(order2)
	if err != nil {
		fmt.Println("incorrect order2 convert to []byte ")
	}
	checkInvokeConfirm(t, stub, order3)
}
