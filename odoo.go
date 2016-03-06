package main

import (
	"fmt"
	"github.com/puper/go-jsonrpc/jsonrpc"
)
func main() {
	params := map[string]interface{}{"service":"object","method":"execute","args":[]interface{}{"oev8",1,"admin","res.partner","search",[]interface{}{}}}
	conn  := jsonrpc.NewClient("http://127.0.0.1:8069/jsonrpc")
	var response   *jsonrpc.Response
	var erreur error
	response, erreur = conn.Call("call",params)
	if erreur == nil {
		fmt.Printf("Response %s",response.Result)
	}
}

