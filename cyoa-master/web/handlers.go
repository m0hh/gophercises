package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type Whole struct {
	Intro     Compononet `json:"intro"`
	NewYork   Compononet `json:"new-york"`
	Debate    Compononet `json:"debate"`
	SeanKelly Compononet `json:"sean-kelly"`
	MarkBates Compononet `json:"mark-bates"`
	Denver    Compononet `json:"denver"`
	Home      Compononet `json:"home"`
}

type Compononet struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func (app *application) handleMain(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("story")

	temp, err := template.ParseFiles("templates/home.html")

	if err != nil {
		app.errorLogger.Println(err)
		return
	}

	if query == "" {
		temp.Execute(w, app.mapp["intro"])
	} else {
		temp.Execute(w, app.mapp[query])
	}
}

func (app *application) buildMap() map[string]Compononet {

	file, err := os.ReadFile("gopher.json")
	if err != nil {
		app.errorLogger.Println(err)
	}

	var story Whole

	err = json.Unmarshal(file, &story)

	if err != nil {
		app.errorLogger.Println(err)
	}

	mapp := make(map[string]Compononet)
	mapp["intro"] = story.Intro
	mapp["new-york"] = story.NewYork
	mapp["debate"] = story.Debate
	mapp["sean-kelly"] = story.SeanKelly
	mapp["mark-bates"] = story.MarkBates
	mapp["denver"] = story.Denver
	mapp["home"] = story.Home

	return mapp
}
