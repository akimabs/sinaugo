package response

// Success -> general response success format
type Success struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Response -> general response format
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []Error     `json:"errors"`
}


// Error -> general response error format
type Error struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}