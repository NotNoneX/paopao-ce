// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sakila

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/alimy/cfg"
	"github.com/alimy/yesql"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/dao/cache"
	"github.com/rocboss/paopao-ce/internal/dao/security"
	"github.com/sirupsen/logrus"
)

var (
	_ core.DataService = (*dataSrv)(nil)
	_ core.VersionInfo = (*dataSrv)(nil)

	_ core.WebDataServantA = (*webDataSrvA)(nil)
	_ core.VersionInfo     = (*webDataSrvA)(nil)

	_onceInitial sync.Once
)

type dataSrv struct {
	core.IndexPostsService
	core.WalletService
	core.MessageService
	core.TopicService
	core.TweetService
	core.TweetManageService
	core.TweetHelpService
	core.CommentService
	core.CommentManageService
	core.UserManageService
	core.ContactManageService
	core.SecurityService
	core.AttachmentCheckService
}

type webDataSrvA struct {
	core.TopicServantA
	core.TweetServantA
	core.TweetManageServantA
	core.TweetHelpServantA
}

func NewDataService() (core.DataService, core.VersionInfo) {
	lazyInitial()

	var (
		v   core.VersionInfo
		cis core.CacheIndexService
		ips core.IndexPostsService
	)
	pvs := security.NewPhoneVerifyService()
	ams := NewAuthorizationManageService()
	ths := newTweetHelpService(_db)

	// initialize core.IndexPostsService
	if cfg.If("Friendship") {
		ips = newFriendIndexService(_db, ams, ths)
	} else if cfg.If("Followship") {
		ips = newFollowIndexService(_db, ths)
	} else if cfg.If("Lightship") {
		ips = newLightIndexService(_db, ths)
	} else {
		// default use lightship post index service
		ips = newLightIndexService(_db, ths)
	}

	// initialize core.CacheIndexService
	if cfg.If("SimpleCacheIndex") {
		// simpleCache use special post index service
		ips = newSimpleIndexPostsService(_db, ths)
		cis, v = cache.NewSimpleCacheIndexService(ips)
	} else if cfg.If("BigCacheIndex") {
		// TODO: make cache index post in different scence like friendship/followship/lightship
		cis, v = cache.NewBigCacheIndexService(ips, ams)
	} else {
		cis, v = cache.NewNoneCacheIndexService(ips)
	}
	logrus.Infof("use %s as cache index service by version: %s", v.Name(), v.Version())

	query := yesql.MustParseBytes(yesqlBytes)
	ds := &dataSrv{
		IndexPostsService:      cis,
		WalletService:          newWalletService(_db),
		MessageService:         newMessageService(_db),
		TopicService:           newTopicService(_db, query),
		TweetService:           newTweetService(_db),
		TweetManageService:     newTweetManageService(_db, cis),
		TweetHelpService:       newTweetHelpService(_db),
		CommentService:         newCommentService(_db),
		CommentManageService:   newCommentManageService(_db),
		UserManageService:      newUserManageService(_db),
		ContactManageService:   newContactManageService(_db),
		SecurityService:        newSecurityService(_db, pvs),
		AttachmentCheckService: security.NewAttachmentCheckService(),
	}
	return ds, ds
}

func NewWebDataServantA() (core.WebDataServantA, core.VersionInfo) {
	lazyInitial()
	// db := conf.MustSqlxDB()
	ds := &webDataSrvA{}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	lazyInitial()
	return newAuthorizationManageService(_db)
}

func (s *dataSrv) Name() string {
	return "Sqlx"
}

func (s *dataSrv) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *webDataSrvA) Name() string {
	return "Sqlx"
}

func (s *webDataSrvA) Version() *semver.Version {
	return semver.MustParse("v0.0.0")
}

// lazyInitial do some package lazy initialize for performance
func lazyInitial() {
	_onceInitial.Do(func() {
		initSqlxDB()
	})
}
