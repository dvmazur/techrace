package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/nlopes/slack"
)

var jokes []Joke

type Joke struct {
	Text string `json:"text"`
}

func init() {
	data, err := ioutil.ReadFile("jokes.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &jokes)
}

func sendRandom(api *slack.Client, channelName string, jokes []Joke) error {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	joke := jokes[r.Intn(len(jokes))]

	_, _, _, err := api.SendMessage(channelName, slack.MsgOptionText(joke.Text, false))
	return err
}

func main() {
	api := slack.New("xoxp-331495915735-329890989345-740906777732-27f6366524085866ac80355a3605cbb5")
	api1 := slack.New("xoxp-331495915735-329890989345-743581978470-f47e2dfba98fd156ee2d187672611cbc")

	channel, err := api.JoinChannel("general")
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 5)
		sendRandom(api, channel.Name, jokes)

		time.Sleep(time.Second * 5)
		sendRandom(api1, channel.Name, jokes)
	}
}
