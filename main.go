package main

import (
	"fmt"
	"time"
)


func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i:=0;i<len(people); i++ {
		fmt.Println(<-c)
	}


}

func sexyCount(person string){
	for i:=0;i<10;i++{
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Request failed")

// func main() {
// 	var results = make(map[string]string)
// 	urls := []string {
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"https://www.academy.nomadcoders.co/",
// 		"https://www.raddit.com/",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	// fmt.Println(results)
// 	for url, result := range(results) {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking:", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println("resp:", resp)
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}
// 	return nil	
// }