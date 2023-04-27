package cronjobs

import (
	"os"

	"github.com/Toaaa/archiv-backend/pkg/external_apis"
	"github.com/Toaaa/archiv-backend/pkg/logger"
	"github.com/Toaaa/archiv-backend/pkg/models"
	"github.com/Toaaa/archiv-backend/pkg/queries"
)

func SetStreamStatus() {
	var streams external_apis.TwitchStreamResponse
	if err := external_apis.TwitchGetHelixStreams(&streams); err != nil {
		logger.Error.Println(err)
		return
	}
	isLive := len(streams.Data) > 0

	var settings models.Settings
	if err := queries.GetSettings(&settings); err != nil {
		logger.Error.Println(err)
		return
	}
	err := external_apis.UpdateBearer(&settings)
	if err != nil {
		logger.Error.Println(err)
		return
	}

	if isLive != settings.IsLive {
		isLiveSettings := models.Settings{IsLive: true}
		if err := queries.PartiallyUpdateSettings(&isLiveSettings); err != nil {
			logger.Error.Println(err)
		}
		if isLive {
			logger.Debug.Println("[cronjob] stream live")
			if os.Getenv("DISCORD_WEBHOOK") != "" {
				if err := external_apis.DiscordSendWebhook(streams); err != nil {
					logger.Error.Println(err)
					return
				}
			}
		} else {
			settings.IsLive = false
			if err := queries.OverwriteAllSettings(&settings); err != nil {
				logger.Error.Println(err)
				return
			}
		}
	}
}
