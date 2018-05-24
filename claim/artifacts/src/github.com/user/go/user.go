/*
 * The  smart contract for user:
 * This will store user creditial temportly, user authenication should moved to CA later
 * 
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the user structure, with 4 properties.  Structure tags are used by encoding/json library
type User struct {
	LoginID   string `json:"loginID"`
	PWD  string `json:"pwd"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

/*
 * The Init method is called when the Smart Contract "claimUser" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initUser()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "claimUser"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "login" {
		return s.login(APIstub, args)
	}  else if function == "initUser" {
		return s.initUser(APIstub)
	} else if function == "register" {
		return s.register(APIstub, args)
	} else if function == "update" {
		return s.update(APIstub, args)
	}

	return shim.Error("Invalid user function name.")
}

/**
  *
  * login function to check user name and password is correct
  *
  */

func (s *SmartContract) login(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2") //Expected loginid and password
	}

	startKey := "USER0"
	endKey := "USER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()


	var user User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf(string(queryResponse.Value))
		

        err = json.Unmarshal(queryResponse.Value, &user)
        if err != nil {
            return shim.Error(err.Error())
        }
    

	    if(user.LoginID==args[0] && user.PWD==args[1]){
    	    return shim.Success([]byte(queryResponse.Key))
        }
		
	}

	return shim.Success([]byte("fail"))
	
}


/**
 * create a new user
 */
func (s *SmartContract) register(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	
	
	startKey := "USER0"
	endKey := "USER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()


	var user User
	lastKey:="USER0"
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Printf(string(queryResponse.Value))
		

        err = json.Unmarshal(queryResponse.Value, &user)
        if err != nil {
            return shim.Error(err.Error())
        }
    

	    if(user.LoginID==args[0]){
    	    return shim.Error("User is already exist:"+args[0])
        }
		
        lastKey=queryResponse.Key
	}
	fmt.Println("- register last key:\n"+lastKey+"\n")
    
	lastKey=strings.Replace(lastKey, "USER", "", -1)
	
	index, err := strconv.Atoi(lastKey)
    if err != nil {
        return shim.Error(err.Error())
    }
	fmt.Printf("- register last key index:\n %d \n", index)
	index=index+1;
    lastKey="USER"+strconv.Itoa(index)
	fmt.Println("- register the new last key:\n"+lastKey+"\n")

    
	user = User{LoginID: args[0], PWD: args[1], FirstName: args[2], LastName: args[3]}

	userAsBytes, _ := json.Marshal(user)
	APIstub.PutState(lastKey, userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) searchAll(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "USER0"
	endKey := "USER999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- searchAll:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) update(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	userAsBytes, _ := APIstub.GetState(args[0])
	user := User{}

	json.Unmarshal(userAsBytes, &user)
	user.LastName = args[1]

	userAsBytes, _ = json.Marshal(user)
	APIstub.PutState(args[0], userAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) initUser(APIstub shim.ChaincodeStubInterface) sc.Response {
	users := []User{

		User{LoginID: "test01", PWD: "pass01", FirstName: "First01,", LastName: "Last01"},
		User{LoginID: "test02", PWD: "pass02", FirstName: "First02", LastName: "Last02"},
	}

	i := 0
	for i < len(users) {
		fmt.Println("i is ", i)
		userAsBytes, _ := json.Marshal(users[i])
		APIstub.PutState("USER"+strconv.Itoa(i), userAsBytes)
		fmt.Println("Added", users[i])
		i = i + 1
	}

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
