package internal

import (
	"encoding/json"
	"os"
)

const (
	channelFile = "result.json"
	typeMessage = "message"
)

type Channel struct {
	Messages []PostJson
}

func (c Channel) GetPosts() []Post {
	posts := []Post{}

	for _, post := range c.Messages {
		if (post.Type != typeMessage) || (len(post.TextEntities) == 0) {
			continue
		}

		posts = append(posts, Post{
			BasePost: post.BasePost,
			Text:     post.GetText(),
		})
	}

	return posts
}

func ReadChannelJson() Channel {
	channelJson, fileErr := os.ReadFile(channelFile)
	CheckError(fileErr, "error when opening channel json")

	var channelJsonParsed Channel
	jsonErr := json.Unmarshal(channelJson, &channelJsonParsed)
	CheckError(jsonErr, "error when parsing channel json")

	return channelJsonParsed
}

type PostJson struct {
	BasePost
	TextEntities []Text `json:"text_entities"`
}

type Text struct {
	Type string
	Text string
}

func (p PostJson) GetText() string {
	var text string

	for _, textItem := range p.TextEntities {
		text += textItem.Text
	}

	return text
}
