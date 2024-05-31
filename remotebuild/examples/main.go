package main

import (
	"context"
	"log"
	"os"

	pbbuild "github.com/streamingfast/substreams/remotebuild/pb/sf/remotebuild/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Small GRPC Client that connects to a local remote build server and sends a build request
func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("go run main.go <zipped_source_code>")
		return
	}

	filepath := args[1]

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return
	}

	b, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
		return
	}

	client := pbbuild.NewBuildServiceClient(conn)
	buildResponse, err := client.Build(context.Background(), &pbbuild.BuildRequest{
		SourceCode:     b,
		Env:            []string{"ENV1=test", ""},
		CollectPattern: "substreams.spkg",
	})
	if err != nil {
		log.Fatalf("failed to build: %v", err)
		return
	}

	log.Printf("Build response - Logs: %v", buildResponse.Logs)
	log.Printf("Build response - Artifacts: %v", len(buildResponse.Artifacts))

	for _, artifact := range buildResponse.Artifacts {
		err = os.WriteFile(artifact.Filename, artifact.Content, 0644)
		if err != nil {
			log.Fatalf("failed to write file: %v", err)
			return
		}
	}
}
