// Code generated by Yesql. DO NOT EDIT.
// versions:
// - Yesql v1.9.0

package ac

import (
	"context"
	"fmt"

	"github.com/alimy/yesql"
	"github.com/bitbus/sqlx"
)

var (
	_ = fmt.Errorf("error for placeholder")
)

const (
	_ShipIndexA_UserInfo        = `SELECT * FROM @user WHERE username=?`
	_SimpleIndexA_UserInfo      = `SELECT * FROM @user WHERE username=?`
	_TopicA_DecrTagsById        = `UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?)`
	_TopicA_ExistTopicUser      = `SELECT 1 FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0`
	_TopicA_FollowPinTags       = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id "u.id", 1 as is_following, c.is_top, c.is_pin, u.nickname "u.nickname", u.username "u.username", u.status "u.status", u.avatar "u.avatar", u.is_admin "u.is_admin" FROM @topic_user c JOIN @user u ON c.user_id = u.id JOIN @tag t ON c.topic_id = t.id WHERE c.is_del = 0 AND t.quote_num > 0 AND c.user_id=? AND c.is_pin=1 ORDER BY c.is_top DESC, t.quote_num DESC LIMIT ? OFFSET ?`
	_TopicA_FollowTags          = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id "u.id", 1 as is_following, c.is_top, c.is_pin, u.nickname "u.nickname", u.username "u.username", u.status "u.status", u.avatar "u.avatar", u.is_admin "u.is_admin" FROM @topic_user c JOIN @user u ON c.user_id = u.id JOIN @tag t ON c.topic_id = t.id WHERE c.is_del = 0 AND t.quote_num > 0 AND c.user_id=? ORDER BY c.is_top DESC, t.quote_num DESC LIMIT ? OFFSET ?`
	_TopicA_FollowTopic         = `INSERT INTO @topic_user(user_id, topic_id, created_on) VALUES (?, ?, ?)`
	_TopicA_HotTags             = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id "u.id", u.nickname "u.nickname", u.username "u.username", u.status "u.status", u.avatar "u.avatar", u.is_admin "u.is_admin" FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.quote_num DESC LIMIT ? OFFSET ?`
	_TopicA_IncrTagsById        = `UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?)`
	_TopicA_InsertTag           = `INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1)`
	_TopicA_NewestTags          = `SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id "u.id", u.nickname "u.nickname", u.username "u.username", u.status "u.status", u.avatar "u.avatar", u.is_admin "u.is_admin" FROM @tag t JOIN @user u ON t.user_id = u.id WHERE t.is_del = 0 AND t.quote_num > 0 ORDER BY t.id DESC LIMIT ? OFFSET ?`
	_TopicA_PinTopic            = `UPDATE @topic_user SET is_pin=1-is_pin, modified_on=? WHERE user_id=? AND topic_id=? AND is_del=0`
	_TopicA_StickTopic          = `UPDATE @topic_user SET is_top=1-is_top, modified_on=? WHERE user_id=? AND topic_id=? AND is_del=0`
	_TopicA_TagsByIdA           = `SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0`
	_TopicA_TagsByIdB           = `SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?)`
	_TopicA_TagsByKeywordA      = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6`
	_TopicA_TagsByKeywordB      = `SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6`
	_TopicA_TagsForIncr         = `SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?)`
	_TopicA_TopicInfos          = `SELECT topic_id, is_top FROM @topic_user WHERE is_del=0 AND user_id=? AND topic_id IN (?)`
	_TopicA_TopicIsPin          = `SELECT is_pin FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0`
	_TopicA_TopicIsTop          = `SELECT is_top FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0`
	_TopicA_UnfollowTopic       = `DELETE FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0`
	_TweetA_AttachmentByTweetId = `SELECT * FROM @user WHERE username=?`
	_TweetA_FavoriteByTweetId   = `SELECT * FROM @user WHERE username=?`
	_TweetA_ReactionByTweetId   = `SELECT * FROM @user WHERE username=?`
	_TweetA_TweetInfoById       = `SELECT * FROM @user WHERE username=?`
	_TweetA_TweetItemById       = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserFavorites       = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserInfo            = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserReactions       = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserTweetsByAdmin   = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserTweetsByFriend  = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserTweetsByGuest   = `SELECT * FROM @user WHERE username=?`
	_TweetA_UserTweetsBySelf    = `SELECT * FROM @user WHERE username=?`
	_TweetHelpA_UserInfo        = `SELECT * FROM @user WHERE username=?`
	_TweetManageA_UserInfo      = `SELECT * FROM @user WHERE username=?`
)

