package main

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"

	"golang.org/x/net/http2"
	"net/http"
)

type chatService struct {
}

func (s *chatService) Join(c context.Context, r *JoinRequest) (*JoinResponse, error) {
	log.Printf("join: %s", r.GetName())
	return &JoinResponse{}, nil
}

func (s *chatService) Say(c context.Context, r *SayRequest) (*SayResponse, error) {
	log.Printf("say: %s", r.GetMes())
	return &SayResponse{}, nil
}

func main() {
	srv := grpc.NewServer()
	RegisterChatServer(srv, &chatService{})
	wsrv := grpcweb.WrapServer(srv)

	/*
		mux := http.NewServeMux()
		mux.HandleFunc("/", wsrv.ServeHTTP)
	*/

	var s http.Server
	http2.VerboseLogs = true
	s.Addr = ":5000"

	http2.ConfigureServer(&s, nil)

	http.HandleFunc("/", wsrv.ServeHTTP)

	panic(s.ListenAndServeTLS("server.crt", "server.key"))
}
