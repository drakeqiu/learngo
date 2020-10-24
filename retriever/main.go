package main

import (
	"fmt"
	"github.com/drakeqiu/learngo/retriever/mock"
	re "github.com/drakeqiu/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,
		form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}
const url = "http://www.baidu.com"
func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":"james",
			"course":"golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) string {
	rp.Post(url, map[string]string {
		"contents": "another faked www.baidu.com",
	})
	return rp.Get(url)
	//rp.Post()
	//rp.Get()
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake www.baidu.com"}
	r = &retriever
	inspect(r)

	r = &re.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*re.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever.")
	}
	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *re.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
