package services

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/dickeyy/perpbot/config"
	"github.com/rs/zerolog/log"
)

var Discord *discordgo.Session
var Commands = make(map[string]*Command)
var RegisteredCommands = make([]*discordgo.ApplicationCommand, len(Commands))

type Command struct {
	*discordgo.ApplicationCommand
	Handler func(*discordgo.Session, *discordgo.InteractionCreate) *discordgo.InteractionResponse
}

func ConnectDiscord(events []interface{}) {
	token := ""
	env := os.Getenv("ENVIORNMENT")
	registerCmds := os.Getenv("REGISTER_CMDS")

	if env == "prod" {
		token = os.Getenv("TOKEN")
	} else {
		token = os.Getenv("DEV_TOKEN")
	}

	Discord, _ = discordgo.New(token)
	Discord.Identify.Intents = discordgo.Intent(discordgo.IntentsAllWithoutPrivileged)
	Discord.StateEnabled = true
	Discord.State.MaxMessageCount = 5000

	for _, h := range events {
		Discord.AddHandler(h)
	}

	err := Discord.Open()
	if err != nil {
		log.Fatal().AnErr("Error opening connection to Discord", err)
	}

	// register commands
	if registerCmds == "true" {
		if env == "prod" {
			log.Info().Msgf("Registering %d global commands", len(Commands))
			RegisterCommands(Discord, "") // register globally
		} else {
			log.Info().Msgf("Registering %d dev commands", len(Commands))
			RegisterCommands(Discord, config.Bot.DevGuildID) // just register for the dev guild
		}
	}
}

func RegisterCommands(s *discordgo.Session, g string) {
	i := 0
	for _, v := range Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, g, v.ApplicationCommand)
		if err != nil {
			log.Fatal().AnErr("Error registering command: "+v.Name, err)
		}
		i += 1
		log.Info().Msgf("Registered %d/%d commands", i, len(Commands))
	}
}

func DisconnectDiscord() {
	Discord.Close()
	log.Info().Msg("Disconnected from Discord")
}
