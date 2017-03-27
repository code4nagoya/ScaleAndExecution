package order

import (
	"fmt"
	//	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestOrderChainCodeValidate(t *testing.T) {

	fmt.Printf("start  order chaincode validate\n")

	orderChainCode := new(OrderChainCode)
	orderChainCode.validate()

	fmt.Println("end order chaincode validate")

}

func TestOrderChainCodeInit(t *testing.T) {

	fmt.Printf("start  order chaincode init\n")

	orderChainCode := new(OrderChainCode)
	//	orderChainCode.Init(new(shim.ChaincodeStubInterface))
	orderChainCode.init()
	fmt.Println("end order chaincode init")

}

func TestOrderChainCodeInvoke(t * testing.T){
	fmt.Println("start order chaincode invoke" )
	
	orderChainCode := new(OrderChainCode)
	orderChainCode.invoke()
	fmt.Println("end order chaincode invoke " )
}


func TestOrderChainCodeDelete(t *testing.T){
	
	fmt.Println("start ChainCode Delete" )
	

	orderChainCode := new(OrderChainCode)
	//	orderChainCode.Init(new(shim.ChaincodeStubInterface))
	orderChainCode.delete()
	
	
	
	fmt.Println("end ChainCode Delete")
	
}

//
//
// test query
func TestOrderChainCodeQuery(t *testing.T) {

	fmt.Printf("start  order chaincode init\n")

	orderChainCode := new(OrderChainCode)
	//	orderChainCode.Init(new(shim.ChaincodeStubInterface))
	orderChainCode.query()
	fmt.Println("end order chaincode init")

}

