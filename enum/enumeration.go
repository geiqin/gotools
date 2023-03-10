package enum

//bool类型
type BoolType string

const (
	BoolTypeTrue  BoolType = "0" //false
	BoolTypeFalse BoolType = "1" //true
)

//应用程序
type Application string

const (
	ApplicationApp Application = "app" //APP端
	ApplicationWeb Application = "web" //WEB端
)

//VIP类型
type VipType string

const (
	VipTypeFree VipType = "0" //免费会员
	VipTypePay  VipType = "1" //付费会员
)

//访问终端类型
type ClientType string

const (
	ClientTypeAppPC           ClientType = "pc"       //电脑PC
	ClientTypeAppH5           ClientType = "h5"       //手机H5
	ClientTypeAppMiniWeixin   ClientType = "weixin"   //微信小程序
	ClientTypeAppMiniAlipay   ClientType = "alipay"   //支付宝小程序
	ClientTypeAppMiniBaidu    ClientType = "baidu"    //百度小程序
	ClientTypeAppMiniQQ       ClientType = "qq"       //QQ小程序
	ClientTypeAppMiniToutiao  ClientType = "toutiao"  //今日头条小程序
	ClientTypeAppMiniKuaishou ClientType = "kuaishou" //快手小程序
)

//注册来源
type RegSourceType string

const (
	RegSourceTypeHand            RegSourceType = "hand"     //手动添加(无法登录)
	RegSourceTypeAccount         RegSourceType = "account"  //账号密码登录
	RegSourceTypeMobile          RegSourceType = "mobile"   //手机一键登录/短信验证登录
	RegSourceTypeAppMiniWeixin   RegSourceType = "weixin"   //微信小程序
	RegSourceTypeAppMiniBaidu    RegSourceType = "baidu"    //百度小程序
	RegSourceTypeAppMiniAlipay   RegSourceType = "alipay"   //支付宝小程序
	RegSourceTypeAppMiniToutiao  RegSourceType = "toutiao"  //今日头条小程序
	RegSourceTypeAppMiniKuaishou RegSourceType = "kuaishou" //快手小程序
	RegSourceTypeAppMiniQQ       RegSourceType = "qq"       //QQ小程序
	RegSourceTypeAlipay          RegSourceType = "alipay"   //支付宝登录
	RegisterSourceDouYin         RegSourceType = "douyin"   //抖音登录
	RegSourceTypeWeibo           RegSourceType = "weibo"    //微博登录
)
