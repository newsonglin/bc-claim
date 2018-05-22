/**
 *This is a test file, don't worry about it. It has no functional features, but it do useful
 */
package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkClaimInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("claimTest", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkClaimInvoke(t *testing.T, stub *shim.MockStub, args [][]byte, retval []byte) {
	res := stub.MockInvoke("claimTest", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
	
    if retval != nil {
		if res.Payload == nil {
			fmt.Printf("Invoke returned nil, expected %s", string(retval))
			t.FailNow()
		}
		if string(res.Payload) != string(retval) {
			fmt.Printf("Invoke returned %s, expected %s", string(res.Payload), string(retval))
			t.FailNow()
		}
	}	
	fmt.Println(string(res.Payload))
	
}

func TestClaim_Init(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("claimTest", scc)
	checkClaimInit(t, stub, [][]byte{[]byte("init"), nil})
}

func TestClaimAll(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("claimTest", scc)
	//checkClaimInit(t, stub, [][]byte{[]byte("init"), nil})
	
	// Invoke initMD
	fmt.Println("Test medical record init ==================")
	checkClaimInvoke(t, stub, [][]byte{[]byte("initMD")},nil)
	
	//Invoke searchMD
	fmt.Println("Test search medical record ==================")
	expectedVal:=[] byte("{\"medicalID\":\"MDCN00000001\",\"hospital\":\"Ruijin Hospital\",\"medicines\":\"Aspirin,\",\"owner\":\"Zhao Yi\"}")
	checkClaimInvoke(t, stub, [][]byte{[]byte("searchMD"), []byte("MR0")}, expectedVal)
	
	//Invoke createMD
	fmt.Println("Test create medical record ==================")
	newMD:=[] byte("{\"medicalID\":\"TEST\",\"hospital\":\"TEST Hospital\",\"medicines\":\"TEST Aspirin\",\"owner\":\"TEST Zhao Yi\"}")
	checkClaimInvoke(t, stub, [][]byte{[]byte("createMD"),[]byte("MR9"), []byte("TEST"),[]byte("TEST Hospital"),[]byte("TEST Aspirin"),[]byte("TEST Zhao Yi"),}, nil)
	checkClaimInvoke(t, stub, [][]byte{[]byte("searchMD"), []byte("MR9")}, newMD)
	
	//Invoke searchAll
	fmt.Println("Test search all medical records ==================")
	checkClaimInvoke(t, stub, [][]byte{[]byte("searchAll")},nil)
}