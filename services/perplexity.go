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
}
