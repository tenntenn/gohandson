package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/events", eventsHandler)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	token := os.Getenv("SLACK_BOT_TOKEN")
	slack.SetHTTPClient(urlfetch.Client(ctx))
	api := slack.New(token)

	defer r.Body.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	body := buf.String()

	verifytoken := os.Getenv("SLACK_VERIFY_TOKEN")
	opt := slackevents.OptionVerifyToken(&slackevents.TokenComparator{verifytoken})
	evt, err := slackevents.ParseEvent(json.RawMessage(body), opt)
	if err != nil {
		log.Errorf(ctx, "ParseEvent: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if evt.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Errorf(ctx, "%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, r.Challenge)
		return
	}

	log.Infof(ctx, "Event:%#v", evt)
	if evt.Type == slackevents.CallbackEvent {
		var postParams slack.PostMessageParameters
		switch evt := evt.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			_, _, err := api.PostMessage(evt.Channel, "こんにちは", postParams)
			if err != nil {
				log.Errorf(ctx, "%v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
