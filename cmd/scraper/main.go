package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/anthdm/hollywood/actor"
)

const scrapeInterval = time.Second

type scraper struct {
	url      string
	storePID *actor.PID
	engine   *actor.Engine
}

func newScraper(url string, storePID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &scraper{
			url:      url,
			storePID: storePID,
		}
	}
}

func (s *scraper) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.engine = c.Engine()
		go s.scrapeLoop()
	case actor.Stopped:
	default:
		_ = msg
	}
}

func (s *scraper) scrapeLoop() {
	for {
		resp, err := http.Get(s.url)
		if err != nil {
			panic(err)
		}
		var fact CatFact
		if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
			log.Println("failed to decode the response body")
			continue
		}
		time.Sleep(scrapeInterval)
	}
}

type CatFact struct {
	Fact string `jsopn:"fact"`
}

func main() {

}
