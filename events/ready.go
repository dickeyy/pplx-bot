package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func init() {
	Events = append(Events, onReady)
}

func onReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Info().Msgf("Signed in as %s", s.State.User.String())
}
