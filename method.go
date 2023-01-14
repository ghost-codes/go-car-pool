package main

type Method struct {
	GET    string
	POST   string
	PUT    string
	DELETE string
	PATCH  string
}

var Methods = Method{
	GET:    "GET",
	POST:   "POST",
	PUT:    "PUT",
	DELETE: "DELETE",
	PATCH:  "PATCH",
}
