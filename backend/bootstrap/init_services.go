package bootstrap

import (
	"mellow/services"
	"mellow/services/servimpl"
)

type Services struct {
	UserService    services.UserService
	AuthService    services.AuthService
	PostService    services.PostService
	GroupService   services.GroupService
	CommentService services.CommentService
}

func InitServices(repos *Repositories) *Services {
	userService := servimpl.NewUserService(repos.UserRepository)
	authService := servimpl.NewAuthService(repos.UserRepository, repos.AuthRepository, userService)
	postService := servimpl.NewPostService(repos.PostRepository)
	groupService := servimpl.NewGroupService(repos.GroupRepository)
	commentService := servimpl.NewCommentService(repos.CommentRepository)

	return &Services{
		UserService:    userService,
		AuthService:    authService,
		PostService:    postService,
		GroupService:   groupService,
		CommentService: commentService,
	}
}
