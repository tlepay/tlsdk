package tlsdk

import (
	"bytes"
	"fmt"
	"time"

	"github.com/codingeasygo/util/converter"
	"github.com/codingeasygo/util/xhttp"
	"github.com/codingeasygo/util/xmap"
)

func (c *Config) V1TradePaymentJspayRequest(reqParam *V1TradePaymentJspayRequest) (resp CommonResp, err error) {
	reqUrl := c.APIURL + V1_TRADE_PAYMENT_JSPAY
	return c.PostJSON(reqUrl, reqParam)
}

type CommonReq struct {
	MchID     string `json:"mch_id" valid:"mch_id,r|s,l:0;"`       // 商户号
	APPID     string `json:"app_id" valid:"app_id,r|s,l:0;"`       // APPID
	Timestamp int64  `json:"timestamp" valid:"timestamp,r|i,r:0;"` // 时间戳
	Sign      string `json:"sign" valid:"sign,r|s,l:0;"`           // 加签结果
	Data      string `json:"data" valid:"data,r|s,l:0;"`           // json数据
}

type CommonResp struct {
	Code      int    `json:"code" valid:"code,r|i,r:0;"`           // 返回码
	MchID     string `json:"mch_id" valid:"mch_id,r|s,l:0;"`       // 商户号
	APPID     string `json:"app_id" valid:"app_id,r|s,l:0;"`       // APPID
	Timestamp int64  `json:"timestamp" valid:"timestamp,r|i,r:0;"` // 时间戳
	Sign      string `json:"sign" valid:"sign,r|s,l:0;"`           // 加签结果
	Data      string `json:"data" valid:"data,r|s,l:0;"`           // json数据
	DataM     xmap.M `json:"data_m" valid:"data_m,o|o;"`           // data转map
}

func (c CommonResp) VerifySign(pubKey string) (err error) {
	return RSAVerify([]byte(pubKey), []byte(c.Data), c.Sign)
}

func (c CommonReq) CalcSign(priKey string) (sign string, err error) {
	return RSASign([]byte(priKey), []byte(c.Data))
}

var DEBUG = false

func (c *Config) PostJSON(reqUrl string, reqParam interface{}) (resp CommonResp, err error) {
	commonReq := CommonReq{
		MchID:     c.MchID,
		APPID:     c.APPID,
		Timestamp: time.Now().UnixMilli(),
		Data:      converter.JSON(reqParam),
	}
	sign, err := commonReq.CalcSign(c.RSAMCHPrivateKey)
	if err != nil {
		return
	}
	commonReq.Sign = sign
	if DEBUG {
		fmt.Printf("PostJSON reqUrl:%s, reqParam:%s \n", reqUrl, converter.JSON(commonReq))
	}
	err = xhttp.PostJSON(&resp, bytes.NewBuffer([]byte(converter.JSON(commonReq))), reqUrl)
	if err != nil {
		return
	}
	err = resp.VerifySign(c.RSATLPublicKey)
	if err != nil {
		return
	}
	_, err = converter.UnmarshalJSON(bytes.NewBuffer([]byte(resp.Data)), &resp.DataM)
	return
}
