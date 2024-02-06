// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.1.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/model/web"
)

type Trends interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	GetIndexTrends(*gin.Context, *web.GetIndexTrendsReq) (*web.GetIndexTrendsResp, mir.Error)

	mustEmbedUnimplementedTrendsServant()
}

// RegisterTrendsServant register Trends servant to gin
func RegisterTrendsServant(e *gin.Engine, s Trends) {
	router := e.Group("v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("GET", "/trends/index", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.GetIndexTrendsReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.GetIndexTrends(c, req)
		if err != nil {
			s.Render(c, nil, err)
			return
		}
		var rv _render_ = resp
		rv.Render(c)
	})
}

// UnimplementedTrendsServant can be embedded to have forward compatible implementations.
type UnimplementedTrendsServant struct{}

func (UnimplementedTrendsServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedTrendsServant) GetIndexTrends(c *gin.Context, req *web.GetIndexTrendsReq) (*web.GetIndexTrendsResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedTrendsServant) mustEmbedUnimplementedTrendsServant() {}
