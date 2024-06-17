package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{

}

func NewHandler()Handler{

	return Handler{}
}

func (h *Handler)Pong( resp http.ResponseWriter, req *http.Request){

	fmt.Println("req method", req.Method)

	response:=map[string]string{"massage":"pong"}

	encoder:=json.NewEncoder(resp)

	encoder.Encode(response)


}