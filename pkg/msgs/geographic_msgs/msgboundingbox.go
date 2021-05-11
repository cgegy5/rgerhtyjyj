//nolint:golint
package geographic_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type BoundingBox struct {
	msg.Package `ros:"geographic_msgs"`
	MinPt       GeoPoint
	MaxPt       GeoPoint
}