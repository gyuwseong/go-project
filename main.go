package main

import (
	"log"
	"net/http"

	campaignv1connect "go-project/gen/campaign/v1/campaignv1connect"
	"go-project/server"
)

func main() {
	mux := http.NewServeMux()
	campaignServer := server.NewCampaignServer()
	path, handler := campaignv1connect.NewCampaignServiceHandler(campaignServer)
	mux.Handle(path, handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
