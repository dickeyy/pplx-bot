package services

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sgaunet/perplexity-go/v2"
)

var PPLX *perplexity.Client

func ConnectPerplexity() {
	PPLX = perplexity.NewClient(os.Getenv("PPLX_API_KEY"))
	log.Info().Msg("Connected to Perplexity")

	// temp, ping the api to see if its working
	res, err := PPLX.SendCompletionRequest(perplexity.NewCompletionRequest(
		perplexity.WithMessages([]perplexity.Message{
			{
				Role:    "user",
				Content: "Hello, how are you?",
			},
		}),
	))
	if err != nil {
		log.Error().Msgf("Error pinging Perplexity: %s", err.Error())
	}
	log.Info().Msgf("Perplexity API is working: %s", res.GetLastContent())
}
