package xkcdbot

import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/xml"
import "time"

import "model"
import "utils"

func main() {
	resp, err := http.Get("http://xkcd.com/rss.xml")
	if err != nil {
		fmt.Printf("An error happened, shit!\n")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var r model.Rss
	e := xml.Unmarshal(body, &r)
	if e != nil {
		fmt.Printf("Unmarshall error: %s\n", e.Error)
	}

	now := time.Now()
	fmt.Println(utils.ReadReferenceId())

	for _, item := range r.Items {
		date := utils.ParseDate(item.PubDate)
		if now.After(date) {
			//fmt.Printf("viejuno")
		} else {
			//fmt.Printf("Esto no va a pasar")
		}

		fmt.Println(item)
	}
}
