package igin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"
)

const (
	UserId  = "trace_user_id"
	TraceId = "trace_id"
)

type (
	transCtx   struct{}
	userIDCtx  struct{}
	traceIDCtx struct{}
)

type gi struct {
	g *gin.Context
	c context.Context
}

func NewG(g *gin.Context) *gi {
	return &gi{g: g}
}

//设置uid
func (g *gi) SetContextUserId(uid string) {
	g.g.Set(UserId, uid)
}

//获取uid
func (g *gi) GetContextUserId() string {
	return g.g.GetString(UserId)
}

//获取链路追踪id
func (g *gi) GetTraceId() string {
	return g.g.GetString(TraceId)
}

//NewTraceId 创建追踪id的上下文
func (g *gi) NewTraceId(traceId string) context.Context {

	return context.WithValue(g.c, transCtx{}, traceId)
}

//gin.Context--->context.context()

//todo
func (g *gi) NewContext() context.Context {
	parent := context.Background()
	if v := g.GetTraceId(); v != "" {
		g.c = parent
		parent = g.NewTraceId(v)
	}
}
