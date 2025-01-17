package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/dickeyy/perpbot/events"
	"github.com/dickeyy/perpbot/services"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		return
	}

	// set log format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set up console writer
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Set global logger
	log.Logger = log.Output(consoleWriter)

	// set the environment
	env := os.Getenv("ENVIORNMENT")

	if env == "dev" {
		log.Warn().Msg("Running in development mode")
	}
}

func main() {
	services.ConnectDiscord(events.Events)
	services.ConnectPerplexity()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info().Msg("Press Ctrl+C to exit")

	// handle shutdown
	<-stop
	log.Warn().Msg("Shutting down")
	services.DisconnectDiscord()

	log.Info().Msg("Goodbye!")
}
