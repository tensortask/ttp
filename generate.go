//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tensortask/boltz/internal/log"
)

var (
	goPath = os.Getenv("GOPATH")
)

func main() {
	if err := generateProtos(); err != nil {
		fmt.Printf("%+v", err)
	}
}

func generateProtos() error {
	args := []string{
		"-I=protos",
		fmt.Sprintf("-I=%s", filepath.Join(goPath, "src")),
		fmt.Sprintf("-I=%s", filepath.Join(goPath, "src", "github.com", "gogo", "protobuf", "protobuf")),
		fmt.Sprintf("--proto_path=%s", filepath.Join(goPath, "src", "github.com")),
		"--gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:gen",
		"tensor.proto",
		"transport.proto",
		"type.proto",
	}
	cmd := exec.Command("protoc", args...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	log.Info().Msg("Successfully compiled TTP protocol buffers.")
	return nil
}
