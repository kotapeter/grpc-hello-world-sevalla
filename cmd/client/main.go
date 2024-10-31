package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cshep4/grpc-course/grpc-hello-world-sevalla/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	host, ok := os.LookupEnv("GRPC_HOST")
	if !ok {
		host = "localhost:50051"
	}
	conn, err := grpc.NewClient(host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloServiceClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := client.SayHello(ctx, &proto.SayHelloRequest{Name: "Chris"})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		// return file contents to user
		if _, err := w.Write([]byte(res.GetMessage())); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	log.Printf("starting http server on address: :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
