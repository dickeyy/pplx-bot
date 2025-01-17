package components

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var Components = make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate) *discordgo.InteractionResponse)

func OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	handler, ok := Components[data.CustomID]
	if !ok {
		return
	}

	start := time.Now()
	resp := handler(s, i)
	if resp != nil && time.Since(start) <= 3*time.Second {
		s.InteractionRespond(i.Interaction, resp)
	}
}

func ComplexResponse(d *discordgo.InteractionResponseData) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	}
}

func ContentResponse(c string, e bool) *discordgo.InteractionResponse {
	d := &discordgo.InteractionResponseData{Content: c}
	if e {
		d.Flags = discordgo.MessageFlagsEphemeral
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: d,
	}
}

func EmptyResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	}
}

func UpdateResponse(i *discordgo.InteractionResponseData) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: i,
	}
}
