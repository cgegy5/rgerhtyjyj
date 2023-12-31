//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"time"

	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type TimeReference struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	TimeRef     time.Time
	Source      string
}
