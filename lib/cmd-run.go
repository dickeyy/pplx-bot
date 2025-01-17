package lib

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func CmdRun(s *discordgo.Session, i *discordgo.InteractionCreate, d time.Duration) {
	data := i.ApplicationCommandData()
	log.Info().Str("command", data.Name).Str("guild", i.GuildID).Str("user", i.Member.User.ID).Str("took", d.String()).Msg("Command executed")
}
