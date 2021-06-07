package handler

import (
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric/common/flogging"
)

type Data struct{
	Key string `json:"key"`
	Value []byte `json:"value"`
}

var logger = flogging.MustGetLogger("handler")

func QueryData(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Infof("Get value by key: %s", args[0])
	return stub.GetState(args[0])
}

func SaveData(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	logger.Infof("Enter .....%s", function)
	data := &Data{}
	txId := stub.GetTxID()
	err := json.Unmarshal([]byte(args[0]), data)
	if err != nil {
		return nil, err
	}
	//policy.TxID = txId
	//bytes, err := json.Marshal(policy)
	//if err != nil {
	//	return nil, err
	//}
	err = stub.PutState(txId, data.Value)
	if err != nil {
		return nil, err
	}
	err = stub.PutState(data.Key, data.Value)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
