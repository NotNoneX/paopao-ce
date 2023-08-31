// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: tweets.sql

package dbr

import (
	"context"
)

const getPostById = `-- name: GetPostById :one

SELECT id, user_id, comment_count, collection_count, upvote_count, is_top, is_essence, is_lock, latest_replied_on, tags, attachment_price, ip, ip_loc, created_on, modified_on, deleted_on, is_del, visibility, share_count FROM p_post WHERE id=$1 AND is_del=0
`

// ------------------------------------------------------------------------------
// tweet sql dml
// ------------------------------------------------------------------------------
func (q *Queries) GetPostById(ctx context.Context, id int64) (*PPost, error) {
	row := q.db.QueryRow(ctx, getPostById, id)
	var i PPost
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CommentCount,
		&i.CollectionCount,
		&i.UpvoteCount,
		&i.IsTop,
		&i.IsEssence,
		&i.IsLock,
		&i.LatestRepliedOn,
		&i.Tags,
		&i.AttachmentPrice,
		&i.Ip,
		&i.IpLoc,
		&i.CreatedOn,
		&i.ModifiedOn,
		&i.DeletedOn,
		&i.IsDel,
		&i.Visibility,
		&i.ShareCount,
	)
	return &i, err
}
