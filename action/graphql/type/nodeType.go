package graphqltype

import "github.com/graphql-go/graphql"

// NodeType : Graphql object type of Node
var NodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Node",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"server_uuid": &graphql.Field{
				Type: graphql.String,
			},
			"bmc_mac_addr": &graphql.Field{
				Type: graphql.String,
			},
			"bmc_ip": &graphql.Field{
				Type: graphql.String,
			},
			"pxe_mac_addr": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"cpu_cores": &graphql.Field{
				Type: graphql.Int,
			},
			"memory": &graphql.Field{
				Type: graphql.Int,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"active": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// NodeListType : Graphql object type of NodeList
var NodeListType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "NodeList",
		Fields: graphql.Fields{
			"node_list": &graphql.Field{
				Type: graphql.NewList(NodeType),
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)

// PowerStateNodeType : Graphql object type of PowerStateNode
var PowerStateNodeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PowerStateNode",
		Fields: graphql.Fields{
			"result": &graphql.Field{
				Type: graphql.String,
			},
			"errors": &graphql.Field{
				Type: graphql.NewList(Errors),
			},
		},
	},
)