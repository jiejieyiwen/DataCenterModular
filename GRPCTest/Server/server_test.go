package Server

import (
	"testing"
)

const (
	PORT = ":50051"
)

func TestServer(t *testing.T) {
	server := GRpcServer{}
	server.Service(PORT)
}
