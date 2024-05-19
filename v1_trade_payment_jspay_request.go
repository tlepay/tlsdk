package tlsdk

type V1TradePaymentJspayRequest struct {
	// 商户订单号
	OutOrderID string `json:"out_order_id" valid:"out_order_id,r|s,l:0;"`
	// 商品描述
	GoodsDesc string `json:"goods_desc" valid:"goods_desc,r|s,l:0;"`
	// 交易类型
	TradeType string `json:"trade_type" valid:"trade_type,r|s,l:0;"`
	// 金额，单位分
	Amount float64 `json:"amount" valid:"amount,r|i,r:0;"`
	// 买家备注
	BuyerMemo string `json:"buyer_memo" valid:"buyer_memo,o|s,l:0;"`
	// 交易有效期
	ExpiredTime int64 `json:"expired_time" valid:"expired_time,o|i,r:0;"`
	// 支付成功回调地址
	NotifyURL string `json:"notify_url" valid:"notify_url,o|s,l:0;"`
	// 用户ip
	FromAddrIp string `json:"from_addr_ip" valid:"from_addr_ip,r|s,l:0;"`
	// 微信openid
	OpenID string `json:"openid" valid:"openid,o|s,l:0;"`
}
