package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"

	dbprovider "github.com/deanonqq/microservice-user-balance/internal/dbprovider"

	"io/ioutil"
)

type ProductView struct {
	Code  string
	Price uint
}

func decodeBody(msg *dbprovider.MessageUser, w http.ResponseWriter, r *http.Request) (*dbprovider.MessageUser, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	err = json.Unmarshal(b, msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	return msg, err
}

func (s *APIServer) handleUserBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			arg := r.URL.Query()
			argToFind := arg.Get("Id")
			argToFind_to_uint, err := strconv.ParseUint(argToFind, 10, 64)
			if err != nil {
				log.Print("Argument in invalid!")
			}
			js := dbprovider.Mgr.GetBalance(argToFind_to_uint)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		case http.MethodPost:
			var msg dbprovider.MessageUser
			decodeBody(&msg, w, r)

			if &msg.Action == "add" {

			} else if &msg.Action == "substract" {

			}

			dbprovider.Mgr.AddUser(&msg)

			output, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.Header().Set("content-type", "application/json")
			w.Write(output)
		case http.MethodPut:
			fmt.Println("PUT")
		case http.MethodDelete:
			fmt.Println("DELETE")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
