package bootstrap

import (
	"database/sql"
	"mellow/repositories"
	"mellow/repositories/repoimpl"
)

type Repositories struct {
	UserRepository             repositories.UserRepository
	AuthRepository             repositories.AuthRepository
	PostRepository             repositories.PostRepository
	GroupRepository            repositories.GroupRepository
	GroupJoinRequestRepository repositories.GroupJoinRequestRepository
	CommentRepository          repositories.CommentRepository
	MessageRepository          repositories.MessageRepository
	NotificationRepository     repositories.NotificationRepository
}

func InitRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository:             repoimpl.NewUserRepository(db),
		AuthRepository:             repoimpl.NewAuthRepository(db),
		PostRepository:             repoimpl.NewPostRepository(db),
		GroupRepository:            repoimpl.NewGroupRepository(db),
		GroupJoinRequestRepository: repoimpl.NewGroupJoinRequestRepository(db),
		CommentRepository:          repoimpl.NewCommentRepository(db),
		MessageRepository:          repoimpl.NewMessageRepository(db),
		NotificationRepository:     repoimpl.NewNotificationRepository(db),
	}
}
