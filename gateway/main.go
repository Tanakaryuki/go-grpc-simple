package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Tanakaryuki/go-grpc-simple/proto"
	"google.golang.org/grpc"
)

func main() {
	// helloサービスへの接続先（docker-composeではサービス名で解決される）
	helloServiceAddress := os.Getenv("HELLO_SERVICE_ADDR")
	if helloServiceAddress == "" {
		helloServiceAddress = "hello:50051"
	}

	// gRPCクライアントの作成
	conn, err := grpc.Dial(helloServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to hello service: %v", err)
	}
	defer conn.Close()
	helloClient := proto.NewHelloServiceClient(conn)

	// HTTPハンドラ：/hello
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		req := &proto.HelloRequest{}
		res, err := helloClient.SayHello(ctx, req)
		if err != nil {
			http.Error(w, "Error calling hello service", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, res.Message)
	})

	// HTTPハンドラ：/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Gateway server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
