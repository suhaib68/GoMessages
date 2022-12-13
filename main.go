// GoMessages project main.go
package main

import (
	"net/http"
)

func main() {
	print("http://localhost:10085")
	http.ListenAndServe(":10085", nil)
}
