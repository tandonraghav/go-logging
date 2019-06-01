package utilities

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type RequestCtxKey struct{
	name string
}

var (
	CtxKey = &RequestCtxKey{"RequestContext"}
)

type RequestContext struct {
	ReqID string
	Username string
}

func NewContext() *RequestContext{
	return &RequestContext{}
}

func (x *RequestContext) AddReqID() {
	x.ReqID=createRequestID()
}

func createRequestID() string {
	var sb strings.Builder
	t := time.Now()
	str := t.Format("20060102")
	sb.WriteString(str)
	if t.Hour() < 10 {
		sb.WriteString("0")
		sb.WriteString(strconv.Itoa(t.Hour()))
	} else {
		sb.WriteString(strconv.Itoa(t.Hour()))
	}

	if t.Minute() < 10 {
		sb.WriteString("0")
		sb.WriteString(strconv.Itoa(t.Minute()))
	} else {
		sb.WriteString(strconv.Itoa(t.Minute()))
	}
	sb.WriteString(viper.GetString("server.serverId"))
	sb.WriteString(strconv.FormatInt(t.UnixNano(), 10))
	return sb.String()
}

func GetRequestID(ctx context.Context) string {
	rCtx:=GetRequestContext(ctx)
	if rCtx!=nil{
		return (*rCtx).ReqID
	}
	return "NA"
}

func GetValueFromReqContext(ctx context.Context,key string) string {
	fmt.Print(key)
	rCtx:=GetRequestContext(ctx)
	if rCtx!=nil{
		r := reflect.ValueOf(rCtx)
		f := reflect.Indirect(r).FieldByName(key)

		if f.String()!=""{
			return f.String()
		}
	}
	return "NA"
}

func GetRequestContext(ctx context.Context) *RequestContext {
	if ctx!=nil{
		if v, ok:=ctx.Value(CtxKey).(*RequestContext); ok{
			return v
		}
	}
	return nil
}

func WithContext(ctx context.Context) context.Context{
	rCtx:=NewContext()
	rCtx.AddReqID()
	return context.WithValue(ctx,CtxKey,rCtx)
}

func AddUsernameToContext(ctx context.Context,username string){
	GetRequestContext(ctx).Username=username
}