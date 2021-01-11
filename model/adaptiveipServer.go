package model

import (
	"hcc/piccolo/action/grpc/errconv"
	"time"
)

// AdaptiveIPServer - ish
type AdaptiveIPServer struct {
	ServerUUID     string                `json:"server_uuid"`
	PublicIP       string                `json:"public_ip"`
	PrivateIP      string                `json:"private_ip"`
	PrivateGateway string                `json:"private_gateway"`
	CreatedAt      time.Time             `json:"created_at"`
	Errors         []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPServerList : Contain list of AdaptiveIPServers
type AdaptiveIPServerList struct {
	AdaptiveIPServers []AdaptiveIPServer    `json:"adaptiveip_server_list"`
	Errors            []errconv.PiccoloHccError `json:"errors"`
}

// AdaptiveIPServerNum - ish
type AdaptiveIPServerNum struct {
	Number int                   `json:"number"`
	Errors []errconv.PiccoloHccError `json:"errors"`
}
