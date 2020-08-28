package users

import "encoding/json"

// PublicUser is the struct to process the user when returning to requests without a permission
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser is the struct to process the user when returning to requests with a permission
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall is a function used to process the user and return the user to marshalled to json (Public or Private)
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJSON, _ := json.Marshal(user)
	var private PrivateUser
	json.Unmarshal(userJSON, &private)
	return private
}

// Marshall is a function used to process the user and return the user to marshalled to json (Public or Private)
func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
