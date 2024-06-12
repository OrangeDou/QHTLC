#！/bin/bash
cd /root/go-workspace/src/github.com/hyperledger/fabric/scripts/fabric-samples/test-network
./network.sh down

./network.sh up

./network.sh createChannel -c mychannel

#部署链码
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go

#调用invoke函数
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"InitLedger","Args":[]}'


