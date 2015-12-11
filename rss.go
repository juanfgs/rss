package rss

import "encoding/xml"

type Channel struct {
	XMLName xml.Name `xml:"channel"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"items"`

}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

func NewChannel(title, link, description string) Channel {
	return Channel{Title:title,Link:link,Description:description}
}

func (self *Channel) Add(title, link, description string) {
	self.Items = append(self.Items, Item{Title:title,Link:link,Description:description})

}

func (self Channel) Marshal() []byte{
	output,err := xml.Marshal(self)
	if err != nil {
		panic(err)
	}
	return output
}
