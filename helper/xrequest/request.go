package xrequest

import (
	"context"
	"gorm.io/gorm/utils"
	"strings"
)

//获得来源类型
func GetFromType(ctx context.Context) string {
	val := ctx.Value("From-Type")
	if val != nil {
		v := utils.ToString(val)
		return v
	}
	return ""
}

//当前访问端（基于FromType）
func IsTerminal(ctx context.Context, terminal string) bool {
	return strings.ToLower(GetFromType(ctx)) == terminal
}

//是否是H5端（基于FromType）
func IsH5(ctx context.Context) bool {
	return strings.ToLower(GetFromType(ctx)) == "h5"
}

//是否是微信端（基于FromType）
func IsWechat(ctx context.Context) bool {
	return strings.ToLower(GetFromType(ctx)) == "wechat"
}

//是否是小程序端（基于FromType）
func IsRoutine(ctx context.Context) bool {
	return strings.ToLower(GetFromType(ctx)) == "routine"
}

//是否是app端（基于FromType）
func IsApp(ctx context.Context) bool {
	return strings.ToLower(GetFromType(ctx)) == "app"
}

//是否是pc端（基于FromType）
func IsPc(ctx context.Context) bool {
	return strings.ToLower(GetFromType(ctx)) == "pc"
}

//获取小程序类型
func GetRoutineType(ctx context.Context) string {
	val := ctx.Value("Routine-Type")
	if val != nil {
		v := utils.ToString(val)
		return v
	}
	return ""
}

//是否微信环境 （基于UserAgent）
func IsWeixinEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "MicroMessenger") != -1 {
		return true
	}
	return false
}

//是否钉钉环境 （基于UserAgent）
func IsDingDingEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "DingTalk") != -1 {
		return true
	}
	return false
}

//是否QQ环境 （基于UserAgent）
func IsQQEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "QQ/") != -1 {
		return true
	}
	return false
}

//是否支付宝环境 （基于UserAgent）
func IsAlipayEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "AlipayClient") != -1 {
		return true
	}
	return false
}

//是否新浪微博环境 （基于UserAgent）
func IsWeiboEnv(ctx context.Context) bool {
	userAgent := GetHttpUserAgent(ctx)
	if userAgent != "" && strings.Index(userAgent, "Weibo") != -1 {
		return true
	}
	return false
}

//获取UserAgent
func GetHttpUserAgent(ctx context.Context) string {
	val := ctx.Value("Http-User-Agent")
	if val != nil {
		v := utils.ToString(val)
		return v
	}
	return ""
}
