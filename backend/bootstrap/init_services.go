package bootstrap

import (
	"mellow/services"
	"mellow/services/servimpl"
)

type Services struct {
	UserService             services.UserService
	AuthService             services.AuthService
	PostService             services.PostService
	GroupService            services.GroupService
	GroupJoinRequestService services.GroupJoinRequestService
	CommentService          services.CommentService
	MessageService          services.MessageService
	NotificationService     services.NotificationService
}

func InitServices(repos *Repositories) *Services {
	userService := servimpl.NewUserService(repos.UserRepository)
	authService := servimpl.NewAuthService(repos.UserRepository, repos.AuthRepository, userService)
	notificationService := servimpl.NewNotificationService(repos.NotificationRepository, repos.UserRepository)
	groupService := servimpl.NewGroupService(repos.GroupRepository, notificationService)
	postService := servimpl.NewPostService(repos.PostRepository, userService, groupService)
	commentService := servimpl.NewCommentService(repos.CommentRepository, repos.UserRepository, postService)
	messageService := servimpl.NewMessageService(repos.MessageRepository)
	groupJoinRequestService := servimpl.NewGroupJoinRequestService(repos.GroupJoinRequestRepository, repos.GroupRepository, notificationService)

	return &Services{
		UserService:             userService,
		AuthService:             authService,
		PostService:             postService,
		GroupService:            groupService,
		GroupJoinRequestService: groupJoinRequestService,
		CommentService:          commentService,
		NotificationService:     notificationService,
		MessageService:          messageService,
	}
}
