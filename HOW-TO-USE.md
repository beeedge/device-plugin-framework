# Converter Example

This example builds a simple key/converter store CLI where the mechanism
for storing and retrieving keys is pluggable. However, in this example we don't
trust the plugin to do the summation work. We use bi-directional plugins to
call back into the main proccess to do the sum of two numbers. To build this example:

```sh
# This builds the main CLI
$ go build -o converter

# This builds the plugin written in Go
$ go build -o converter-go-grpc ./plugin-go-grpc

# This tells the Converter binary to use the "converter-go-grpc" binary
$ export CONVERTER_PLUGIN="./converter-go-grpc"

# Read and write
$ ./converter put hello 1
$ ./converter put hello 1

$ ./converter get hello
2
```

### Plugin: plugin-go-grpc

This plugin uses gRPC to serve a plugin that is written in Go:

```
# This builds the plugin written in Go
$ go build -o converter-go-grpc ./plugin-go-grpc

# This tells the KV binary to use the "kv-go-grpc" binary
$ export CONVERTER_PLUGIN="./converter-go-grpc"
```

## Updating the Protocol

If you update the protocol buffers file, you can regenerate the file
using the following command from this directory. You do not need to run
this if you're just trying the example.

For Go:

```sh
$ protoc -I proto/ proto/kv.proto --go_out=plugins=grpc:proto/
```
