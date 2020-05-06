// Code generated by protoc-gen-hb-grpc-gateway. DO NOT EDIT.
// source: proto/service.proto

/*
Package proto is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package proto

import (
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule"
	"github.com/hb-go/grpc-contrib/registry"
)

var (
	route_ExampleService_ApiCall_0 = registry.Binding{
		Method: "GET",
		PathTmpl: &httprule.Template{
			Version: 1,
			OpCodes: []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3},
			Pool:    []string{"v1", "example", "call", "name"},
			Verb:    "",
		},
		AssumeColonVerb: true,
	}

	route_ExampleService_ApiCall_1 = registry.Binding{
		Method: "GET",
		PathTmpl: &httprule.Template{
			Version: 1,
			OpCodes: []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4},
			Pool:    []string{"v1", "example", "call", "responsebody", "name"},
			Verb:    "",
		},
		AssumeColonVerb: true,
	}

	route_ExampleService_Call_0 = registry.Binding{
		Method: "POST",
		PathTmpl: &httprule.Template{
			Version: 1,
			OpCodes: []int{2, 0, 2, 1, 2, 2},
			Pool:    []string{"v1", "example", "call"},
			Verb:    "",
		},
		AssumeColonVerb: true,
	}

	route_ExampleService_Call_1 = registry.Binding{
		Method: "POST",
		PathTmpl: &httprule.Template{
			Version: 1,
			OpCodes: []int{2, 0, 2, 1, 2, 2, 2, 3},
			Pool:    []string{"v1", "example", "call", "responsebody"},
			Verb:    "",
		},
		AssumeColonVerb: true,
	}
)

var GatewayServiceExampleService = registry.Service{
	Name: _ExampleService_serviceDesc.ServiceName,
	Methods: []*registry.Method{
		&registry.Method{
			Name: "ApiCall",
			Bindings: []*registry.Binding{
				&route_ExampleService_ApiCall_0,
				&route_ExampleService_ApiCall_1,
			},
		},

		&registry.Method{
			Name: "Call",
			Bindings: []*registry.Binding{
				&route_ExampleService_Call_0,
				&route_ExampleService_Call_1,
			},
		},
	},
}
