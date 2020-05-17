simple wrapper to send text file content to discord channel

main work is done by https://github.com/bwmarrin/discordgo

you can install it with go get -v github.com/codecare/discordbot/cmd/discordbot

this will install an executable in your default path for go executables $GOPATH/bin (default GOPATH=$HOME/go)

discord bot needs:

	export DISCORD_TOKEN="***"
	export DISCORD_CHANNEL_ID="***"

token should be valid for all channels of server

beware of maximum character limit of discord

to set up channel for bot usage and token see the internet

1. use web portal https://discordapp.com/developers
2. define application
3. define bot
4. write down client id and secret
5. enable bot in browser
   https://discordapp.com/api/oauth2/authorize?scope=bot&client_id=12345
6. export token as DISCORD_TOKEN
7. to get the channel id enable developer mode for discord user
   then right click on channel to get id
8. export channel id as DISCORD_CHANNEL_ID

then run with file name of data to transfer

	discordbot sendthis.txt



    