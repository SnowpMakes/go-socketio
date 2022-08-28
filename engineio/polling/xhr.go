package polling

import (
	"github.com/SnowpMakes/socketio/engineio/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}
