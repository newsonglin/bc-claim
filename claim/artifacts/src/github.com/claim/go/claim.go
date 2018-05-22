/*
 * The  smart contract for claim:
 * Note, we create a medical record struct.
 * That will go to the block chain when patient go to see a doctor
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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the medicalRecord structure, with 4 properties.  Structure tags are used by encoding/json library
type MedicalRecord struct {
	MedicalID   string `json:"medicalID"`
	Hospital  string `json:"hospital"`
	Medicines string `json:"medicines"`
	Owner  string `json:"owner"`
}

/*
 * The Init method is called when the Smart Contract "claimMedicalRecord" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initMD()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "claimMedicalRecord"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "searchMD" {
		return s.searchMD(APIstub, args)
	} else if function == "initMD" {
		return s.initMD(APIstub)
	} else if function == "createMD" {
		return s.createMD(APIstub, args)
	} else if function == "searchAll" {
		return s.searchAll(APIstub)
	} else if function == "updateOwner" {
		return s.updateOwner(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) searchMD(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	medicalRecordAsBytes, _ := APIstub.GetState(args[0])

	return shim.Success(medicalRecordAsBytes)
}

func (s *SmartContract) initMD(APIstub shim.ChaincodeStubInterface) sc.Response {
	medicalRecords := []MedicalRecord{
		MedicalRecord{MedicalID: "MDCN00000001", Hospital: "Ruijin Hospital", Medicines: "Aspirin,", Owner: "Zhao Yi"},
		MedicalRecord{MedicalID: "MDCN00000002", Hospital: "Ruijin Hospital", Medicines: "Penicillin", Owner: "Qian Er"},
		MedicalRecord{MedicalID: "MDCN00000003", Hospital: "Heping Hospital", Medicines: "Penicillin", Owner: "Sun San"},
		MedicalRecord{MedicalID: "MDCN00000004", Hospital: "Heping Hospital", Medicines: "Erythromycin", Owner: "Li Si"},
		MedicalRecord{MedicalID: "MDCN00000005", Hospital: "336 Hospital", Medicines: "Erythromycin", Owner: "Zhou Wu"},
		MedicalRecord{MedicalID: "MDCN00000006", Hospital: "Redcross Hospital", Medicines: "Penicillin", Owner: "Zheng Liu"},
		MedicalRecord{MedicalID: "MDCN00000007", Hospital: "Redcross Hospital", Medicines: "Aspirin", Owner: "Wang Qi"},
		MedicalRecord{MedicalID: "MDCN00000008", Hospital: "Suzhou Children Hospital", Medicines: "ChuanBeiPiPa", Owner: "Feng Ba"},
		MedicalRecord{MedicalID: "MDCN00000009", Hospital: "Suzhou Children Hospital", Medicines: "Erythromycin", Owner: "Chen Jiu"},
		MedicalRecord{MedicalID: "MDCN00000010", Hospital: "336 Hospital", Medicines: "Erythromycin", Owner: "Shotaro"},
	}

	i := 0
	for i < len(medicalRecords) {
		fmt.Println("i is ", i)
		medicalRecordAsBytes, _ := json.Marshal(medicalRecords[i])
		APIstub.PutState("MR"+strconv.Itoa(i), medicalRecordAsBytes)
		fmt.Println("Added", medicalRecords[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createMD(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var medicalRecord = MedicalRecord{MedicalID: args[1], Hospital: args[2], Medicines: args[3], Owner: args[4]}

	medicalRecordAsBytes, _ := json.Marshal(medicalRecord)
	APIstub.PutState(args[0], medicalRecordAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) searchAll(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "MR0"
	endKey := "MR999"

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

func (s *SmartContract) updateOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	medicalRecordAsBytes, _ := APIstub.GetState(args[0])
	medicalRecord := MedicalRecord{}

	json.Unmarshal(medicalRecordAsBytes, &medicalRecord)
	medicalRecord.Owner = args[1]

	medicalRecordAsBytes, _ = json.Marshal(medicalRecord)
	APIstub.PutState(args[0], medicalRecordAsBytes)

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
