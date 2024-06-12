package blockchain

import (
	model "doctorProject/model/user"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/hyperledger/fabric-contract-api-go/serializer"
	"github.com/hyperledger/fabric-protos-go/peer"
)

// TODO 增加日志记录
type ContractChaincode struct {
	DefaultContract string

	Info                  metadata.InfoMetadata
	TransactionSerializer serializer.TransactionSerializer
	// contains filtered or unexported fields
}

// 初始化数据状态，实例化/升级链码时被自动调用
func (cc *ContractChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()

	var (
		alice, bob  model.User //	定义两个用户
		aliceAmount = alice.GetWalletAmount()
		bobAmount   = bob.GetWalletAmount()
		err         error
	)
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	alice.Username = args[0]
	aliceAmount, err = strconv.Atoi(args[1]) //alice钱包余额
	if err != nil {
		return shim.Error("Expecting integer value for wallet amount")
	}

	bob.Username = args[2]
	bobAmount, err = strconv.Atoi(args[3]) //bob钱包余额
	if err != nil {
		return shim.Error("Expecting integer value for wallet amount")
	}

	fmt.Printf("Alice's wallet amount: %d\n", aliceAmount)
	fmt.Printf("Bob's wallet amount: %d\n", bobAmount)

	err = stub.PutState(alice.Username, []byte(strconv.Itoa(aliceAmount)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(bob.Username, []byte(strconv.Itoa(bobAmount)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (cc *ContractChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()

	// 调用 invoke 函数实现转账操作
	if function == "Transfer" {
		return cc.Transfer(stub, args)
	}
	// 调用 query 函数实现账户查询操作
	//TODO 补充查询，删除逻辑
	if function == "Query" {
		return cc.Query(stub, args)
	}
	// 调用 delete 函数实现账户注销操作
	if function == "Delete" {
		return cc.Delete(stub, args)
	}

	// 传递的函数名出错，返回 shim.Error()
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

func (cc *ContractChaincode) Transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var (
		alice, bob model.User
		err        error
		aval       = alice.Wallet.Amount
		bval       = bob.Wallet.Amount
	)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	alice.Username = args[0]
	bob.Username = args[1]

	// 获取账本中alice账户余额
	aliceAmountByte, err := stub.GetState(alice.Username)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if aliceAmountByte == nil {
		return shim.Error("Entity not found")
	}
	aval, _ = strconv.Atoi(string(aliceAmountByte))

	//	获取账本中bob账户余额
	bobAmountByte, err := stub.GetState(bob.Username)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if bobAmountByte == nil {
		return shim.Error("Entity not found")
	}
	bval, _ = strconv.Atoi(string(bobAmountByte))

	// 转账
	transferValue, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	aval = aval - transferValue
	bval = bval + transferValue
	fmt.Printf("Aval = %d, Bval = %d\n", aval, bval)

	// 更新转账后账本中 A 余额
	err = stub.PutState(alice.Username, []byte(strconv.Itoa(aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	// 更新转账后账本中 B 余额
	err = stub.PutState(bob.Username, []byte(strconv.Itoa(bval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// 账户余额查询
func (cc *ContractChaincode) Query(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var userName string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	userName = args[0] // 账户用户名

	// 从账本中获取该账户余额
	avalbytes, err := stub.GetState(userName)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + userName + "\"}"
		return shim.Error(jsonResp)
	}

	if avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + userName + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + userName + "\",\"Amount\":\"" + string(avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	// 返回转账金额
	return shim.Success(avalbytes)
}

// TODO账户查询
func (cc *ContractChaincode) WalletExists(stub shim.ChaincodeStubInterface, args []string) (bool, error) {
	result, err := stub.GetState(args[0])
	if err != nil {
		_ = shim.Error("Failed to get the Wallet," + err.Error())
		return false, err
	}
	if result == nil {
		return false, nil
	}
	return true, nil
}

// 账户注销
func (cc *ContractChaincode) Delete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userName := args[0] // 账户用户名

	// 从账本中删除该账户状态
	err := stub.DelState(userName)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// // 更新账户数据
//
//	func (cc *ContractChaincode) UpdateAsset(stub shim.ChaincodeStubInterface, id string) error {
//		exists, err := cc.WalletExists(stub, id)
//		if err != nil {
//			return err
//		}
//		if !exists {
//			return fmt.Errorf("the asset %s does not exist", id)
//		}
//		asset, err := cc.ReadAsset(stub, id)
//		if err != nil {
//			return fmt.Errorf("Failed to read the asset :%s ", id)
//		}
//	}
//
// TODO账户创建 修改逻辑
func (cc *ContractChaincode) CreateAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var (
		alice model.User
		err   error
	)

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	alice.Username = args[0]
	amount, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset amount")
	}

	err = stub.PutState(alice.Username, []byte(strconv.Itoa(amount)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
