package xrequest

import (
	"context"
	"gorm.io/gorm/utils"
	"strings"
)

//来源类型
type FromType string

const (
	FromTypeApp     FromType = "app"     //APP
	FromTypePc      FromType = "pc"      //PC
	FromTypeH5      FromType = "h5"      //H5
	FromTypeRoutine FromType = "routine" //小程序
	FromTypeWechat  FromType = "wechat"  //微信
	FromTypeCashier FromType = "cashier" //收银台
)

//小程序类型
type RoutineType string

const (
	RoutineTypeWeixin   RoutineType = "MP-WEIXIN"   //微信小程序
	RoutineTypeAlipay   RoutineType = "MP-ALIPAY"   //支付宝小程序
	RoutineTypeBaidu    RoutineType = "MP-BAIDU"    //百度小程序
	RoutineTypeTouTiao  RoutineType = "MP-TOUTIAO"  //头条小程序
	RoutineTypeKuaiShou RoutineType = "MP-KUAISHOU" //快手小程序
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

//是否微信小程序（基于RoutineType）
func IsMPWeixin(ctx context.Context) bool {
	return strings.ToUpper(GetRoutineType(ctx)) == "MP-WEIXIN"
}

//是否支付宝小程序（基于RoutineType）
func IsMPAlipay(ctx context.Context) bool {
	return strings.ToUpper(GetRoutineType(ctx)) == "MP-ALIPAY"
}

//是否头条小程序（基于RoutineType）
func IsMPToutiao(ctx context.Context) bool {
	return strings.ToUpper(GetRoutineType(ctx)) == "MP-TOUTIAO"
}

//是否百度小程序（基于RoutineType）
func IsMPBaidu(ctx context.Context) bool {
	return strings.ToUpper(GetRoutineType(ctx)) == "MP-BAIDU"
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
