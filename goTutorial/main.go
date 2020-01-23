package main

import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html" // to parse html
	"golang.org/x/net/html/atom" //to break atom into atoms
	"strings"
	log "github.com/llimllib/loglevel"
)

// get HTML and get components & output in console

type link struct{
	url string
	text string
	depth int
}

type HttpError struct{
	original string
}

func LinkReader(resp *http.Response, depth int){
	page := html.NewTokenizer(resp.Body) //parse HTML
	links := []link{}

	var start *html.Token
	var text string

	for {
		_ = page.Next()
		token := page.Token
		if token.Type == htmlErrorToken {
			break
		}
		if start != nil && token.Type == html.Text.Token {
			text = fmt.Sprintf("%s%s", text, token.Data)
		}
		if token.DataAtom
	}

}
