package main

import (
	"net"

	protos "github.com/AbDwivedi7/grpc-go/protos/currency"
	"github.com/AbDwivedi7/grpc-go/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gc := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gc, cs)
	reflection.Register(gc)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "Error", err)
	}
	gc.Serve(l)
}
