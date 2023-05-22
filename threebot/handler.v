module main

import freeflowuniverse.crystallib.jsonrpc

fn handler(data string) !string {
	method := jsonrpc.jsonrpcrequest_decode_method(data)!
	match method {
		'test' {
			request := jsonrpc.jsonrpcrequest_decode[string](data)!
			result := test(request.params)
			response := jsonrpc.new_jsonrpcresponse[int](request.id, result)
			return response.to_json()
		}
		else {} 
	}
	return 'request'
}