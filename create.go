package oilPaintingChaincode

import (
	"log"
	"oilPaintingChaincode/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(new(OilChainCode))
	if err != nil {
		log.Fatalf("Error create ChainCode: %v\n", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		log.Fatalf("Error starting ChainCode: %v\n", err)
	}
}
