package worker

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/model"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func Start() {
	urls := model.ListUrls()

	for _, url := range urls {
		log.Print("Working on url id: " + url.Id)
		go Run(url.Id)
	}
}

func Run(id string) {
	var content []byte
	for {
		url := model.GetUrl(id)

		if url.Url == "" {
			log.Print("Empty url.")
			return
		}

		startTime := time.Now()

		resp, err := client.Get(url.Url)

		if err != nil {
			log.Print(err)
			content = nil
		} else {
			content, _ = ioutil.ReadAll(resp.Body)
		}

		duration := time.Now().Sub(startTime).Seconds()
		model.SaveToHistory(id, string(content), duration)

		time.Sleep(time.Duration(url.Interval) * time.Second)
	}
}
