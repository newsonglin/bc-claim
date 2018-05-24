# Invoke initMD

echo -e  "\n********************initMD*********************"

curl -s -X POST \
  http://localhost:4000/services/channels/mychannel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjcxNjY1MjAsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1MjcxMzA1MjB9.11pZDEA_0tHHC2-Rore-s6N2-Hi4KqVd2jLG60BncZs" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.exigen.com","peer1.org1.exigen.com"],
	"fcn":"initMD",
	"args":["ANYTHING"]
}'


echo -e  "\n********************searchMD*********************"

#Invoke searchMD
curl -s -X GET \
  "http://localhost:4000/services/channels/mychannel/chaincodes/mycc?peer=peer0.org1.exigen.com&fcn=searchAll&args=%5B%22MR0%22%5D" \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjcxNjY1MjAsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1MjcxMzA1MjB9.11pZDEA_0tHHC2-Rore-s6N2-Hi4KqVd2jLG60BncZs" \
  -H "content-type: application/json"