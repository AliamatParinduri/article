package http

import (
	"article_app/middleware"
	usecase2 "article_app/modules/jwt/usecase"
	"article_app/modules/post/delivery/http/controller"
	repository "article_app/modules/post/repository/postgres"
	"article_app/modules/post/usecase"
	pg "article_app/repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	server                         = pg.ServerPG{}
	DB         *gorm.DB            = server.InitialPostgres()
	jwtUsecase usecase2.JWTUsecase = usecase2.NewJWTUsecase()

	postRepository repository.PostRepository = repository.NewPostRepository(DB)
	postUsecase    usecase.PostUsecase       = usecase.NewPostUsecase(postRepository)
	postController controller.PostController = controller.NewPostController(postUsecase)
)

func PostRouter(r *mux.Router) {
	apiRoute := r.PathPrefix("/api/v1/posts").Subrouter()
	apiRoute.HandleFunc("/", postController.GetPosts).Methods("GET")
	apiRoute.HandleFunc("/{id}/user", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(postController.GetPostByUser))).Methods("GET")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(postController.GetPostByID))).Methods("GET")
	apiRoute.HandleFunc("/", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(postController.CreatePost))).Methods("POST")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(postController.UpdatePost))).Methods("PUT")
	apiRoute.HandleFunc("/{id}", middleware.AuthorizeJWT(middleware.AuthorizeIsAdmin(postController.DeletePost))).Methods("DELETE")
}
