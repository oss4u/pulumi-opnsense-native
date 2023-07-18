package opnsense

import (
	"github.com/go-resty/resty/v2"
)

func get_client() *resty.Request {
	client := resty.New()

	return client.R().
		SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")
}
