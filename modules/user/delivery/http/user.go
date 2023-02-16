package http

import (
	"article_app/middleware"
	usecase2 "article_app/modules/jwt/usecase"
	"article_app/modules/user/delivery/http/controller"
	repository "article_app/modules/user/repository/postgres"
	"article_app/modules/user/usecase"
	pg "article_app/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	server                         = pg.ServerPG{}
	DB         *gorm.DB            = server.InitialPostgres()
	jwtUsecase usecase2.JWTUsecase = usecase2.NewJWTUsecase()

	userRepository repository.UserRepository = repository.NewUserRepository(DB)
	userUsecase    usecase.UserUsecase       = usecase.NewUserUsecase(userRepository)
	userController controller.UserController = controller.NewUserController(userUsecase)
)

func UserRouter(r *mux.Router) {
	apiRoute := r.PathPrefix("/api/v1/users").Subrouter()
	apiRoute.HandleFunc("/", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(userController.GetUsers))).Methods("GET")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(userController.GetUserByID))).Methods("GET")
	apiRoute.HandleFunc("/", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(userController.CreateUser))).Methods("POST")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(userController.UpdateUser))).Methods("PUT")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(userController.DeleteUser))).Methods("DELETE")
}
