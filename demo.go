package tlsdk

import (
	"fmt"

	"github.com/codingeasygo/util/converter"
)

func V1TradePaymentJspayRequestDemo() {
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
	// 3. 发起API调用
	resp, err := config.V1TradePaymentJspayRequest(reqParam)
	if err != nil {
		fmt.Println(resp, err)
		return
	}
	// 3.处理返回结果
	fmt.Printf("V1TradePaymentJspayRequestDemo response:\n %v\n", converter.JSON(resp.DataM))
}
