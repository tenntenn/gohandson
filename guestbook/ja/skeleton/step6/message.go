package main

import "time"

// Message はゲストブックに投稿されるメッセージです。
type Message struct {
	Name      string    `datastore:"name"`
	Text      string    `datastore:"text"`
	CreatedAt time.Time `datastore:"createdAt"`
}
