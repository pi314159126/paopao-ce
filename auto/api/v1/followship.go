// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.0.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

type Followship interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	ListFollowers(*web.ListFollowersReq) (*web.ListFollowersResp, mir.Error)
	ListFollowings(*web.ListFollowingsReq) (*web.ListFollowingsResp, mir.Error)
	DeleteFollowing(*web.DeleteFollowingReq) mir.Error
	AddFollowing(*web.AddFollowingReq) mir.Error

	mustEmbedUnimplementedFollowshipServant()
}

// RegisterFollowshipServant register Followship servant to gin
func RegisterFollowshipServant(e *gin.Engine, s Followship) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/follower/list", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(web.ListFollowersReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.ListFollowers(req)
		s.Render(c, resp, err)
	})
	router.Handle("GET", "/following/list", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(web.ListFollowingsReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.ListFollowings(req)
		s.Render(c, resp, err)
	})
	router.Handle("POST", "/following/delete", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(web.DeleteFollowingReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.DeleteFollowing(req))
	})
	router.Handle("POST", "/following/add", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		req := new(web.AddFollowingReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.AddFollowing(req))
	})
}

// UnimplementedFollowshipServant can be embedded to have forward compatible implementations.
type UnimplementedFollowshipServant struct{}

func (UnimplementedFollowshipServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedFollowshipServant) ListFollowers(req *web.ListFollowersReq) (*web.ListFollowersResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFollowshipServant) ListFollowings(req *web.ListFollowingsReq) (*web.ListFollowingsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFollowshipServant) DeleteFollowing(req *web.DeleteFollowingReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFollowshipServant) AddFollowing(req *web.AddFollowingReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedFollowshipServant) mustEmbedUnimplementedFollowshipServant() {}
