package story

import (
	"encoding/json"
	"io/ioutil"
)

type Story map[string]StoryPart

func (s *Story) ReadStoryFromFile(filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, s)
	return err
}

type StoryPart struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
