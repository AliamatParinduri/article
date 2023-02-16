package http

import (
	"article_app/middleware"
	usecase2 "article_app/modules/jwt/usecase"
	"article_app/modules/tag/delivery/http/controller"
	repository "article_app/modules/tag/repository/postgres"
	"article_app/modules/tag/usecase"
	pg "article_app/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	server                         = pg.ServerPG{}
	DB         *gorm.DB            = server.InitialPostgres()
	jwtUsecase usecase2.JWTUsecase = usecase2.NewJWTUsecase()

	tagRepository repository.TagRepository = repository.NewTagRepository(DB)
	tagUsecase    usecase.TagUsecase       = usecase.NewTagUsecase(tagRepository)
	tagController controller.TagController = controller.NewTagController(tagUsecase)
)

func TagRouter(r *mux.Router) {
	apiRoute := r.PathPrefix("/api/v1/tags").Subrouter()
	apiRoute.HandleFunc("/", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(tagController.GetTags))).Methods("GET")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(tagController.GetTagByID))).Methods("GET")
	apiRoute.HandleFunc("/", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(tagController.CreateTag))).Methods("POST")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(tagController.UpdateTag))).Methods("PUT")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(tagController.DeleteTag))).Methods("DELETE")
}
