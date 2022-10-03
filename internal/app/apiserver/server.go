/**
 * более легковесная часть сервера без лишних методов
 */

package apiserver

import (
	"fmt"
	"go_restapi/internal/app/store"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

// новый сервер
func newServer(store store.Store) *server {
	fmt.Println("newServer") //TODO debug
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// конфигурируем роуты
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/", s.handleMain())
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	 }
}

func (s *server) handleMain() http.HandlerFunc  {
	type request struct {
		name string
	}

	return func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "Server is working \n Main router")
	}
}