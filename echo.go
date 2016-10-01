// Copyright 2016 LINE Corporation
// Copyright 2016 FUJIURA Toyonori
// This is branched from https://github.com/line/line-bot-sdk-go/blob/master/examples/echo_bot/server.go
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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
