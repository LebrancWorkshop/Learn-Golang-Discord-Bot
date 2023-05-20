package server

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/LebrancWorkshop/Learn-Golang-Discord-Bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	BotToken       = flag.String("token", "", "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type IDiscordServer interface {
	Start()
}

type discordServer struct {
	cfg 				config.IConfig
	dg 					*discordgo.Session
	commands 		[]*discordgo.ApplicationCommand
}

func NewDiscordServer(cfg config.IConfig) IDiscordServer {
	dg, err := discordgo.New("Bot " + cfg.App().GetToken())
	if err != nil {
		log.Fatalf("Invalid Bot Parameters: %s", err)
	}
	return &discordServer{
		cfg: cfg,
		dg: dg,
		commands: make([]*discordgo.ApplicationCommand, 0),
	}
}

// Code From: https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go 
func (s *discordServer) Start() {
	s.dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.dg.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(s.commands))
	for i, v := range s.commands {
		cmd, err := s.dg.ApplicationCommandCreate(s.dg.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.dg.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.dg.ApplicationCommandDelete(s.dg.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}

