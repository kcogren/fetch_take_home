package main

import (
	"fetch_api/api/handler"
	"log/slog"
	"net/http"
	"os"
)

// Developed on Windows by KC Ogren
// Trying out standard library mux changes and slog changes
func main() {
	logLevel := &slog.LevelVar{} // INFO
	opts := &slog.HandlerOptions{

		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	logLevel.Set(slog.LevelInfo)

	slog.SetDefault(logger)

	slog.Info("Receipt Processor Starting")

	// Persistance
	tempStorage := make(map[string]uint64)
	receiptHandeler := &handler.ReceiptHandler{Storage: tempStorage}

	// Request in
	mux := http.NewServeMux()

	// Route Handler
	mux.Handle("/", &handler.HomeHandler{})

	// Endpoint: Process Receipts
	// Path: /receipts/process
	// Method: POST
	// Payload: Receipt JSON
	// Response: JSON containing an id for the receipt.
	mux.HandleFunc("POST /receipts/process", receiptHandeler.Post)

	// Endpoint: Get Points
	// Path: /receipts/{id}/points
	// Method: GET
	// Response: A JSON object containing the number of points awarded.
	mux.HandleFunc("GET /receipts/{id}/points", receiptHandeler.GetPoints)

	slog.Info("Receipt Processor Ready")
	http.ListenAndServe(":8080", mux)

}
