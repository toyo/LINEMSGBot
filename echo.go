package linemsgbot

import (
	"os"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"github.com/line/line-bot-sdk-go/linebot"
)


func init(){
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		c := appengine.NewContext(req)

		bot, err := linebot.New(
			os.Getenv("CHANNEL_SECRET"),
			os.Getenv("CHANNEL_TOKEN"),
			linebot.WithHTTPClient(urlfetch.Client(c)),
		)

		if err != nil {
			log.Criticalf(c,"linebot init error")
			log.Criticalf(c,err.Error())
			w.WriteHeader(500)
			return
		}

		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					source := event.Source
					if source.Type == linebot.EventSourceTypeUser {
						log.Debugf(c,string(message.Text))
						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage(message.Text)).Do(); err != nil {
							log.Debugf(c,string(err.Error()))
						}
					}
				default:
					source := event.Source
					if source.Type == linebot.EventSourceTypeUser {
						log.Debugf(c,"Got message!")
						if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage("Got message!")).Do(); err != nil {
							log.Debugf(c,string(err.Error()))
						}
					}
				}
			case linebot.EventTypePostback:
				source := event.Source
				if source.Type == linebot.EventSourceTypeUser {
					log.Debugf(c,"Got PostBack?!")
					if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage("Got PostBack?!")).Do(); err != nil {
						log.Debugf(c,string(err.Error()))
					}
				}
			case linebot.EventTypeBeacon:
				source := event.Source
				if source.Type == linebot.EventSourceTypeUser {
					log.Debugf(c,"Got beacon")
					if _, err = bot.PushMessage(source.UserID, linebot.NewTextMessage("Got beacon!")).Do(); err != nil {
						log.Debugf(c,string(err.Error()))
					}
				}
			default:
			}
		}
		w.WriteHeader(http.StatusOK)  
	})
	http.ListenAndServe(":8080", nil)
}
