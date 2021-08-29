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

func decodeBodyPost(msg *dbprovider.PostMessageUser, w http.ResponseWriter, r *http.Request) (*dbprovider.PostMessageUser, error) {
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
			argToFindToUint, err := strconv.ParseUint(argToFind, 10, 64)
			if err != nil {
				log.Print("Argument in invalid!")
			}
			js := dbprovider.Mgr.GetBalance(argToFindToUint)
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)

		case http.MethodPost:
			var msg dbprovider.PostMessageUser
			var ms dbprovider.MessageUser
			// decodeBody(&ms, w, r)
			decodeBodyPost(&msg, w, r)

			if msg.Action == "Add" {

				js := dbprovider.Mgr.AddBalance(uint64(msg.Id), uint64(msg.Sum))
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)

			} else if msg.Action == "Substract" {

				js := dbprovider.Mgr.SubBalance(uint64(msg.Id), uint64(msg.Sum))
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)

			} else if msg.Action == "Send" {

			} else {
				dbprovider.Mgr.AddUser(&ms)
				output, err := json.Marshal(ms)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				w.Header().Set("content-type", "application/json")
				w.Write(output)
			}

		case http.MethodPut:
			fmt.Println("PUT")
		case http.MethodDelete:
			fmt.Println("DELETE")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
