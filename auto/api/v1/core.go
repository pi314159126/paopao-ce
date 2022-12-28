// Code generated by go-mir. DO NOT EDIT.

package v1

import (
	"net/http"

	"github.com/alimy/mir/v3"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
	"github.com/rocboss/paopao-ce/internal/servants/base"
)

type Core interface {
	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	TweetCollectionStatus() mir.Error
	TweetStarStatus() mir.Error
	SuggestTags() mir.Error
	SuggestUsers() mir.Error
	ChangeAvatar(*web.ChangeAvatarReq) mir.Error
	ChangeNickname() mir.Error
	ChangePassword() mir.Error
	UserPhoneBind() mir.Error
	GetStars() mir.Error
	GetCollections() mir.Error
	SendUserWhisper() mir.Error
	ReadMessage(*web.ReadMessageReq) mir.Error
	GetMessages(*web.GetMessagesReq) (*base.PageResp, mir.Error)
	GetUnreadMsgCount(*web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, mir.Error)
	GetUserInfo(*web.UserInfoReq) (*web.UserInfoResp, mir.Error)
	SyncSearchIndex(*web.SyncSearchIndexReq) mir.Error

	mustEmbedUnimplementedCoreServant()
}

type CoreBinding interface {
	BindChangeAvatar(*gin.Context) (*web.ChangeAvatarReq, mir.Error)
	BindReadMessage(*gin.Context) (*web.ReadMessageReq, mir.Error)
	BindGetMessages(*gin.Context) (*web.GetMessagesReq, mir.Error)
	BindGetUnreadMsgCount(*gin.Context) (*web.GetUnreadMsgCountReq, mir.Error)
	BindGetUserInfo(*gin.Context) (*web.UserInfoReq, mir.Error)
	BindSyncSearchIndex(*gin.Context) (*web.SyncSearchIndexReq, mir.Error)

	mustEmbedUnimplementedCoreBinding()
}

type CoreRender interface {
	RenderTweetCollectionStatus(*gin.Context, mir.Error)
	RenderTweetStarStatus(*gin.Context, mir.Error)
	RenderSuggestTags(*gin.Context, mir.Error)
	RenderSuggestUsers(*gin.Context, mir.Error)
	RenderChangeAvatar(*gin.Context, mir.Error)
	RenderChangeNickname(*gin.Context, mir.Error)
	RenderChangePassword(*gin.Context, mir.Error)
	RenderUserPhoneBind(*gin.Context, mir.Error)
	RenderGetStars(*gin.Context, mir.Error)
	RenderGetCollections(*gin.Context, mir.Error)
	RenderSendUserWhisper(*gin.Context, mir.Error)
	RenderReadMessage(*gin.Context, mir.Error)
	RenderGetMessages(*gin.Context, *base.PageResp, mir.Error)
	RenderGetUnreadMsgCount(*gin.Context, *web.GetUnreadMsgCountResp, mir.Error)
	RenderGetUserInfo(*gin.Context, *web.UserInfoResp, mir.Error)
	RenderSyncSearchIndex(*gin.Context, mir.Error)

	mustEmbedUnimplementedCoreRender()
}

// RegisterCoreServant register Core servant to gin
func RegisterCoreServant(e *gin.Engine, s Core, b CoreBinding, r CoreRender) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/post/collection", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderTweetCollectionStatus(c, s.TweetCollectionStatus())
	})

	router.Handle("GET", "/post/star", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderTweetStarStatus(c, s.TweetStarStatus())
	})

	router.Handle("GET", "/suggest/tags", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderSuggestTags(c, s.SuggestTags())
	})

	router.Handle("GET", "/suggest/users", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderSuggestUsers(c, s.SuggestUsers())
	})

	router.Handle("POST", "/user/avatar", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindChangeAvatar(c)
		if err != nil {
			r.RenderChangeAvatar(c, err)
			return
		}
		r.RenderChangeAvatar(c, s.ChangeAvatar(req))
	})

	router.Handle("POST", "/user/nickname", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderChangeNickname(c, s.ChangeNickname())
	})

	router.Handle("POST", "/user/password", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderChangePassword(c, s.ChangePassword())
	})

	router.Handle("POST", "/user/phone", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderUserPhoneBind(c, s.UserPhoneBind())
	})

	router.Handle("GET", "/user/stars", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderGetStars(c, s.GetStars())
	})

	router.Handle("GET", "/user/collections", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderGetCollections(c, s.GetCollections())
	})

	router.Handle("POST", "/user/whisper", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		r.RenderSendUserWhisper(c, s.SendUserWhisper())
	})

	router.Handle("POST", "/user/message/read", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindReadMessage(c)
		if err != nil {
			r.RenderReadMessage(c, err)
			return
		}
		r.RenderReadMessage(c, s.ReadMessage(req))
	})

	router.Handle("GET", "/user/messages", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindGetMessages(c)
		if err != nil {
			r.RenderGetMessages(c, nil, err)
			return
		}
		resp, err := s.GetMessages(req)
		r.RenderGetMessages(c, resp, err)
	})

	router.Handle("GET", "/user/msgcount/unread", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindGetUnreadMsgCount(c)
		if err != nil {
			r.RenderGetUnreadMsgCount(c, nil, err)
			return
		}
		resp, err := s.GetUnreadMsgCount(req)
		r.RenderGetUnreadMsgCount(c, resp, err)
	})

	router.Handle("GET", "/user/info", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindGetUserInfo(c)
		if err != nil {
			r.RenderGetUserInfo(c, nil, err)
			return
		}
		resp, err := s.GetUserInfo(req)
		r.RenderGetUserInfo(c, resp, err)
	})

	router.Handle("GET", "/sync/index", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req, err := b.BindSyncSearchIndex(c)
		if err != nil {
			r.RenderSyncSearchIndex(c, err)
			return
		}
		r.RenderSyncSearchIndex(c, s.SyncSearchIndex(req))
	})

}

