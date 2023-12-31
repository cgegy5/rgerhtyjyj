//autogenerated:yes
//nolint:revive,lll
package audio_common_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type AudioDataStamped struct {
	msg.Package `ros:"audio_common_msgs"`
	Header      std_msgs.Header
	Audio       AudioData
}
