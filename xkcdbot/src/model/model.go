package model

import "fmt"
import "regexp"

type Rss struct {
  Title       string `xml:"channel>title"`
  Link        string `xml:"channel>link"`
  Description string `xml:"channel>description"`
  Items       []Item `xml:"channel>item"`
}

type Item struct {
  Title       string `xml:"title"`
  Link        string `xml:"link"`
  Description string `xml:"description"`
  PubDate     string `xml:"pubDate"`
  Guid        string `xml:"guid"`
}

func (i Item) String() string {
  return fmt.Sprintf("%-20s %-50s %s",
    i.Id(),
    i.ImgUrl(),
    i.PubDate)
}

func (i Item) ImgUrl() string {
  r, _ := regexp.Compile("http://imgs.xkcd.com/comics/[^\\.]+\\.[\\w]{3}")
  return r.FindString(i.Description)
}

func (i Item) Id() string {
  r, _ := regexp.Compile("http://xkcd.com/(\\d+)/")
  matches := r.FindStringSubmatch(i.Guid)
  return matches[1]
}

func (r Rss) String() string {
  return fmt.Sprintf("%s", r.Title)
}
