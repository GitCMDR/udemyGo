package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Go methods can return two variables and you always need to use ALL the variables. The blank identifier is
	// used to be able to acknowledge both variables and throw away the one that you don't want to use in this
	// case, being error or 'err'.
	res, _ := http.Get("http://t.uber.com/sticker-request")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
