package main

import (
	"encoding/json"
	"fmt"
	// "bytes"
	"time"
	// "strconv"


	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type ChainCode struct {
}

// // 키 구조체
// type User struct {
// 	Phone   string `json:"phone"`
// 	Truck Truck `json:"truck"`
// }

// 키 구조체
type Truck struct {
	ObjectType	string `json:"docType"` //  이 오브젝트 타입에 만든 구조체 이름을 넣으면 인덱스를 찾을 수 있음
	Key string `json:"key"` // 키
	StartPoint string `json:"startpoint"` // 출발지
	EndPoint string `json:"endpoint"` // 도착지 
	CarWeight string `json:"carweight"`  // 차 톤수
	Car string `json:"car"` // 차 종류
	Weight string `json:"weight"`   // 적재 중량
	TransPort string `json:"transport"` // 운행방법 1:편도 2:왕복
	Cost string `json:"cost"` // 금액 
	Average string `json:"average"` // 평균 금액
	Date string `json:"date"` // 완료 시간 
}


func (s *ChainCode) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}



func (s *ChainCode) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()

	// if function == "addUser" {
	// 	return s.addUser(APIstub, args)
	// } 
	if function == "addTruck" {
		return s.addTruck(APIstub, args)
	} else if function == "getTruck" {
		return s.getTruck(APIstub, args)
	} else if function == "getAllTruck" {
		return s.getAllTruck(APIstub)
	} else if function == "getHistory" {
		return s.getHistory(APIstub, args)
	} else if function == "getBatteryValue" {
		return s.getBatteryValue(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// 유저 등록
// func (s *ChainCode) addUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

// 	if len(args) != 1 {
// 		return shim.Error("fail!")
// 	}
// 	var user = User{Phone: args[0]}
// 	userAsBytes, _ := json.Marshal(user)
// 	APIstub.PutState(args[0], userAsBytes)

// 	return shim.Success(nil)
// }

// 데이터 입력
func (s *ChainCode) addTruck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var data = Truck{ObjectType: "Truck",Key:args[0],StartPoint:args[1], EndPoint:args[2], CarWeight:args[3], Car:args[4], Weight:args[5], TransPort:args[6], Cost:args[7], Average:args[7] , Date:time.Now().Format("20060102150405") }
	userAsBytes,_:=json.Marshal(data)

	// 월드스테이드 업데이트 
	APIstub.PutState(args[0], userAsBytes)

	indexName := "startpoint~id"
	startpointidIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{data.StartPoint, data.Key})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	APIstub.PutState(startpointidIndexKey, value)


	return shim.Success([]byte("rating is updated"))

}

// 키값 데이터 조회
func (s *ChainCode) getTruck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	value, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to get Battery")
	}
	if value == nil {
		return shim.Error("value not found")
	}
	return shim.Success(value)
}

// 모든 데이터 조회
func (s *ChainCode) getAllTruck(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "00000000000"
	endKey := "999999999999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	var buffer string
	buffer ="["

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer += ","
		}

		buffer += string(response.Value)

		bArrayMemberAlreadyWritten = true
	}
	buffer += "]"
	return shim.Success([]byte(buffer))
}


//밸류 조회
func (s *ChainCode) getBatteryValue(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	value := args[0]
	queriedIdByValueIterator, err := APIstub.GetStateByPartialCompositeKey("startpoint~id", []string{value})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer queriedIdByValueIterator.Close()


	var buffer string
	buffer = "["
	bArrayMemberAlreadyWritten := false
	
	var i int

	for i = 0; queriedIdByValueIterator.HasNext(); i++ {
		response, err := queriedIdByValueIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(response.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		returnedName := compositeKeyParts[0]
		returnedKey := compositeKeyParts[1]

		fmt.Printf("- found a key from index:%s name:%s key:%s\n", objectType, returnedName, returnedKey)
		if bArrayMemberAlreadyWritten == true {
			buffer += ", "
		}

		truckAsBytes, err := APIstub.GetState(returnedKey)
		if err != nil {
			return shim.Error(err.Error())
		}

		buffer += string(truckAsBytes)
		bArrayMemberAlreadyWritten = true
	}

	buffer += "]"

	return shim.Success([]byte(buffer))
}





// 키 이력 조회
func (s *ChainCode) getHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	truckName := args[0]

	fmt.Printf("- start getHistoryForBattery: %s\n", truckName)

	resultsIterator, err := APIstub.GetHistoryForKey(truckName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()


	var buffer string
	buffer ="["

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer += ","
		}

			buffer += string(response.Value)

		bArrayMemberAlreadyWritten = true
	}
	buffer += "]"

	return shim.Success([]byte(buffer))
}


func main() {
	if err := shim.Start(new(ChainCode)); err != nil {
		fmt.Printf("Error starting ChainCode chaincode: %s", err)
	}
}