sleep 5s
#Invoke users web service and then extract the token exactaly,I'm proud of this method
MY_TOKEN=`curl -s -X POST http://localhost:4000/services/users -H "content-type: application/x-www-form-urlencoded" -d 'username=Jim&orgName=Org1'|grep -oP '(?<="token":").*?(?=")'`

echo "there is  $MY_TOKEN "

#Create Channel request
echo -e  "\n********************Create Channel request*********************"
curl -s -X POST \
  http://localhost:4000/services/channels \
  -H "authorization: Bearer $MY_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	    "channelName":"mychannel",
	    "channelConfigPath":"../../artifacts/channel/mychannel.tx"
      }'

#Join Channel request

echo -e "\n********************Join Channel request*********************"

curl -s -X POST \
  http://localhost:4000/services/channels/mychannel/peers \
  -H "authorization: Bearer $MY_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.exigen.com","peer1.org1.exigen.com"]
}'

# Install chaincode
echo -e  "\n********************Install chaincode*********************"


curl -s -X POST \
  http://localhost:4000/services/chaincodes \
  -H "authorization: Bearer $MY_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.exigen.com","peer1.org1.exigen.com"],
	"chaincodeName":"mycc",
	"chaincodePath":"github.com/claim/go",	
	"chaincodeType": "golang",
	"chaincodeVersion":"v0"
}'




# Instantiate chaincode
echo -e  "\n********************Wait for 15s*********************"
sleep 15s

echo -e  "\n********************Instantiate chaincode*********************"

curl -s -X POST \
  http://localhost:4000/services/channels/mychannel/chaincodes \
  -H "authorization: Bearer $MY_TOKEN" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.exigen.com","peer1.org1.exigen.com"],
	"chaincodeName":"mycc",
	"chaincodeVersion":"v0",
	"chaincodeType": "golang",
	"args":["a","100","b","200"]
}'
      