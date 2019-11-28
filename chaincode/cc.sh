docker exec cli peer chaincode install -n cargo -v 1.0 -p github.com/trcc/
docker exec cli peer chaincode install -n cargo -v 1.1 -p github.com/trcc/



export CORE_PEER_LOCALMSPID=Org2MSP
export PEER0_ORG2_CA=/etc/hyperledger/crypto/peerOrganizations/org2.trucker.com/peers/peer0.org2.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.trucker.com/users/Admin@org2.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org2.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/trcc/

export CORE_PEER_LOCALMSPID=Org3MSP
export PEER0_ORG3_CA=/etc/hyperledger/crypto/peerOrganizations/org3.trucker.com/peers/peer0.org3.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.trucker.com/users/Admin@org3.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org3.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/trcc/


export CORE_PEER_LOCALMSPID=Org1MSP
export PEER0_ORG1_CA=/etc/hyperledger/crypto/peerOrganizations/org1.trucker.com/peers/peer0.org1.trucker.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.trucker.com/users/Admin@org1.trucker.com/msp
export CORE_PEER_ADDRESS=peer0.org1.trucker.com:7051
docker exec cli peer chaincode install -n truck -v 1.0 -p github.com/trcc/


docker exec cli peer chaincode instantiate -v 1.0 -C trucker -n cargo -c '{"Args":["Init"]}' -P 'OR ("Org1MSP.member", "Org2MSP.member")'
# docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addUser","00000000000"]}'
docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addTruck","00000000000","서울특별시 강남구","서울","1T","트럭","300kg","편도","70000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getTruck","00000000000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getAllTruck"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getHistory","00000000000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getBatteryValue","인천"]}'

docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addGeo","서울특별시 강남"]}'
docker exec cli peer chaincode invoke -o orderer.trucker.com:7050 -C trucker -n cargo -c '{"Args":["addAverage","서울특별시","50000"]}'
docker exec cli peer chaincode query -C trucker -n cargo -c '{"Args":["getAverage","서울특별시"]}'

docker exec cli peer chaincode upgrade -v 1.1 -C trucker -n cargo -c '{"Args":["Init"]}' -P 'OR ("Org1MSP.member", "Org2MSP.member")'

echo '-------------------------------------END-------------------------------------'

