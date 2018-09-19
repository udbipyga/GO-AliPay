package model

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
