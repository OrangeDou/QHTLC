package blockchain

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shimtest"
)

// 测试init方法
func TestContractChaincode_Init(t *testing.T) {
	//调用NewMockStub方法返回一个MockStub类，来模拟背书节点对链码的调用
	chaincode := new(ContractChaincode)
	stub := shimtest.NewMockStub("test", chaincode)

	res := stub.MockInit("1", [][]byte{[]byte(""), []byte("alice"), []byte("1000"), []byte("bob"), []byte("2000")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}
	if len(res.Message) != 0 {
		fmt.Print(res.Message)
	}
}

// 测试invoke方法
func TestContractChaincode_Invoke(t *testing.T) {
	//调用NewMockStub方法返回一个MockStub类，来模拟背书节点对链码的调用
	chaincode := new(ContractChaincode)
	stub := shimtest.NewMockStub("test", chaincode)

	//调用init方法初始化链码
	res := stub.MockInit("1", [][]byte{[]byte(""), []byte("alice"), []byte("1000"), []byte("bob"), []byte("2000")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}

	//调用invoke方法实现转账操作
	res = stub.MockInvoke("1", [][]byte{[]byte("transfer"), []byte("alice"), []byte("bob"), []byte("100")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}
	if len(res.Message) != 0 {
		fmt.Print(res.Message)
	}
}

// 测试query方法
func TestContractChaincode_Query(t *testing.T) {
	//调用NewMockStub方法返回一个MockStub类，来模拟背书节点对链码的调用
	chaincode := new(ContractChaincode)
	stub := shimtest.NewMockStub("test", chaincode)

	//调用init方法初始化链码
	res := stub.MockInit("1", [][]byte{[]byte(""), []byte("alice"), []byte("1000"), []byte("bob"), []byte("2000")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}

	//调用invoke方法实现转账操作
	res = stub.MockInvoke("1", [][]byte{[]byte("transfer"), []byte("alice"), []byte("bob"), []byte("100")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}

	//调用query方法查询余额
	res = stub.MockInvoke("1", [][]byte{[]byte("query"), []byte("alice")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}
	if len(res.Message) != 0 {
		fmt.Print(res.Message)
	}
}

// TODO测试delete方法
func TestContractChaincode_Delete(t *testing.T) {
	//调用NewMockStub方法返回一个MockStub类，来模拟背书节点对链码的调用
	chaincode := new(ContractChaincode)
	stub := shimtest.NewMockStub("test", chaincode)

	//调用init方法初始化链码
	res := stub.MockInit("1", [][]byte{[]byte(""), []byte("alice"), []byte("1000"), []byte("bob"), []byte("2000")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}
}

func TestContractChainCode_CreateAsset(t *testing.T) {
	chaincode := new(ContractChaincode)
	stub := shimtest.NewMockStub("test", chaincode)
	//调用init方法初始化链码
	res := stub.MockInit("1", [][]byte{[]byte(""), []byte("alice"), []byte("1000"), []byte("bob"), []byte("2000")})
	if res.Status != 200 {
		t.Fatalf("Expected CreateAsset to succeed, got %s", res.Message)
	}

}
