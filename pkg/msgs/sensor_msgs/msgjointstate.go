//autogenerated:yes
//nolint:revive,lll
package sensor_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type JointState struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Name        []string
	Position    []float64
	Velocity    []float64
	Effort      []float64
}
