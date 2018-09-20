package model

import (
	"encoding/json"
	"fmt"
)

type AliPayParam interface {
	APIName() string
	Params() map[string]string
	ExtJSONParamName() string
	ExtJSONParamValue() string
}

//AliPayTradeWapPay https://docs.open.alipay.com/203/107090
type AliPayTradeWapPay struct {
	NotifyURL string `json:"-"`

	//biz content
	Subject     string `json:"subject"`
	OutTradeNo  string `json:"out_trade_no"`
	TotalAmount string `json:"total_amount"`
	ProductCode string `json:"product_code"`

	Body               string `json:"body,omitempty"`
	TimeoutExpress     string `json:"timeout_express,omitempty"`
	SellerId           string `json:"seller_id,omitempty"`
	AuthToken          string `json:"auth_token,omitempty"`
	GoodsType          string `json:"goods_type,omitempty"`
	PassbackParams     string `json:"passback_params,omitempty"`
	PromoParams        string `json:"promo_params,omitempty"`
	ExtendParams       string `json:"extend_params,omitempty"`
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`
	DisablePayChannels string `json:"disable_pay_channels,omitempty"`
	StoreId            string `json:"store_id,omitempty"`
}

func (this AliPayTradeWapPay) APIName() string {
	return "alipay.trade.wap.pay"
}

func (this AliPayTradeWapPay) Params() map[string]string {
	var m = make(map[string]string)
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeWapPay) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeWapPay) ExtJSONParamValue() string {
	bytes, err := json.Marshal(this)
	if err != nil {
		fmt.Errorf("JSON marshal err", err)
		return ""
	}
	return string(bytes)
}

func marshal(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)

}
