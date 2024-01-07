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

type AlipayPub interface {
	_default_

	AlipayNotify(*web.AlipayNotifyReq) mir.Error

	mustEmbedUnimplementedAlipayPubServant()
}

// RegisterAlipayPubServant register AlipayPub servant to gin
func RegisterAlipayPubServant(e *gin.Engine, s AlipayPub) {
	router := e.Group("v1")

	// register routes info to router
	router.Handle("POST", "/alipay/notify", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(web.AlipayNotifyReq)
		var bv _binding_ = req
		if err := bv.Bind(c); err != nil {
			s.Render(c, nil, err)
			return
		}
		s.Render(c, nil, s.AlipayNotify(req))
	})
}

// UnimplementedAlipayPubServant can be embedded to have forward compatible implementations.
type UnimplementedAlipayPubServant struct{}

func (UnimplementedAlipayPubServant) AlipayNotify(req *web.AlipayNotifyReq) mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedAlipayPubServant) mustEmbedUnimplementedAlipayPubServant() {}
