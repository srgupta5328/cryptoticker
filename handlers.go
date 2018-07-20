package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	query        string
	start        int
	limit        int
	convert      string
	queryBuilder string
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Newsfeed Gopher!</h1>")
}

func (conf *Config) GlobalHandler(w http.ResponseWriter, r *http.Request) {
	//localhost:8080/api/global/
	url := conf.URL + conf.API_VERSION + "global/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Making Request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Client couldn't complete the request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(body)

}

func (conf *Config) TickerHandler(w http.ResponseWriter, r *http.Request) {
	//localhost:8080/api/ticker/
	url := conf.URL + conf.API_VERSION + "ticker/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Making Request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error Client couldn't make a request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(body)

}

func (conf *Config) TargetTickerHandler(w http.ResponseWriter, r *http.Request) {
	//localhost:8080//api/ticker/{id}
	uri := conf.URL + conf.API_VERSION + "/ticker/"
	vars := mux.Vars(r)
	Id := vars["id"]

	url := uri + Id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Making Request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error Client couldn't make a request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(body)

}

func (conf *Config) ListingHandler(w http.ResponseWriter, r *http.Request) {
	//localhost:8080//api/listings/
	url := conf.URL + conf.API_VERSION + "listings/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Making the Request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Client couldn't execute request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.Write(body)
}
