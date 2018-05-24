/**
 *This is a test file, don't worry about it. It has no functional features, but it do useful
 */
package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkUserInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("userTest", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkUserInvoke(t *testing.T, stub *shim.MockStub, args [][]byte, retval []byte) {
	res := stub.MockInvoke("userTest", args)
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

func TestUser_Init(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("userTest", scc)
	checkUserInit(t, stub, [][]byte{[]byte("init"), nil})
}

func TestUserAll(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("userTest", scc)
	
	// Invoke initUser
	fmt.Println("Test user init ==================")
	checkUserInvoke(t, stub, [][]byte{[]byte("initUser")},nil)
	
	//Invoke login
	fmt.Println("Test user login ==================")
	checkUserInvoke(t, stub, [][]byte{[]byte("login"), []byte("test02"), []byte("pass02")}, nil)
	
    //Invoke login
	fmt.Println("Test user register ==================")
	checkUserInvoke(t, stub, [][]byte{[]byte("register"), []byte("test03"), []byte("pass03"),[]byte("First03"),[]byte("Last03")}, nil)
	
   }