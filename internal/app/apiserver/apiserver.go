package apiserver

import (
	"go_restapi/internal/app/store/sqlstore"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *sqlstore.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// старт сервера
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	// if err := s.configureStore(); err != nil {
	// 	return err
	// }

	s.logger.Info("start API server on port 8080") //!TODO брать значение номера порта из конфига

	return http.ListenAndServe(s.config.BindAddress, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

// конфигурируем роуты
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/", s.handleMain())
}

// конфигурируем хранилище
// func (s *APIServer) configureStore() error  {
// 	st := sqlstore.New(s.config.Store)
// 	// if err := st.Open(); err != nil {
// 	// 	return err
// 	// }
// 	s.store = st
// 	return nil
// }

func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *APIServer) handleMain() http.HandlerFunc {
	type request struct {
		name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Main router")
	}
}
