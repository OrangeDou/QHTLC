package utils

import (
	"fmt"

	"doctorProject/model/blockchain"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func StartLedger() error {
	contractChaincode := blockchain.ContractChaincode{}
	err := shim.Start(&contractChaincode)
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
		return err
	}
	return nil
}
