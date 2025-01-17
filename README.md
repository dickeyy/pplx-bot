# PPLX Bot

A simple Discord bot for searching with [Perplexity](https://perplexity.ai/).

Currently, the bot is private, meaning you'll have to host your own instance, however, I plan to make it public soon.

## Hosting

To host an instance of the bot, you'll need to follow these steps:

### Prerequisites

- Go 1.23+
- Docker (optional)
- Git
- A Discord bot token and app ID - you can get one [here](https://discord.com/developers/applications)
- A Perplexity API key - you can get one [here](https://www.perplexity.ai/settings/api)

### Setup

First, clone the repository:

```bash
git clone https://github.com/dickeyy/pplx-bot.git
cd pplx-bot
```

Next, create a `.env` file in the root directory of the project, and add the following variables:

```txt
ENVIORNMENT=<prod|dev>
TOKEN=Bot <your token>
DEV_TOKEN=Bot <your dev token> # Optional
APP_ID=<your app id>
DEV_APP_ID=<your dev app id> # Optional
REGISTER_CMDS=<true|false> # Tells the bot whether to register slash commands (needed on first run)
PPLX_API_KEY=<your perplexity api key>
```

_Note: The `TOKEN` and `DEV_TOKEN` need the `Bot` prefix before the actual token._

### Running

Finally, build and run the bot, either using Docker or locally:

**Docker:**

```bash
docker build -t pplx-bot .
docker run -d --name pplx-bot --env-file .env pplx-bot
```

**Locally:**

```bash
go build -o pplx-bot .
./pplx-bot
```

Now you can invite the bot to your server and start using it!

## Commands

The bot only has one command: `/search <query> [debug]`, where `query` is the search query and `debug` is an optional boolean to send back some debug information.

## Contributing

Contributions are welcome! Please open an issue for discussion before working on a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
