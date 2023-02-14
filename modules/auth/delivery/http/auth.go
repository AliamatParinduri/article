package http

import (
	"article_app/modules/Auth/usecase"
	"article_app/modules/auth/delivery/http/controller"
	repository "article_app/modules/auth/repository/postgres"
	pg "article_app/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB           = pg.InitialPostgres()
	jwtUsecase usecase.JWTUsecase = usecase.NewJWTUsecase()

	authRepository repository.AuthRepository = repository.NewAuthRepository(DB)
	authUsecase    usecase.AuthUsecase       = usecase.NewAuthUsecase(authRepository)
	authController controller.AuthController = controller.NewAuthController(authUsecase, jwtUsecase)
)

func AuthRouter(r *mux.Router) {
	apiRoute := r.PathPrefix("/api/v1/auth").Subrouter()
	apiRoute.HandleFunc("/register", authController.Register).Methods("POST")
	apiRoute.HandleFunc("/login", authController.Login).Methods("POST")
}