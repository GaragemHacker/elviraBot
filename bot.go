package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	urlT = "https://api.telegram.org/"
)

type Elvira struct {
	Token string `json:"token"`
}

func (e *Elvira) sendMsg(chatId int, msg string) {
	urlSite := urlT + e.Token + "/sendMessage"
	//fmt.Println("URL:>", urlSite)

	req, err := http.NewRequest("GET", urlSite, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content Type", "application/x-www-form-urlencoded;")
	req.Header.Set("charset:", "UTF-8")

	response, err := http.PostForm(urlSite, url.Values{"chat_id": {strconv.Itoa(chatId)}, "parse_mode": {"HTML"}, "text": {msg}})

	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()

		if err != nil {
			fmt.Printf("%s", err)
		}

	}
}

func (e *Elvira) sendPhoto(chatId int, img string) {
	urlSite := urlT + e.Token + "/sendPhoto"

	req, err := http.NewRequest("GET", urlSite, nil)

	req.Header.Set("Content Type", "application/x-www-form-urlencoded;")

	req.Header.Set("charset:", "UTF-8")

	response, err := http.PostForm(urlSite, url.Values{"chat_id": {strconv.Itoa(chatId)}, "photo": {"AgADAQADqqcxG9hEuAbvX1Tw8KBML0PH5y8ABPHpJLGgcfhStB4AAgI"}})

	if err != nil {
		return
	}

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("%s\n", string(contents))
	}
}
