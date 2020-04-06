package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro/go-micro-part03/basic/config"
	"go-micro/go-micro-part03/plugins/session"
	auth "go-micro/go-micro-part03/proto"
	"net/http"
)


// parseToken 解析token
func  parseToken(tk string) (c jwt.MapClaims, err error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.Jwt.SecretKey), nil
	})

	// jwt 框架自带了一些检测，如过期，发布者错误等
	if err != nil {
		switch e := err.(type) {
		case *jwt.ValidationError:
			switch e.Errors {
			case jwt.ValidationErrorExpired:
				return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
			default:
				break
			}
			break
		default:
			break
		}

		return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
	}

	// 检测合法
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("[parseToken] 不合法的token")
	}

	return claims, nil
}

// AuthWrapper 认证wrapper
//func AuthWrapper(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		ck, _ := r.Cookie("remember-me-token")
//		// token不存在，则状态异常，无权限
//		if ck == nil {
//			log.Error("token不存在")
//			http.Error(w, "非法请求", 400)
//			return
//		}
//
//		token, err := parseToken(ck.Value)
//		if err != nil {
//			http.Error(w, "非法请求", 400)
//			return
//		}
//
//		userId , _ := strconv.ParseUint(token["sub"].(string), 10, 64)
//
//		if userId != 0 {
//			rsp, err := authClient.GetToken(context.TODO(), &auth.TokenRequest{
//				UserId: userId,
//			})
//			if err != nil {
//				log.Errorf("[AuthWrapper]，err：%s", err)
//				http.Error(w, "非法请求", 400)
//				return
//			}
//
//			// token不一致
//			if rsp.Token != ck.Value {
//				log.Errorf("[AuthWrapper]，token不一致")
//				http.Error(w, "非法请求", 400)
//				return
//			}
//		} else {
//			log.Errorf("[AuthWrapper]，session不合法，无用户id")
//			http.Error(w, "非法请求", 400)
//			return
//		}
//
//		h.ServeHTTP(w, r)
//		return
//
//		sess := session.GetSession(w, r)
//		if sess.ID != "" {
//			// 检测是否通过验证
//			if sess.Values["valid"] != nil {
//				h.ServeHTTP(w, r)
//				return
//			} else {
//				userId := sess.Values["userId"].(uint64)
//				if userId != 0 {
//					rsp, err := authClient.GetToken(context.TODO(), &auth.TokenRequest{
//						UserId: userId,
//					})
//					if err != nil {
//						log.Logf("[AuthWrapper]，err：%s", err)
//						http.Error(w, "非法请求", 400)
//						return
//					}
//
//					// token不一致
//					if rsp.Token != ck.Value {
//						log.Error("[AuthWrapper]，token不一致")
//						http.Error(w, "非法请求", 400)
//						return
//					}
//				} else {
//					log.Error("[AuthWrapper]，session不合法，无用户id")
//					http.Error(w, "非法请求", 400)
//					return
//				}
//			}
//		} else {
//			http.Error(w, "非法请求", 400)
//			return
//		}
//
//		h.ServeHTTP(w, r)
//	})
//}


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
				userId := sess.Values["userId"].(uint64)
				if userId != 0 {
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

