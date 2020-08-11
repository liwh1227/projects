package contractapi

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"reflect"
	"unicode"
)

type GeneralChaincode struct {
	contract                  ContractInterface
	transactionContextHandler reflect.Type
}

func NewChaincode(inputContract ContractInterface) (*GeneralChaincode, error) {
	cc := new(GeneralChaincode)
	cc.contract = inputContract
	return cc, nil
}

// Start starts the chaincode in the fabric shim
func (cc *GeneralChaincode) Start() error {
	return shim.Start(cc)
}

func (cc *GeneralChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Chaincode Initing...")
	return shim.Success(nil)
}

// 对账本数据进行操作时被自动调用(query, invoke)
func (cc *GeneralChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()

	fmt.Println("-------------------------------------")
	fmt.Println("Invoking:", function)
	fmt.Println("args:", args)
	defer fmt.Println("-------------------------------------")
	//利用反射来选择对应功能

	fnRune := []rune(function)
	if unicode.IsLower(fnRune[0]) {
		fnRune[0] = unicode.ToUpper(fnRune[0])
		function = string(fnRune)
	}

	if exist, allFuncName := IfExistFuncName(function, cc.contract); !exist {
		var str string
		for _, name := range allFuncName {
			str = str + " " + name
		}
		return shim.Error(fmt.Sprintf("Invalid invoke function name. Expecting: %s", str))
	}

	cc.transactionContextHandler = reflect.ValueOf(cc.contract.GetTransactionContextHandler()).Elem().Type()
	ctx := reflect.New(cc.transactionContextHandler)
	ctxIface := ctx.Interface().(SettableTransactionContextInterface)
	ctxIface.SetStub(stub)

	canvasCC := reflect.ValueOf(cc.contract)
	funcSelect := canvasCC.MethodByName(function)

	//fmt.Println(args)

	reflectArgs := []reflect.Value{ctx, reflect.ValueOf(args)}
	response := funcSelect.Call(reflectArgs)

	//fmt.Println("hello")

	pe := &peer.Response{}
	bytes, _ := json.Marshal(response[0].Interface())
	json.Unmarshal(bytes, pe)
	return *pe
}

func IfExistFuncName(funcName string, contract ContractInterface) (bool, []string) {
	allFuncName := []string{}
	t := reflect.TypeOf(contract)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i).Name
		if m == funcName {
			return true, []string{}
		}
		allFuncName = append(allFuncName, m)
	}
	return false, allFuncName
}