// UnimplementedCoreServant can be embedded to have forward compatible implementations.
type UnimplementedCoreServant struct {
}

func (UnimplementedCoreServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedCoreServant) TweetCollectionStatus() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) TweetStarStatus() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) SuggestTags() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) SuggestUsers() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) ChangeAvatar(req *web.ChangeAvatarReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) ChangeNickname() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) ChangePassword() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) UserPhoneBind() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) GetStars() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) GetCollections() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) SendUserWhisper() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) ReadMessage(req *web.ReadMessageReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) GetMessages(req *web.GetMessagesReq) (*base.PageResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) GetUnreadMsgCount(req *web.GetUnreadMsgCountReq) (*web.GetUnreadMsgCountResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) GetUserInfo(req *web.UserInfoReq) (*web.UserInfoResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) SyncSearchIndex(req *web.SyncSearchIndexReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedCoreServant) mustEmbedUnimplementedCoreServant() {}

// UnimplementedCoreRender can be embedded to have forward compatible implementations.
type UnimplementedCoreRender struct {
	RenderAny func(*gin.Context, any, mir.Error)
}

func (r *UnimplementedCoreRender) RenderTweetCollectionStatus(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderTweetStarStatus(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderSuggestTags(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderSuggestUsers(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderChangeAvatar(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderChangeNickname(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderChangePassword(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderUserPhoneBind(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderGetStars(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderGetCollections(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderSendUserWhisper(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderReadMessage(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) RenderGetMessages(c *gin.Context, data *base.PageResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedCoreRender) RenderGetUnreadMsgCount(c *gin.Context, data *web.GetUnreadMsgCountResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedCoreRender) RenderGetUserInfo(c *gin.Context, data *web.UserInfoResp, err mir.Error) {
	r.RenderAny(c, data, err)
}

func (r *UnimplementedCoreRender) RenderSyncSearchIndex(c *gin.Context, err mir.Error) {
	r.RenderAny(c, nil, err)
}

func (r *UnimplementedCoreRender) mustEmbedUnimplementedCoreRender() {}

// UnimplementedCoreBinding can be embedded to have forward compatible implementations.
type UnimplementedCoreBinding struct {
	BindAny func(*gin.Context, any) mir.Error
}

func (b *UnimplementedCoreBinding) BindChangeAvatar(c *gin.Context) (*web.ChangeAvatarReq, mir.Error) {
	obj := new(web.ChangeAvatarReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) BindReadMessage(c *gin.Context) (*web.ReadMessageReq, mir.Error) {
	obj := new(web.ReadMessageReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) BindGetMessages(c *gin.Context) (*web.GetMessagesReq, mir.Error) {
	obj := new(web.GetMessagesReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) BindGetUnreadMsgCount(c *gin.Context) (*web.GetUnreadMsgCountReq, mir.Error) {
	obj := new(web.GetUnreadMsgCountReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) BindGetUserInfo(c *gin.Context) (*web.UserInfoReq, mir.Error) {
	obj := new(web.UserInfoReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) BindSyncSearchIndex(c *gin.Context) (*web.SyncSearchIndexReq, mir.Error) {
	obj := new(web.SyncSearchIndexReq)
	err := b.BindAny(c, obj)
	return obj, err
}

func (b *UnimplementedCoreBinding) mustEmbedUnimplementedCoreBinding() {}
