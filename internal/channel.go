package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	channelFile = "result.json"
	typeMessage = "message"
)

type Channel struct {
	Messages []MixedPost
}

func (c Channel) GetPosts() []Post {
	posts := []Post{}

	for _, post := range c.Messages {
		if post.Type != typeMessage || post.Text == "" {
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

type MixedPost struct {
	BasePost
	Text any // string | (string | { text: string })[]
}

func (m MixedPost) GetText() string {
	switch text := m.Text.(type) {
	case string:
		return text
	case []any:
		var combinedText string

		for _, textItem := range text {
			str, ok := textItem.(string)
			if ok {
				combinedText += str
				continue
			}

			for key, value := range textItem.(map[string]any) {
				if key == "text" {
					combinedText += value.(string)
				}
			}
		}

		return combinedText
	}

	fmt.Println(separator)
	fmt.Println(m.Text)
	ThrowError("couldn't parse text")
	return ""
}
