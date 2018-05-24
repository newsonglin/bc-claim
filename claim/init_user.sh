#Invoke users web service and then extract the token exactaly,I'm proud of this method
echo -e "\n********************create new user*********************"
MY_TOKEN=`curl -s -X POST http://localhost:4000/services/users -H "content-type: application/x-www-form-urlencoded" -d 'username=Songlin&orgName=Org1'|grep -oP '(?<="token":").*?(?=")'`

echo "this is  $MY_TOKEN "

echo -e  "\n********************searchMD*********************"

#Invoke searchMD
curl -s -X GET \
  "http://localhost:4000/services/channels/mychannel/chaincodes/mycc?peer=peer0.org1.exigen.com&fcn=searchAll&args=%5B%22MR0%22%5D" \
  -H "authorization: Bearer $MY_TOKEN" \
  -H "content-type: application/json"



      