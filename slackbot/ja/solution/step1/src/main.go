package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func main() {
	http.HandleFunc("/events", eventsHandler)
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)

	defer r.Body.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	body := buf.String()

	verifytoken := os.Getenv("SLACK_VERIFY_TOKEN")
	opt := slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: verifytoken})
	evt, err := slackevents.ParseEvent(json.RawMessage(body), opt)
	if err != nil {
		log.Printf("ParseEvent: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if evt.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, r.Challenge)
		return
	}

	log.Printf("Event:%#v", err)
	if evt.Type == slackevents.CallbackEvent {
		var postParams slack.PostMessageParameters
		switch evt := evt.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			_, _, err := api.PostMessage(evt.Channel, "こんにちは", postParams)
			if err != nil {
				log.Printf("%v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
