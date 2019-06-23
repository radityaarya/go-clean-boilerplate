package utils

import "net/http"

type Response struct {
	Meta meta `json:"meta"`
	Payload payload `json:"payload"`
}

type meta struct {
	Status      string `json:"status"`
	ErrorMessage string `json:"error"`
}

type payload struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{}  `json:"pagination,omitempty"`
}

type Pagination struct {
	Page      int    `json:"page,omitempty"`
	Count     int    `json:"count,omitempty"`
	Total     int    `json:"total,omitempty"`
	TotalPage int    `json:"total_page,omitempty"`
}

var (
	OK 		 			= http.StatusText(http.StatusOK)
	NotFound 			= http.StatusText(http.StatusNotFound)
	InternalServerError = http.StatusText(http.StatusInternalServerError)
)

// ResponseSuccess used for handling response with http code 200 (OK / Success)
func ResponseSuccess() *Response {
	return &Response{
		Meta: meta{
			Status: OK,
		},
	}
}

// ResponseNotFound used for handling response with http code 404 (Not found)
func ResponseNotFound() *Response {
	return &Response{
		Meta: meta{
			Status: NotFound,
		},
	}
}

// ResponseInternalServerError used for handling response with http code 500 (Internal server error)
func ResponseInternalServerError() *Response {
	return &Response{
		Meta: meta{
			Status: InternalServerError,
		},
	}
}