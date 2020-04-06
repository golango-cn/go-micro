package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/plugins/session"
	auth "go-micro/go-micro-part03/proto"
	"net/http"
)


// AuthWrapper 认证wrapper
func AuthWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie("remember-me-token")
		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}

		sess := session.GetSession2(w, r)
		if sess.ID != "" {
			// 检测是否通过验证
			if sess.Values["valid"] != nil {
				h.ServeHTTP(w, r)
				return
			} else {
				userId := sess.Values["userId"].(int32)
				if userId != 0 {
					rsp, err := authClient.GetToken(context.TODO(), &auth.TokenRequest{
						UserId: uint64(userId),
					})
					if err != nil {
						log.Errorf("[AuthWrapper]，err：%s", err)
						http.Error(w, "非法请求", 400)
						return
					}

					// token不一致
					if rsp.Token != ck.Value {
						log.Errorf("[AuthWrapper]，token不一致")
						http.Error(w, "非法请求", 400)
						return
					}
				} else {
					log.Errorf("[AuthWrapper]，session不合法，无用户id")
					http.Error(w, "非法请求", 400)
					return
				}
			}
		} else {
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// AuthWrapper 认证wrapper
func AuthWrapper2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, _ := r.Cookie("remember-me-token")

		log.Info("1", ck)

		// token不存在，则状态异常，无权限
		if ck == nil {
			http.Error(w, "非法请求", 400)
			return
		}

		//token, err := parseToken(ck.Value)
		//if err != nil {
		//	http.Error(w, "非法请求", 400)
		//	return
		//}

		sess := session.GetSession(r)

		if sess == nil {
			log.Error("请求缺少session-id-")
			http.Error(w, "非法请求", 400)
			return
		}

		userId := uint64(sess.Values["userId"].(float64))
		rsp, err := authClient.GetToken(context.TODO(), &auth.TokenRequest{
			UserId: userId,
		})
		if err != nil {
			log.Errorf("[AuthWrapper]，err：%s", err)
			http.Error(w, "非法请求", 400)
			return
		}

		// token不一致
		if rsp.Token != ck.Value {
			log.Errorf("[AuthWrapper]，token不一致")
			http.Error(w, "非法请求", 400)
			return
		}

		h.ServeHTTP(w, r)
	})
}
