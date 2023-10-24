package bot

import (
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/arno4000/schaebigctl/pkg/ai"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartBot(token string) {
	logrus.Infoln("Started schaebigctl bot")
	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		logrus.Errorln(err)
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}

		if strings.Contains(strings.ToLower(m.Content), "marcus") {
			logrus.Infof("Received message in channel %s: %s", m.ChannelID, m.Content)
			prompts := []string{
				`Your persona: Your name is Marcus. You are an immortal adventurer in the MMO Amazon Game New World, you've lived in Aeternum for a long time and are very wise. You love a good fight. Make your answer short. Here's what the player says to you:`,
				`Your persona: Your name is Marcus. You are an immortal adventurer in the MMO Amazon Game New World, you've lived in Aeternum for a long time and are very wise. Compare one of the feats you accomplished to whatever the player tells you in a very exaggerated way. Make your answer short. Here's what the player says to you:`,
			}
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			randomNumber := r.Intn(2)
			prompt := prompts[randomNumber]
			timeout := viper.GetInt("maxTimeoutMinutes")
			var waitDuration time.Duration
			if timeout > 0 {
				waitDuration = time.Duration(rand.Intn(viper.GetInt("maxTimeoutMinutes")))
				logrus.Infof("Posting message to channel %s in %d minutes", m.ChannelID, waitDuration)
				time.Sleep(waitDuration * time.Minute)
			} else {
				logrus.Infof("Posting message to thread %s now", m.ChannelID)
			}

			answer, err := ai.GenAIResponse(prompt, m.Content)
			if err != nil {
				logrus.Errorln(err)
			}
			logrus.Infof("Posted message to thread %s, took %d minutes", m.ChannelID, waitDuration)
			_, err = s.ChannelMessageSend(m.ChannelID, answer)
			if err != nil {
				logrus.Errorln(err)
			}
		}

	})

	err = dg.Open()
	if err != nil {
		logrus.Errorln(err)
	}

	// // Get the list of guilds the bot is a member of
	// guilds, err := dg.UserGuilds(100, "", "")
	// if err != nil {
	// 	logrus.Error("Error fetching guilds: ", err)
	// 	return
	// }

	// // Leave all guilds
	// for _, guild := range guilds {
	// 	err := dg.GuildLeave(guild.ID)
	// 	if err != nil {
	// 		logrus.Error("Error leaving guild ", guild.ID, ": ", err)
	// 	} else {
	// 		logrus.Info("Left guild: ", guild.Name)
	// 	}
	// }

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGILL, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
	logrus.Infoln("Exiting schaebigctl bot. Bye. Stay schaebig!")

}
