//autogenerated:yes
//nolint:revive,lll
package std_srvs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
)

type TriggerReq struct {
	msg.Package `ros:"std_srvs"`
}

type TriggerRes struct {
	msg.Package `ros:"std_srvs"`
	Success     bool
	Message     string
}

type Trigger struct {
	msg.Package `ros:"std_srvs"`
	TriggerReq
	TriggerRes
}