// PreparexContext enhances the Conn interface with context.
type PreparexContext interface {
	// PrepareContext prepares a statement.
	// The provided context is used for the preparation of the statement, not for
	// the execution of the statement.
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)

	// PrepareNamedContext returns an sqlx.NamedStmt
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind rebind query to adapte SQL Driver
	Rebind(query string) string
}

// PreparexBuilder preparex builder interface for sqlx
type PreparexBuilder interface {
	PreparexContext
	QueryHook(query string) string
}

type ShipIndexA struct {
	yesql.Namespace `yesql:"ship_index_a"`
	UserInfo        string `yesql:"user_info"`
}

type SimpleIndexA struct {
	yesql.Namespace `yesql:"simple_index_a"`
	UserInfo        string `yesql:"user_info"`
}

type TopicA struct {
	yesql.Namespace `yesql:"topic_a"`
	DecrTagsById    string     `yesql:"decr_tags_by_id"`
	IncrTagsById    string     `yesql:"incr_tags_by_id"`
	TagsByIdA       string     `yesql:"tags_by_id_a"`
	TagsByIdB       string     `yesql:"tags_by_id_b"`
	TagsForIncr     string     `yesql:"tags_for_incr"`
	TopicInfos      string     `yesql:"topic_infos"`
	ExistTopicUser  *sqlx.Stmt `yesql:"exist_topic_user"`
	FollowPinTags   *sqlx.Stmt `yesql:"follow_pin_tags"`
	FollowTags      *sqlx.Stmt `yesql:"follow_tags"`
	FollowTopic     *sqlx.Stmt `yesql:"follow_topic"`
	HotTags         *sqlx.Stmt `yesql:"hot_tags"`
	InsertTag       *sqlx.Stmt `yesql:"insert_tag"`
	NewestTags      *sqlx.Stmt `yesql:"newest_tags"`
	PinTopic        *sqlx.Stmt `yesql:"pin_topic"`
	StickTopic      *sqlx.Stmt `yesql:"stick_topic"`
	TagsByKeywordA  *sqlx.Stmt `yesql:"tags_by_keyword_a"`
	TagsByKeywordB  *sqlx.Stmt `yesql:"tags_by_keyword_b"`
	TopicIsPin      *sqlx.Stmt `yesql:"topic_is_pin"`
	TopicIsTop      *sqlx.Stmt `yesql:"topic_is_top"`
	UnfollowTopic   *sqlx.Stmt `yesql:"unfollow_topic"`
}

type TweetA struct {
	yesql.Namespace     `yesql:"tweet_a"`
	AttachmentByTweetId string `yesql:"attachment_by_tweet_id"`
	FavoriteByTweetId   string `yesql:"favorite_by_tweet_id"`
	ReactionByTweetId   string `yesql:"reaction_by_tweet_id"`
	TweetInfoById       string `yesql:"tweet_info_by_id"`
	TweetItemById       string `yesql:"tweet_item_by_id"`
	UserFavorites       string `yesql:"user_favorites"`
	UserInfo            string `yesql:"user_info"`
	UserReactions       string `yesql:"user_reactions"`
	UserTweetsByAdmin   string `yesql:"user_tweets_by_admin"`
	UserTweetsByFriend  string `yesql:"user_tweets_by_friend"`
	UserTweetsByGuest   string `yesql:"user_tweets_by_guest"`
	UserTweetsBySelf    string `yesql:"user_tweets_by_self"`
}

type TweetHelpA struct {
	yesql.Namespace `yesql:"tweet_help_a"`
	UserInfo        string `yesql:"user_info"`
}

type TweetManageA struct {
	yesql.Namespace `yesql:"tweet_manage_a"`
	UserInfo        string `yesql:"user_info"`
}

func BuildShipIndexA(p PreparexBuilder) (obj *ShipIndexA, err error) {
	obj = &ShipIndexA{
		UserInfo: p.QueryHook(_ShipIndexA_UserInfo),
	}
	return
}

