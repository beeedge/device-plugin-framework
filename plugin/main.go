package main

import (
	"encoding/json"

	da "github.com/beeedge/beethings/pkg/device-access/rest/models"
	"github.com/beeedge/device-plugin-framework/shared"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of device-plugin.
type Converter struct {
}

// ConvertIssueMessage2Device converts issue request to protocol that device understands, which has four return parameters:
// 1. inputMessages: device issue protocols for each of command input param.
// 2. outputMessages: device data report protocols for each of command output param.
// 3. issueTopic: device issue MQTT topic for input params.
// 4. issueResponseTopic: device issue response MQ topic for output params.
func (c *Converter) ConvertIssueMessage2Device(deviceId, modelId, featureId string, values map[string]string, convertedDeviceFeatureMap string) ([]string, []string, string, string, error) {
	var deviceFeatureMap da.DeviceFeatureMap
	if err := json.Unmarshal([]byte(convertedDeviceFeatureMap), &deviceFeatureMap); err != nil {
		return nil, nil, "", "", err
	}
	// TODO: concrete implement
	return nil, nil, "", "", nil
}

// ConvertDeviceMessages2MQFormat receives device command issue responses and converts it to RabbitMQ normative format.
func (c *Converter) ConvertDeviceMessages2MQFormat(messages []string, convertedDeviceFeatureMap string) (string, []byte, error) {
	var deviceFeatureMap da.DeviceFeatureMap
	if err := json.Unmarshal([]byte(convertedDeviceFeatureMap), &deviceFeatureMap); err != nil {
		return "", nil, err
	}
	// TODO: concrete implement
	return "Hello World", nil, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"converter": &shared.ConverterPlugin{Impl: &Converter{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
