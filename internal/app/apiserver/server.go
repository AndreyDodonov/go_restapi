/**
 * более легковесная часть сервера без лишних методов
 */

package apiserver

import (
	"encoding/json"
	"fmt"
	"go_restapi/internal/app/model"
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
	s.logger.Info("start API server")
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// wrapper for CORS start
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	// wrapper for CORS end
	s.router.ServeHTTP(w, r)
}

// конфигурируем роуты
func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/users", s.handleUsersGet()).Methods("GET")
	s.router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

}

func (s *server) handleUsersGet() http.HandlerFunc  {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		 allUser, err := s.store.User().Get()
		 if err != nil {
			s.logger.Warn("request all users error: ", err)
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		fmt.Println("all user: ", allUser) //TODO debug
	  s.respond(w, r, http.StatusOK, allUser)
		s.logger.Info("all users requested")
	}

}

// обработка "/users" Регистрация и аутентификация пользователей
func (s *server) handleUsersCreate() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		req := &request{}

		body, err := io.ReadAll(r.Body)
		fmt.Println("body is: ", body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("body error is: ", err) //TODO debug, need error handler
			return
		}

		err = json.Unmarshal(body, &req)
		fmt.Println("unmarshal email is: ", req.Email)       //TODO debug
		fmt.Println("unmarshal password is: ", req.Password) //TODO debug
		if err != nil {
			fmt.Println("unmarshal error is: ", err) //TODO debug, need error handler
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.logger.Warn("user create error: ", err)
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()

		// s.respond(w, r, http.StatusCreated, u)
		fmt.Println("user respond: ", w) //TODO debug
		s.logger.Info("new user was created")
	}
}

func (s *server) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Server is working \n Main router")
		http.FileServer(http.Dir("./web"))
	}
}

// хелпер для обработки ошибок
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// ответ сервера
func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}