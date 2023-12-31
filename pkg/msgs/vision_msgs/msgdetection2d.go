//autogenerated:yes
//nolint:revive,lll
package vision_msgs

import (
	"github.com/bluenviron/goroslib/v2/pkg/msg"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/sensor_msgs"
	"github.com/bluenviron/goroslib/v2/pkg/msgs/std_msgs"
)

type Detection2D struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Results     []ObjectHypothesisWithPose
	Bbox        BoundingBox2D
	SourceImg   sensor_msgs.Image
}
