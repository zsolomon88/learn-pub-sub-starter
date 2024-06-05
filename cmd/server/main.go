package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
    connStr := "amqp://guest:guest@localhost:5672"
    conn, err := amqp.Dial(connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    fmt.Printf("Connection opened: %v", connStr)
	fmt.Println("Starting Peril server...")

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, os.Interrupt)
    <-signalChan

    fmt.Println("Shutting down")
    conn.Close()
}
