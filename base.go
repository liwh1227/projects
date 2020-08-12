package oilPaintingChaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
	"oilPaintingChaincode/contractapi"
)

func BadRequest(msg string) peer.Response {
	return peer.Response{
		Status:  400,
		Message: msg,
	}
}

func Save(ctx contractapi.TransactionContextInterface, objectType string, attributes []string, payload interface{}) ([]byte, error) {
	key, err := ctx.GetStub().CreateCompositeKey(objectType, attributes)
	if err != nil {
		return nil, err
	}

	bytes, ok := payload.([]byte)
	if !ok {
		bytes, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	err = ctx.GetStub().PutState(key, bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func Read(ctx contractapi.TransactionContextInterface, objectType string, attributes []string) ([]byte, error) {
	key, err := ctx.GetStub().CreateCompositeKey(objectType, attributes)
	if err != nil {
		fmt.Println("CreateCompositeKey")
		return nil, err
	}
	res, err := ctx.GetStub().GetState(key)
	if err != nil {
		fmt.Println("GetState")
		return nil, err
	}
	return res, nil
}

func response(ctx contractapi.TransactionContextInterface, payload interface{}) peer.Response {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling response", payload, err)
		return ctx.BadRequest(err.Error())
	}

	fmt.Println("Success response", payload)
	return ctx.Response(payloadBytes)
}
