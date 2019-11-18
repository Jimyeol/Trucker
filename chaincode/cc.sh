docker exec cli peer chaincode install -n cargo -v 1.0 -p github.com/bacc/


export CORE_PEER_LOCALMSPID=Org2MSP
export PEER0_ORG2_CA=/etc/hyperledger/crypto/peerOrganizations/org2.trucker.com/peers/peer0.org2.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.trucker.com/users/Admin@org2.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org2.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/bacc/

export CORE_PEER_LOCALMSPID=Org3MSP
export PEER0_ORG3_CA=/etc/hyperledger/crypto/peerOrganizations/org3.trucker.com/peers/peer0.org3.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.trucker.com/users/Admin@org3.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org3.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/bacc/


export CORE_PEER_LOCALMSPID=Org1MSP
export PEER0_ORG1_CA=/etc/hyperledger/crypto/peerOrganizations/org1.trucker.com/peers/peer0.org1.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.trucker.com/users/Admin@org1.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org1.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/bacc/

#export ORDERER_CA=/etc/hyperledger/crypto/ordererOrganizations/trucker.com/orderers/orderer.trucker.com/msp/tlscacerts/tlsca.trucker.com-cert.pem
#docker exec cli peer chaincode instantiate -o orderer.trucker.com:7050 -C trucker -n truck -v 1.0 -c '{"Args":["init"]}' -P "OR('Org1MSP.member','Org2MSP.member')" --collections-config /opt/gopath/src/github.com/trucker/collections_config.json



docker exec cli peer chaincode instantiate -v 1.0 -C trucker -n cargo -c '{"Args":["Init"]}' -P 'OR ("Org1MSP.member", "Org2MSP.member")'
docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addUser","00000000000"]}'
docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addTruck","00000000000","인천","서울","1T","트럭","300kg","편도","70000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getTruck","00000000000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getAllTruck"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getHistory","00000000000"]}'



echo '-------------------------------------END-------------------------------------'

