package v1

import "time"

var example = Changes{
	Count:         1,
	LastTimestamp: time.Now(),
	Scope: []Scope{
		{
			Type:       "Namespace",
			Identifier: Identifier{Name: "kube-system"},
		},
		{
			Type:       "Cluster",
			Identifier: Identifier{Name: "localhost"},
		},
		{
			Type:       "Label",
			Identifier: Identifier{Id: "region", Name: "eu-west-1"},
		},
	},
	Changes: []Change{
		{
			Type:        "Canary",
			Category:    "Succeeded",
			Description: "ping timeout exceeded 5ms > 1ms",
			Identifier: Identifier{
				Id:   "check.key",
				Name: "Ping google.com",
			},
		},
	},
}
