package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dickeyy/perpbot/components"
	"github.com/dickeyy/perpbot/services"
	"github.com/rs/zerolog/log"
	"github.com/sgaunet/perplexity-go/v2"
)

func init() {
	services.Commands[searchCmd.Name] = &services.Command{
		ApplicationCommand: searchCmd,
		Handler:            handleSearch,
	}
}

var searchCmd = &discordgo.ApplicationCommand{
	Type:        discordgo.ChatApplicationCommand,
	Name:        "search",
	Description: "Query Perplexity",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "query",
			Description: "The query to search for",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionBoolean,
			Name:        "debug",
			Description: "Enable debug mode, returns some more stats about the request",
			Required:    false,
		},
	},
}

func handleSearch(s *discordgo.Session, i *discordgo.InteractionCreate) *discordgo.InteractionResponse {
	query := i.ApplicationCommandData().Options[0].StringValue()
	debug := false
	if len(i.ApplicationCommandData().Options) > 1 {
		debug = i.ApplicationCommandData().Options[1].BoolValue()
	}

	if query == "" {
		return EmbedResponse(components.ErrorEmbed("Please provide a query"), true)
	}

	go func() {
		// start a timer
		start := time.Now()

		msg := []perplexity.Message{
			{
				Role:    "system",
				Content: "You are a Discord bot that can answer questions. Keep your answers short and concise. Do not include any links or images in your answer. You can use properly formatted Markdown in your answers if you want.",
			},
			{
				Role:    "user",
				Content: query,
			},
		}

		req := perplexity.NewCompletionRequest(
			perplexity.WithMessages(msg),
			perplexity.WithReturnImages(false),
			perplexity.WithDefaultModel(),
			perplexity.WithMaxTokens(500),
		)
		err := req.Validate()
		if err != nil {
			log.Error().AnErr("Error validating request", err)
			errEmbed := components.ErrorEmbed("Failed to validate request\n```" + err.Error() + "```")
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{errEmbed},
			})
			return
		}

		res, err := services.PPLX.SendCompletionRequest(req)
		if err != nil {
			log.Error().AnErr("Error sending request", err)
			errEmbed := components.ErrorEmbed("Failed to send request\n```" + err.Error() + "```")
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{errEmbed},
			})
			return
		}

		cont := res.GetLastContent()

		// end the timer
		end := time.Now()

		desc := fmt.Sprintf("-# %s asks \n> *\"%s\"*\n\n-# perplexity.ai says\n %s\n\n-# Answered in %s", i.Member.User.Username, query, cont, end.Sub(start).String())
		if debug {

			debugInfo := "```asciidoc\n" +
				"Model		 :: " + res.Model + "\n" +
				"Created	   :: " + strconv.Itoa(res.Created) + "\n" +
				"Comp. Tokens  :: " + strconv.Itoa(res.Usage.CompletionTokens) + "\n" +
				"Prompt Tokens :: " + strconv.Itoa(res.Usage.PromptTokens) + "\n" +
				"Total Tokens  :: " + strconv.Itoa(res.Usage.TotalTokens) + "\n" +
				"```"
			desc += fmt.Sprintf("\n%s", debugInfo)
		}

		embed := components.NewEmbed().
			SetDescription(desc).
			SetColor("DarkButNotBlack")

		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embed.MessageEmbed},
		})

	}()

	return LoadingResponse()
}
