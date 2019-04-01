package main

import (
	"fmt"
	"golang-learn/base/retiever/mock"
	"golang-learn/base/retiever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string  {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is fake impl"}
	fmt.Printf("%T %v\n",r,r)
	r=real.Retriever{}

	r = &real.Retriever{
		UserAgent:"Kuuyee2.0",
		TimeOut:time.Minute,
	}
	//fmt.Println(download(r))
	fmt.Printf("%T %v\n",r,r)
	inspect(r)

	realRetriever :=r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)



}
func inspect(r Retriever)  {
	fmt.Printf("%T %v\n",r,r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:",v.UserAgent)
	}
}