# Instantiate chaincode
echo -e  "\n********************Instantiate chaincode*********************"

curl -s -X POST \
  http://localhost:4000/services/channels/mychannel/chaincodes \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjcxMDA2NTQsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1MjcwNjQ2NTR9.zKlEta2YQ-KYJtlNjrPCmFBSju2NI_41SFrO8riKhO4" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.exigen.com","peer1.org1.exigen.com"],
	"chaincodeName":"mycc",
	"chaincodeVersion":"v0",
	"chaincodeType": "golang",
	"args":["a","100","b","200"]
}'
