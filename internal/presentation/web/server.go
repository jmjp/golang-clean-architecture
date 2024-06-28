package web

import (
	"fmt"
	"net/http"
	"onion/internal/application/usecases"
	"onion/internal/infrastructure/database"
	"onion/internal/infrastructure/repositories"
	"onion/internal/presentation/web/handlers"
)

type Server struct {
	port   int
	router *http.ServeMux
}

// NewServer creates a new Server instance with the specified port.
//
// Parameters:
// - port: an integer representing the port number to listen on.
//
// Returns:
// - *Server: a pointer to the newly created Server instance.
func NewServer(port int) *Server {
	return &Server{
		port:   port,
		router: http.NewServeMux(),
	}
}

// Start starts the server and initializes the necessary repositories and use cases.
//
// It takes a pointer to a PostgresDB instance as a parameter and returns an error.
func (s *Server) Start(db *database.PostgresDB) error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.router,
	}

	userRepo := repositories.NewUserPostgresRepository(db)
	otpRepo := repositories.NewOtpPostgresRepository(db)

	useCase := usecases.NewLoginUsecase(userRepo, otpRepo)
	loginHandler := handlers.NewLoginHandler(useCase)

	s.router.HandleFunc("POST /auth", loginHandler.HandleFunc)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
