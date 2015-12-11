package rss


import "testing"

var C Channel

func init() {
	C = NewChannel("foo", "/feeds.xml", "lorem ipsum dolor sit amet")
}

func TestAdd( t *testing.T){
	amount := len(C.Items)
	C.Add("Foo", "http://foo.bar/","Fubar")
	if amount == len(C.Items) {
		t.Fatal("Amount of items should have increased")
	}
}
