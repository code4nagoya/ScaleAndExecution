package order

import (
	"errors"
	"fmt"
	//	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type OrderChainCode struct {
}

func (c *OrderChainCode) validate() {

	fmt.Printf("validate\n")

}

func (c *OrderChainCode) init() {

	fmt.Printf("init\n")
}

func (c *OrderChainCode) invoke() ([]byte, error) {

	fmt.Printf("invoke\n")
	return nil, nil
}

func (c *OrderChainCode) delete() ([]byte, error) {

	fmt.Printf("delete\n")
	return nil, nil
}

func (c *OrderChainCode) query() ([]byte, error) {

	fmt.Printf("query\n")

	return nil, nil

}

func (c *OrderChainCode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {

	_, args := stub.GetFunctionAndParameters()

	fmt.Println(args)

	return nil, nil

}

//
func (c *OrderChainCode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	function, _ := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return c.invoke()
	} else if function == "delete" {
		// Deletes an entity from its state
		return c.delete()
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return c.query()
	}

	return nil, errors.New("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}
