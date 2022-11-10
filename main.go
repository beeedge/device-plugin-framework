package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	da "github.com/beeedge/beethings/pkg/device-access/rest/models"
	"github.com/beeedge/device-plugin-framework/shared"
	"github.com/hashicorp/go-plugin"
)

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("DEVICE_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("converter")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a Converter store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	converter := raw.(shared.Converter)
	var convertedDeviceFeatureMap da.DeviceFeatureMap
	convertedDeviceFeatureMapBytes, err := json.Marshal(&convertedDeviceFeatureMap)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
	result, _, err := converter.ConvertDeviceMessages2MQFormat([]string{""}, string(convertedDeviceFeatureMapBytes))
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
}
