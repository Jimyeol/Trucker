package main

import (
	"encoding/json"
	"fmt"
	// "bytes"
	"time"
	"strconv"


	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)


type ChainCode struct {
}



// 트럭 구조체
type Truck struct {
	ObjectType	string `json:"docType"` //  이 오브젝트 타입에 만든 구조체 이름을 넣으면 인덱스를 찾을 수 있음
	Key string `json:"key"` // 키
	StartPoint string `json:"startpoint"` // 출발지
	EndPoint string `json:"endpoint"` // 도착지 
	CarWeight string `json:"carweight"`  // 차 톤수
	Car string `json:"car"` // 차 종류
	Weight string `json:"weight"`   // 적재 중량
	TransPort string `json:"transport"` // 운행방법 1:편도 2:왕복
	Money string `json:"money"` // 금액 
	Average string `json:"average"` // 평균 금액
	Date string `json:"date"` // 완료 시간 
}

// 지역 평균금액 

type GeoAvg struct {
	Geo string `json:"geo"`
	Average int64 `json:"average"`	
	Cost []Cost `json:"cost"`
}

type Cost struct {
	Money int64 `json:"money"`
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
	} else if function == "getTruckByValue" {
		return s.getTruckByValue(APIstub, args)
	} else if function == "addAverage" {
		return s.addAverage(APIstub, args)
	} else if function == "getAverage" {
		return s.getAverage(APIstub,args)
	} else if function =="addGeo" {
		return s.addGeo(APIstub,args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}



// 데이터 입력
func (s *ChainCode) addTruck(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	var data = Truck{ObjectType: "Truck",Key:args[0],StartPoint:args[1], EndPoint:args[2], CarWeight:args[3], Car:args[4], Weight:args[5], TransPort:args[6], Money:args[7], Average:args[7] , Date:time.Now().Format("20060102150405") }
	userAsBytes,_:=json.Marshal(data)

	// 월드스테이드 업데이트 
	APIstub.PutState(args[0], userAsBytes)



	// StartPoint 인덱스 조회를 위한 CompositeKey 생성 
	indexName := "startpoint~id"
	startpointidIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{data.StartPoint, data.Key})
	if err != nil {
		return shim.Error(err.Error())
	}
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


//StartPoint 조회
func (s *ChainCode) getTruckByValue(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
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

func (s *ChainCode) addGeo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("fail!")
	}
	var geo = GeoAvg{Geo: args[0], Average: 0}
	avgAsBytes, _ := json.Marshal(geo)
	APIstub.PutState(args[0], avgAsBytes)

	return shim.Success(nil)
}


//평균금액 추가 

func (s *ChainCode) addAverage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// 지역 확인
	avgAsBytes, err := APIstub.GetState(args[0])
	if err != nil{
		jsonResp := "\"Error\":\"Failed to get state for "+ args[0]+"\"}"
		return shim.Error(jsonResp)
	} else if avgAsBytes == nil{ // no State! error
		jsonResp := "\"Error\":\"User does not exist: "+ args[0]+"\"}"
		return shim.Error(jsonResp)
	}

	// 평균금액 구조체 선언


	Avg := GeoAvg{}
	err = json.Unmarshal(avgAsBytes, &Avg)
	
	Avg.Geo=args[0]
	newCost,_ := strconv.ParseInt(args[1],0,64)
	costCount := int64(len(Avg.Cost))
	var cost = Cost{Money:newCost}
	Avg.Cost=append(Avg.Cost,cost)
	Avg.Average = (costCount*Avg.Average+newCost)/(costCount+1)

	// update to User World state
	avgAsBytes, err = json.Marshal(Avg);

	APIstub.PutState(args[0], avgAsBytes)

	return shim.Success([]byte("avg is updated"))

}


//평균금액 조회

func (s *ChainCode) getAverage(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AvgAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(AvgAsBytes)
}



func main() {
	if err := shim.Start(new(ChainCode)); err != nil {
		fmt.Printf("Error starting ChainCode chaincode: %s", err)
	}
}