func BuildSimpleIndexA(p PreparexBuilder) (obj *SimpleIndexA, err error) {
	obj = &SimpleIndexA{
		UserInfo: p.QueryHook(_SimpleIndexA_UserInfo),
	}
	return
}

func BuildTopicA(p PreparexBuilder, ctx ...context.Context) (obj *TopicA, err error) {
	var c context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		c = ctx[0]
	} else {
		c = context.Background()
	}
	obj = &TopicA{
		DecrTagsById: p.QueryHook(_TopicA_DecrTagsById),
		IncrTagsById: p.QueryHook(_TopicA_IncrTagsById),
		TagsByIdA:    p.QueryHook(_TopicA_TagsByIdA),
		TagsByIdB:    p.QueryHook(_TopicA_TagsByIdB),
		TagsForIncr:  p.QueryHook(_TopicA_TagsForIncr),
		TopicInfos:   p.QueryHook(_TopicA_TopicInfos),
	}
	if obj.ExistTopicUser, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_ExistTopicUser))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_ExistTopicUser error: %w", err)
	}
	if obj.FollowPinTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_FollowPinTags))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_FollowPinTags error: %w", err)
	}
	if obj.FollowTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_FollowTags))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_FollowTags error: %w", err)
	}
	if obj.FollowTopic, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_FollowTopic))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_FollowTopic error: %w", err)
	}
	if obj.HotTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_HotTags))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_HotTags error: %w", err)
	}
	if obj.InsertTag, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_InsertTag))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_InsertTag error: %w", err)
	}
	if obj.NewestTags, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_NewestTags))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_NewestTags error: %w", err)
	}
	if obj.PinTopic, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_PinTopic))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_PinTopic error: %w", err)
	}
	if obj.StickTopic, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_StickTopic))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_StickTopic error: %w", err)
	}
	if obj.TagsByKeywordA, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_TagsByKeywordA))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_TagsByKeywordA error: %w", err)
	}
	if obj.TagsByKeywordB, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_TagsByKeywordB))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_TagsByKeywordB error: %w", err)
	}
	if obj.TopicIsPin, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_TopicIsPin))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_TopicIsPin error: %w", err)
	}
	if obj.TopicIsTop, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_TopicIsTop))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_TopicIsTop error: %w", err)
	}
	if obj.UnfollowTopic, err = p.PreparexContext(c, p.Rebind(p.QueryHook(_TopicA_UnfollowTopic))); err != nil {
		return nil, fmt.Errorf("prepare _TopicA_UnfollowTopic error: %w", err)
	}
	return
}

func BuildTweetA(p PreparexBuilder) (obj *TweetA, err error) {
	obj = &TweetA{
		AttachmentByTweetId: p.QueryHook(_TweetA_AttachmentByTweetId),
		FavoriteByTweetId:   p.QueryHook(_TweetA_FavoriteByTweetId),
		ReactionByTweetId:   p.QueryHook(_TweetA_ReactionByTweetId),
		TweetInfoById:       p.QueryHook(_TweetA_TweetInfoById),
		TweetItemById:       p.QueryHook(_TweetA_TweetItemById),
		UserFavorites:       p.QueryHook(_TweetA_UserFavorites),
		UserInfo:            p.QueryHook(_TweetA_UserInfo),
		UserReactions:       p.QueryHook(_TweetA_UserReactions),
		UserTweetsByAdmin:   p.QueryHook(_TweetA_UserTweetsByAdmin),
		UserTweetsByFriend:  p.QueryHook(_TweetA_UserTweetsByFriend),
		UserTweetsByGuest:   p.QueryHook(_TweetA_UserTweetsByGuest),
		UserTweetsBySelf:    p.QueryHook(_TweetA_UserTweetsBySelf),
	}
	return
}

func BuildTweetHelpA(p PreparexBuilder) (obj *TweetHelpA, err error) {
	obj = &TweetHelpA{
		UserInfo: p.QueryHook(_TweetHelpA_UserInfo),
	}
	return
}

func BuildTweetManageA(p PreparexBuilder) (obj *TweetManageA, err error) {
	obj = &TweetManageA{
		UserInfo: p.QueryHook(_TweetManageA_UserInfo),
	}
	return
}
