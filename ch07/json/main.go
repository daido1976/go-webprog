package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("./ch07/json/post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	fmt.Println(string(jsonData))
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
	fmt.Println(post.Content)
	fmt.Println(post.Author.ID)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments[0].ID)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
}
