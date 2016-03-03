package main

import (
	"fmt"
	"jsonrpc"
)

func main() {
	params := map[string]interface{}{"service": "object", "method": "execute", "args": []interface{}{"oev8", 1, "admin", "res.partner", "search", []interface{}{}}}
	conn := jsonrpc.NewClient("http://odoo-server:8069/jsonrpc")
	var response *jsonrpc.Response
	var erreur error
	response, erreur = conn.Call("call", params)
	if erreur == nil {
		fmt.Printf("Response %s", response.Result)
	} else {
		fmt.Printf("erreur %s", erreur)
	}
	params = map[string]interface{}{"service": "object", "method": "execute", "args": []interface{}{"oev8", 1, "admin", "res.partner", "read", response.Result}}
	response, erreur = conn.Call("call", params)
	if erreur == nil {
		fmt.Printf("Response %s", response.Result)
	} else {
		fmt.Printf("erreur %s", erreur)
	}

}
