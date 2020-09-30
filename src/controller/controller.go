package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/model"
	"github.com/mpludowski/bWljaGFsLXBsdWRvd3NraQo/worker"
)

type idResponse struct {
	Id string `json:"id"`
}

func PostFetcher(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Url      string `json:"url"`
		Interval int    `json:"interval"`
	}
	err := parseRequest(r, &request)

	if err != nil {
		if err.Error() == "http: request body too large" {
			w.WriteHeader(413)
		} else {
			w.WriteHeader(400)
		}
		return
	}

	var result idResponse
	result.Id = model.SaveUrl(request.Url, request.Interval)

	go worker.Run(result.Id)

	response, _ := json.Marshal(result)

	w.Write(response)
}

func DeleteFetcher(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	model.DeleteUrl(id)

	response, _ := json.Marshal(idResponse{Id: id})
	w.Write(response)
}

func ListFetchers(w http.ResponseWriter, r *http.Request) {
	urls := model.ListUrls()

	response, _ := json.Marshal(urls)

	w.Write(response)
}

func History(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	history := model.GetHistory(id)

	response, _ := json.Marshal(history)

	w.Write(response)
}

func parseRequest(r *http.Request, i interface{}) error {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	return json.Unmarshal(b, i)

}
