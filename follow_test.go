/*
   Created by jinhan on 17-8-14.
   Tip:
   Update:
*/
package zhihuxx

import (
	"fmt"
	"testing"
)

func TestCatchUser(t *testing.T) {
	SetCookie("/home/jinhan/cookie.txt")
	a, e := CatchUser(true, "hunterhug", 1, 20)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		r, _ := Baba.JsonToString()
		fmt.Println(r)
		rr := ParseUser(a)
		for _, v := range rr.Data {
			fmt.Printf("%#v\n", v)
		}
	}
}
