package server

type User struct {
	Guid      string `json:"guid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Admin     bool   `json:"admin"`
}

type UserCreationResponse struct {
	UserGuid string `json:"user_guid"`
	State    string `json:"state"`
}
