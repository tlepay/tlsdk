# GO SDK 简介
    为了提高客户接入的便捷性，本系统提供 SDK 方式接入，使用本 SDK 将极大的简化开发者的工作，开发者将无需考虑通信、签名、验签等，只需要关注业务参数的输入。

## SDK 项目结构说明
- tlsdk -- SDK核心包, 包含通信, 加解签, 接口参数对象等
- config -- 商户配置
- demo -- 接口调用/参数赋值演示demo

## SDK 接入说明 
以下两种方式任选其一
1. 直接在go.mod中引用(require github.com/tlepay/tlsdk [version])
2. 直接下载源码文件, 将TLSDK(SDK核心包)源码放入项目中


## SDK 使用说明
    接口命名直接根据接口URL来命名, 方便用户使用, 需要使用某接口时, 可直接使用接口中文名, 或接口URL(驼峰格式)进行搜索, 找到对应的struct, demo等

1. 配置初始化
- NewConfigWithFile初始化为全局配置(多商户模式下, 可初始化多份)
- 第一个参数 配置文件路径
```
config, _ := tlsdk.NewConfigWithFile("config.json")
```

2. 组装请求参数
- 为了接口使用更加方便, 我们将参数粗分为必填/非必填, 必填直接放在结构体内, 非必填放在结构体的map字段ExtendInfos中
```
// 1.商户配置初始化
config, _ := NewConfigWithFile("config.json")
// 2.组装请求参数
reqParam := &V1TradePaymentJspayRequest{
	OutOrderID: NewOrderID(),
	GoodsDesc:  "测试商品",
	TradeType:  "A_NATIVE", // 支付宝正扫
	Amount:     1,
	FromAddrIp: "127.0.0.1",
}
	
```

3. 发起API调用
```
resp, err := config.V1TradePaymentJspayRequest(reqParam)
if err != nil {
	fmt.Println(resp, err)
	return
}
```

## 关于RSA密钥对生成
- 生成密钥对
```
openssl genpkey -algorithm RSA -out rsa_private_key.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
```
- 或者使用tlsdk.CreateRSAKeyPair()生成密钥对

