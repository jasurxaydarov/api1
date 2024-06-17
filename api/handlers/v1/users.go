package v1

import (
	"app/models"
	"encoding/json"
	"net/http"
)

var users = []models.User{

	{
		User_id: 1,
		User_name: "jasur",
	},
	{
	    User_id: 2,
	   	User_name: "botir",
   	},
	{
		User_id: 3,
		User_name: "aziz",
    },
}

func ( h *Handler)GetUsers(resp http.ResponseWriter,req *http.Request){

	json.NewEncoder(resp).Encode(users)

}