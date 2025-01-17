package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

type Config struct {
	DevGuildID   string `json:"dev_guild_id"`
	ErrMsgPrefix string `json:"error_msg_prefix"`
	BotOwnerID   string `json:"bot_owner_id"`
}

var Bot Config

func init() {
	// This can also be replaced with code to read from a file or environment variables
	Bot = Config{
		DevGuildID:   "val",
		ErrMsgPrefix: "val",
		BotOwnerID:   "val",
	}

	// Optionally, load config from a JSON file
	file, err := os.Open("config.json")
	if err == nil {
		defer file.Close()
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Bot)
		if err != nil {
			log.Error().AnErr("Failed to decode config.json", err)
		}
		log.Info().Msg("Loaded config.json")
	} else {
		log.Error().Msg("Failed to open config.json")
	}
}
