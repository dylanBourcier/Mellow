package servimpl

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"mellow/repositories/repoimpl"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	stmts := []string{
		`CREATE TABLE users (user_id VARCHAR PRIMARY KEY, username TEXT, image_url TEXT);`,
		`CREATE TABLE groups (group_id VARCHAR PRIMARY KEY, user_id VARCHAR NOT NULL, title TEXT, description TEXT, creation_date DATETIME NOT NULL);`,
		`CREATE TABLE groups_member (group_id VARCHAR, user_id VARCHAR, role TEXT, join_date DATETIME, PRIMARY KEY(group_id,user_id));`,
		`CREATE TABLE group_join_requests (id VARCHAR PRIMARY KEY, group_id VARCHAR NOT NULL, requester_id VARCHAR NOT NULL, status TEXT NOT NULL, created_at DATETIME NOT NULL, decided_at DATETIME, decided_by VARCHAR);`,
		`CREATE TABLE notifications (notification_id VARCHAR PRIMARY KEY, user_id VARCHAR NOT NULL, sender_id VARCHAR NOT NULL, request_id VARCHAR, type TEXT NOT NULL, seen BOOLEAN NOT NULL DEFAULT 0, creation_date DATETIME NOT NULL);`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			t.Fatalf("exec schema: %v", err)
		}
	}
	return db
}

func TestRequestJoin_HappyPath(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// seed data
	owner := uuid.New().String()
	user := uuid.New().String()
	group := uuid.New().String()
	_, _ = db.Exec(`INSERT INTO users(user_id,username) VALUES (?,?),(?,?)`, owner, "owner", user, "user")
	_, _ = db.Exec(`INSERT INTO groups(group_id,user_id,title,creation_date) VALUES (?,?,?,?)`, group, owner, "g1", time.Now())

	gjrRepo := repoimpl.NewGroupJoinRequestRepository(db)
	grpRepo := repoimpl.NewGroupRepository(db)
	notifRepo := repoimpl.NewNotificationRepository(db)
	userRepo := repoimpl.NewUserRepository(db)
	notifSvc := NewNotificationService(notifRepo, userRepo)

	svc := NewGroupJoinRequestService(gjrRepo, grpRepo, notifSvc)

	req, err := svc.RequestJoin(context.Background(), user, group)
	if err != nil {
		t.Fatalf("RequestJoin err: %v", err)
	}
	if req.Status != "pending" {
		t.Fatalf("expected pending, got %s", req.Status)
	}

	// duplicate should conflict
	if _, err := svc.RequestJoin(context.Background(), user, group); err == nil {
		t.Fatalf("expected conflict on duplicate pending")
	}
}
