package okclient

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryABCIInfo(t *testing.T){
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryABCIInfo()
	assertNotEqual(t,err,nil)
	jsonBytes,err:=json.Marshal(resp)
	assertNotEqual(t,err,nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryConsenusState(t *testing.T){
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryConsenusState()
	assertNotEqual(t,err,nil)
	fmt.Println(string(resp))
}