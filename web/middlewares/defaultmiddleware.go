package middlewares

import (
	"github.com/tandonraghav/go-logging/utilities"
	logger "github.com/tandonraghav/go-logging/web/logging"
	"context"
	"net/http"
	"runtime/debug"
)


func DefaultMiddleware(next http.Handler) http.Handler {
	fn:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r=r.WithContext(utilities.WithContext(r.Context()))
		r=r.WithContext(logger.WithLogger(r.Context()))
		applyMid := applyMid(w, r)
		if applyMid {
			logger.GetLogger(r.Context()).Info( "applyMid=",applyMid)
			defer excHandler(r.Context())()
			next.ServeHTTP(w,r)
			logger.GetLogger(r.Context()).Info( "Exiting Request")

		} else {
			return
		}
	})
	return fn
}

func excHandler(ctx context.Context) func(){
	fn:=func() {
		if re := recover(); re != nil {
			logger.GetLogger(ctx).Error( "Exception occured ", re)
			logger.GetLogger(ctx).Error( "Exception occured ", (string)(debug.Stack()))
		}

	}
	return fn
}

func applyMid(w http.ResponseWriter, r *http.Request) bool {
	return true
}
