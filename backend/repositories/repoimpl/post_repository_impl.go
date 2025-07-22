package repoimpl

import (
	"context"
	"database/sql"
	"fmt"
	"mellow/models"
	"mellow/repositories"
)

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repositories.PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) InsertPost(ctx context.Context, post *models.Post) error {
	query := `INSERT INTO posts (post_id, group_id, user_id, title, content, creation_date, visibility, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, post.PostID, post.GroupID, post.UserID, post.Title, post.Content, post.CreationDate, post.Visibility, post.ImageURL)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}
	return nil
}
func (r *postRepositoryImpl) GetPostByID(ctx context.Context, postID string) (*models.PostDetails, error) {
	query := `
		SELECT 
			p.post_id, p.group_id, p.user_id, p.title, p.content, 
			p.creation_date, p.visibility, p.image_url, 
			u.username, u.image_url, 
			(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.post_id) AS comment_count
		FROM 
			posts p
		JOIN 
			users u ON p.user_id = u.user_id
		WHERE 
			p.post_id = ?`
	var post models.PostDetails
	err := r.db.QueryRowContext(ctx, query, postID).Scan(
		&post.PostID, &post.GroupID, &post.UserID, &post.Title,
		&post.Content, &post.CreationDate, &post.Visibility, &post.ImageURL,
		&post.Username, &post.AvatarURL, &post.CommentsCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Post not found
		}
		return nil, fmt.Errorf("error retrieving post: %w", err)
	}
	return &post, nil
}

func (r *postRepositoryImpl) UpdatePost(ctx context.Context, post *models.Post) error {
	query := `UPDATE posts SET title = ?, content = ? WHERE post_id = ?`
	_, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.PostID)
	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (r *postRepositoryImpl) DeletePost(ctx context.Context, postID string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM posts WHERE post_id = ?`, postID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}
	return nil
}

func (r *postRepositoryImpl) GetFeed(ctx context.Context, userID *string, limit, offset int) ([]*models.PostDetails, error) {
	query := `
		WITH user_follows AS (
			SELECT followed_id
			FROM follows
			WHERE follower_id = ?
		),
		authorized_private AS (
			SELECT post_id
			FROM posts_viewer
			WHERE user_id = ?
		)
		SELECT 
			p.post_id,p.title, p.content, p.creation_date, p.visibility, p.user_id,
			u.username, u.image_url AS avatar_url,
			g.group_id AS group_id, g.title AS group_title,
			(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.post_id) AS comment_count
		FROM posts p
		JOIN users u ON p.user_id = u.user_id
		LEFT JOIN groups g ON p.group_id = g.group_id
		WHERE
			-- Si post de groupe, visible uniquement par les membres
			(
				p.group_id IS NOT NULL AND EXISTS (
					SELECT 1 FROM groups_member gm
					WHERE gm.group_id = p.group_id AND gm.user_id = ?
				)
			)
			OR
			(
				p.group_id IS NULL AND (
					p.visibility = 'public'
					OR (p.visibility = 'followers' AND (p.user_id IN (SELECT followed_id FROM user_follows) OR p.user_id = ?))
					OR (p.visibility = 'private' AND (p.post_id IN (SELECT post_id FROM authorized_private) OR p.user_id = ?))
				)
			)
		ORDER BY p.creation_date DESC
		LIMIT ? OFFSET ?;
	`

	args := []interface{}{userID, userID, userID, userID, userID, limit, offset}
	if userID == nil {
		// Si l'utilisateur n'est pas connecté, ignorer followers/private
		query = `
			SELECT 
				p.post_id, p.title, p.content, p.creation_date, p.visibility, p.user_id,
				u.username, u.image_url AS avatar_url,
				g.group_id AS group_id, g.title AS group_title,
				(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.post_id) AS comment_count
			FROM posts p
			JOIN users u ON p.user_id = u.user_id
			LEFT JOIN groups g ON p.group_id = g.group_id
			WHERE p.group_id IS NULL AND p.visibility = 'public'
			ORDER BY p.creation_date DESC
			LIMIT ? OFFSET ?;
		`
		args = []interface{}{limit, offset}

	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.PostDetails
	for rows.Next() {
		var p models.PostDetails
		if err := rows.Scan(&p.PostID, &p.Title, &p.Content, &p.CreationDate, &p.Visibility, &p.UserID,
			&p.Username, &p.AvatarURL, &p.GroupID, &p.GroupName, &p.CommentsCount); err != nil {
			return nil, err
		}

		posts = append(posts, &p)
	}

	return posts, nil
}

func (r *postRepositoryImpl) GetUserPosts(ctx context.Context, ownerID string) ([]*models.Post, error) {
	// TODO: SELECT * FROM posts WHERE author_id = ?
	return nil, nil
}

func (r *postRepositoryImpl) InsertPostReport(ctx context.Context, report *models.Report) error {
	// TODO: INSERT INTO reports (id, reporter_id, post_id, reason, created_at) VALUES (?, ?, ?, ?, ?)
	return nil
}

func (r *postRepositoryImpl) IsPostExisting(ctx context.Context, postID string) (bool, error) {
	query := `SELECT COUNT(*) FROM posts WHERE post_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, postID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking post existence: %w", err)
	}
	return count > 0, nil
}

func (r *postRepositoryImpl) IsUserAllowed(ctx context.Context, postID string, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM posts_viewer WHERE post_id = ? AND user_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, postID, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking user access to post: %w", err)
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}
