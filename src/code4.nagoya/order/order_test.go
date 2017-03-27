package order

import (
	"fmt"
	"testing"
)

func TestOrderChainCodeValidate(t *testing.T) {

	fmt.Printf("start  order chaincode validate\n")

	orderChainCode := new(OrderChainCode)
	orderChainCode.validate()
	
	fmt.Println("end order chaincode validate")

}
