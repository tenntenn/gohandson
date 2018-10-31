package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func selectBloodType(api *slack.Client, evt *slackevents.AppMentionEvent) error {
	attachment := slack.Attachment{
		Text:       "血液型を教えてください",
		CallbackID: "select_blood_type",
		Actions: []slack.AttachmentAction{
			{
				Name: "select",
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{Text: "A", Value: "A"},
					{Text: "B", Value: "B"},
					{Text: "O", Value: "O"},
					{Text: "AB", Value: "AB"},
				},
			},

			{
				Name:  "cancel",
				Text:  "キャンセル",
				Type:  "button",
				Style: "danger",
			},
		},
	}

	postParams := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			attachment,
		},
	}

	_, _, err := api.PostMessage(evt.Channel, "★★★血液型占い★★★", postParams)
	if err != nil {
		return err
	}

	return nil
}

func interactionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	slack.SetHTTPClient(urlfetch.Client(ctx))

	defer r.Body.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r.Body)
	body := buf.String()

	verifytoken := os.Getenv("SLACK_VERIFY_TOKEN")

	jsonStr, err := url.QueryUnescape(string(body)[8:])
	if err != nil {
		log.Errorf(ctx, "%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Errorf(ctx, "%v: %s", err, jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if message.Token != verifytoken {
		log.Errorf(ctx, "%v: %s", err, message.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	action := message.Actions[0]
	switch action.Name {
	case "select":
		result := bloodTypeUranai(action.SelectedOptions[0].Value)
		if err := responseMessage(w, message.OriginalMessage, "★占い結果★", result); err != nil {
			log.Errorf(ctx, "Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "cancel":
		title := fmt.Sprintf(":x: @%s キャンセルされました", message.User.Name)
		if err := responseMessage(w, message.OriginalMessage, title, ""); err != nil {
			log.Errorf(ctx, "Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	default:
		log.Errorf(ctx, "不正なアクション: %s", action)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func responseMessage(w http.ResponseWriter, msg slack.Message, title, value string) error {
	msg.Attachments[0].Actions = []slack.AttachmentAction{}
	msg.Attachments[0].Fields = []slack.AttachmentField{
		{
			Title: title,
			Value: value,
			Short: false,
		},
	}
	msg.ResponseType = "in_channel"

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&msg); err != nil {
		return err
	}
	return nil
}

func bloodTypeUranai(bloodType string) string {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	bts := map[string]int64{"A": 1, "B": 2, "AB": 3, "O": 4}
	rnd := rand.New(rand.NewSource(today.Unix() * bts[bloodType]))

	var buf bytes.Buffer

	fmt.Fprintln(&buf, "今日の運勢は...")
	switch rnd.Intn(6) {
	case 0:
		fmt.Fprintln(&buf, "残念...:fearful:「凶」です。")
	case 1, 2:
		fmt.Fprintln(&buf, "「吉」です。")
	case 3, 4:
		fmt.Fprintln(&buf, "「中吉」です。")
	case 5:
		fmt.Fprintln(&buf, "おめでとうございます:tada:「大吉」です。")
	}
	fmt.Fprintln(&buf)

	colors := []string{"赤", "青", "黄色", "緑", "黒", "ピンク"}
	fmt.Fprintf(&buf, "今日のラッキーカラーは%sです。", colors[rnd.Intn(len(colors))])
	fmt.Fprintln(&buf, "今日も一日がんばりましょう:muscle:")

	return buf.String()
}
