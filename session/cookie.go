package main

import (
	"net/http"
	"time"
)

func main() {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "sunlong", Expires: expiration}

}
