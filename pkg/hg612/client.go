package hg612

import "net/http"

//stats = map[string]stat{
//	"deviceInfo": {
//		path:   "html/status/deviceinfo.asp",
//	},
//	"ATM": {
//		path:   "html/status/atmStatus.asp",
//	},
//	"WAN": {
//		path:   "html/status/internetstatus.asp",
//	},
//	"VDSL": {
//		path:   "html/status/xdslStatus.asp",
//	},
//	"LAN": {
//		path:   "html/status/ethenet.asp",
//	},
//}

type Client interface {
	DSLStatus() (*VDSLStatus, error)
}

type HG612Client struct {
	basePath string
	client   *http.Client
}

func New(basePath string, client *http.Client) Client {
	return HG612Client{
		basePath: basePath,
		client:   client,
	}
}
