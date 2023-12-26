
# Discord Greeting Bot (Golang)

## Description
This Discord bot, written in Golang, is a simple and friendly greeting bot that responds when a user says "hello" in a server. The bot aims to create a welcoming atmosphere by acknowledging users who initiate a conversation with a greeting.

## Features
- Responds to the command "!hello" with a friendly greeting.
- Supports multiple server deployments.

## Installation
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/greetbot-discord.git
   cd greetbot-discord
   ```

2. **Build the Bot:**
   ```bash
   go build
   ```

3. **Configure the Bot:**
   - Create a new Discord application and obtain the bot token from the [Discord Developer Portal](https://discord.com/developers/applications).
   - Copy the `config.example.json` file to `config.json` and replace `"YOUR_BOT_TOKEN"` with your actual bot token.

4. **Run the Bot:**
   ```bash
   ./greetbot-discord
   ```

## Usage
1. Invite the bot to your Discord server using the invite link generated when you created your Discord application.
2. Ensure the bot has the necessary permissions to read messages and send messages in the desired channels.
3. In a text channel, type "!hello" to trigger the bot's greeting response.

## Configuration
- `config.json`: Contains the bot configuration, including the bot token.

## Contributing
If you'd like to contribute to the development of this bot, feel free to submit pull requests. Make sure to follow the existing coding style and include comprehensive commit messages.

## Issues
If you encounter any issues or have suggestions for improvements, please open an issue on the GitHub repository.

## License
This Discord Greeting Bot (Golang Version) is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
