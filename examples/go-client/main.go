// Example: register a bundle and resolve a requirement set against the
// control plane using the public SDK.
//
// Run a control plane first:
//
//	cd backend && go run ./cmd/syndovela-api
//
// Then:
//
//	go run ./go-client/main.go --server http://localhost:8080
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	syndovela "github.com/axisrobo/syndovela-open/sdk/go"
)

func main() {
	server := flag.String("server", "http://localhost:8080", "control-plane base URL")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := syndovela.New(*server)

	health, err := client.Health(ctx)
	if err != nil {
		log.Fatalf("health: %v", err)
	}
	fmt.Printf("connected to %s %s\n", health.Product, health.Version)

	bundle := syndovela.Bundle{
		APIVersion: "syndovela.axisrobo.io/v1",
		Kind:       "Bundle",
		Metadata: syndovela.Metadata{
			ID:        "example.identity-core",
			Version:   "0.1.0",
			Publisher: "example",
		},
		Skills: []syndovela.Skill{
			{ID: "identity.lookup", Contract: "identity.lookup/v2", Implementation: "0.1.0"},
		},
		Runtime: syndovela.Runtime{
			Protocol:  "sbrp/v1",
			ABI:       []string{"wasi/preview2"},
			Isolation: []string{"wasm"},
		},
		Security: syndovela.Security{
			Signature:  "required",
			SBOM:       "required",
			Provenance: "required",
		},
	}

	if _, err := client.RegisterBundle(ctx, bundle); err != nil {
		log.Fatalf("register: %v", err)
	}
	fmt.Printf("registered %s@%s\n", bundle.Metadata.ID, bundle.Metadata.Version)

	// Resolving now will fail: the bundle is only REGISTERED, and
	// registry presence is not approval. That failure is the point of
	// the example.
	_, err = client.Resolve(ctx, map[string]any{
		"skills": []map[string]string{{"ref": "identity.lookup/v2"}},
	})
	if apiErr, ok := err.(*syndovela.APIError); ok && apiErr.StatusCode == http.StatusUnprocessableEntity {
		fmt.Println("resolution correctly refused: a REGISTERED bundle is not deployable")
		return
	}
	if err != nil {
		log.Fatalf("resolve: %v", err)
	}
	fmt.Println("resolved")
}
