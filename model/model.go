package model

type (
	Ticket_options struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Desc        string `json:"desc"`
		Allocation  int    `json:"allocation"`
		Create_Date string `json:"Create_Date"`
	}

	Ticket_purchases struct {
		Id       int    `json:"id"`
		Quantity int    `json:"quantity"`
		User_id  string `json:"user_id"`
	}
	Ticket_user struct {
		Id      int    `json:"id"`
		User_id string `json:"user_id"`
		Name    string `json:"name"`
	}


)
