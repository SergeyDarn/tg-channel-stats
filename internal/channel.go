package internal

import (
	"encoding/json"
	"os"
	"strconv"
)

const (
	channelFile     = "result.json"
	postTypeMessage = "message"
	textTypeHashtag = "hashtag"
)

type Channel struct {
	Posts []PostJson `json:"Messages"`
}

func (c *Channel) ProcessedData() (Posts, HashtagMap) {
	posts := Posts{}
	hashtagMap := HashtagMap{}

	for _, post := range c.Posts {
		if (post.Type != postTypeMessage) || (len(post.TextEntities) == 0) {
			continue
		}

		text, hashtags := post.ProcessedText()
		postId := strconv.Itoa(post.Id)
		posts = append(posts, Post{
			BasePost:     post.BasePost,
			Id:           postId,
			ReactionsMap: post.ReactionsMap(),
			Text:         text,
			Hashtags:     hashtags,
		})

		if len(hashtags) == 0 {
			continue
		}

		for _, hashtag := range hashtags {
			postsIds := append(hashtagMap[hashtag].PostsIds, postId)

			reactionsMap := hashtagMap[hashtag].ReactionsMap
			if reactionsMap == nil {
				reactionsMap = ReactionsMap{}
			}

			hashtagMap[hashtag] = Hashtag{
				Text:         hashtag,
				ReactionsMap: reactionsMap.Add(post.Reactions),
				PostsIds:     postsIds,
			}
		}
	}

	return posts, hashtagMap
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
	Id           int
	TextEntities []TextEntity `json:"text_entities"`
	Reactions    []Reaction
}

type TextEntity struct {
	Type string
	Text string
}

func (p *PostJson) ProcessedText() (string, []string) {
	var text string
	var hashtags []string

	for _, textItem := range p.TextEntities {
		text += textItem.Text

		if textItem.Type == textTypeHashtag {
			hashtags = append(hashtags, textItem.Text)
		}
	}

	return text, hashtags
}

func (p *PostJson) ReactionsMap() ReactionsMap {
	reactionsMap := ReactionsMap{}
	return reactionsMap.Add(p.Reactions)
}
