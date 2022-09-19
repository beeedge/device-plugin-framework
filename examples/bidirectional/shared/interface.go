// Package shared contains shared data between the host and plugins.
package shared

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/device-plugin-framework/examples/bidirectional/proto"
	"github.com/hashicorp/go-plugin"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"converter": &ConverterPlugin{},
}

type AddHelper interface {
	Sum(int64, int64) (int64, error)
}

// KV is the interface that we're exposing as a plugin.
type Converter interface {
	ConvertReportMessage(deviceId, modelId, featureId, msgId string) ([]string, error)
	ConvertIssueMessage(deviceId, modelId, featureId, msgId string, values map[string]string) ([]string, []string, error)
	RevConvertMessages(messages []string) (string, []string, error)
	Put(key string, value int64, a AddHelper) error
	Get(key string) (int64, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type ConverterPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Converter
}

func (p *ConverterPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterConverterServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *ConverterPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewConverterClient(c),
		broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &ConverterPlugin{}
