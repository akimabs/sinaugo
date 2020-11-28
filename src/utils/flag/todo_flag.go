package flag

// TodoFlag is response code and message for Todo controller
type TodoFlag struct {
	Message string
	Flag    string
	Error
}

var (
	/********************************
	 * GET MANY TODO ENDPOINT
	 ********************************/

	// GetAllTodo is response code and message
	// for get Todos with success result
	GetAllTodo TodoFlag = TodoFlag{
		Message: "Success get all todos",
	}

	/********************************
	 * GET TODO ENDPOINT
	 ********************************/

	// GetTodoSuccess is response code and message
	// for get Todo with success result
	GetTodoSuccess TodoFlag = TodoFlag{
		Message: "Success get todo",
	}

	// GetTodoNotFound is response code and message
	// For get Todo with Todo not found
	// in database
	GetTodoNotFound TodoFlag = TodoFlag{
		Message: "Todo not found",
		Error: Error{
			Message: "Todo with this ID not enough",
			Flag:    "Todo_NOT_FOUND",
		},
	}

	// GetTodoInvalidParamID is response code and message
	// For get Todo with invalid param uri
	GetTodoInvalidParamID TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}

	/********************************
	 * CREATE TODO ENDPOINT
	 ********************************/

	// CreateTodoSuccess is response code and message
	// for create Todo with success result
	CreateTodoSuccess TodoFlag = TodoFlag{
		Message: "Success create Todo",
	}

	// CreateTodoAlreadyExist is response code and message
	// for create Todo with Todo already
	// exist in database
	CreateTodoAlreadyExist TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Todo with this email already exist",
			Flag:    "Todo_ALREADY_EXIST",
		},
	}

	// CreateTodoInvalidBody is response code and message
	// for create Todo with invalid body
	CreateTodoInvalidBody TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Flag: "INVALID_BODY",
		},
	}

	/********************************
	 * UPDATE TODO ENDPOINT
	 ********************************/

	// UpdateTodoSuccess is response code and message
	// for update Todo with success result
	UpdateTodoSuccess TodoFlag = TodoFlag{
		Message: "Success update Todo",
	}

	// UpdateTodoNotExist is response code and message
	// for update Todo with Todo not
	// exist in database
	UpdateTodoNotExist TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Todo with this id not exist",
			Flag:    "Todo_NOT_FOUND",
		},
	}

	// UpdateTodoInvalidBody is response code and message
	// for update Todo with Todo not
	// exist in database
	UpdateTodoInvalidBody TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Flag: "INVALID_BODY",
		},
	}

	// UpdateTodoInvlidParamURI is response code and message
	// for update Todo with invalid param uri
	UpdateTodoInvlidParamURI TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}

	/********************************
	 * DELETE TODO ENDPOINT
	 ********************************/

	// DeleteTodoSuccess is response code and message
	// for delete Todo with success result
	DeleteTodoSuccess TodoFlag = TodoFlag{
		Message: "Success delete Todo",
	}

	// DeleteTodoNotExist is response code and message
	// for delete Todo with Todo not
	// exist in database
	DeleteTodoNotExist TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Todo with this id not exist",
			Flag:    "Todo_NOT_FOUND",
		},
	}

	// DeleteTodoInvalidParamURI is response code and message
	// for delete Todo with invalid param uri
	DeleteTodoInvalidParamURI TodoFlag = TodoFlag{
		Flag: "BAD_REQUEST",
		Error: Error{
			Message: "Param must be of type integer, required",
			Flag:    "INVALID_BODY",
		},
	}
)