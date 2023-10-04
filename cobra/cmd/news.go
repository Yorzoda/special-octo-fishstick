/*
Copyright © 2023 Yorzoda Behzod <yorzoda.behzod@gmail.com>
This file is part of useless CLI application.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

type (
	Response struct {
		Status       string     `json:"status"`
		TotalResults int        `json:"totalResults"`
		Articles     []Articles `json:"articles"`
	}

	Articles struct {
		Source struct {
			ID   any    `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	}
)

var (
	newsTopic string
	resp      Response
)

// newsCmd represents the News command
var newsCmd = &cobra.Command{
	Use:   "getNews",
	Short: "Wanna know something new?",
	Long:  `news command gives you a news by your topic or by default return actual non-topic news🥸`,
	Run: func(cmd *cobra.Command, args []string) {
		getNews(newsTopic)
	},
}

func init() {

	rootCmd.AddCommand(newsCmd)
	newsCmd.Flags().StringVarP(&newsTopic, "newsTopic", "n", "actual", "for getting news topic that you need")
	if err := newsCmd.MarkFlagRequired("newsTopic"); err != nil {
		fmt.Println(err)
	}
}

func getNews(topic string) {
	url := `https://newsapi.org/v2/everything?apiKey=8f4f98352f1b4a40baac9c7fa1db5e74&q=` + topic
	response, err := http.Get(url)
	if err != nil {
		log.Print("http.Get err:", err)
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&resp)

	if resp.TotalResults == 0 {
		fmt.Println("Oops there's no news for this topic🙁,maybe you should try for another one")
	}

	for _, i := range resp.Articles {
		fmt.Printf("\"Source\":\"%v\"©️\n", i.Source.Name)
		fmt.Printf("\"Title\"\":\"%v\" 🆕\n", i.Title)
		fmt.Printf("\"Description\":\"%v\"📰\n", i.Description)
		fmt.Println("")
	}

}
