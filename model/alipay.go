package model

const (
	K_ALI_PAY_SANDBOX_API_URL    = "https://openapi.alipaydev.com/gateway.do"
	K_ALI_PAY_PRODUCTION_API_URL = "https://openapi.alipay.com/gateway.do"
)

type AliPay struct {
	appId      string
	apiDomain  string
	publickKey []byte
	privateKey []byte
}

func New(appId string, publicKey, privateKey []byte, isProduction bool) (client *AliPay) {
	client = &AliPay{
		appId:      appId,
		privateKey: privateKey,
		publickKey: publicKey,
	}
	if isProduction {
		client.apiDomain = K_ALI_PAY_PRODUCTION_API_URL
	} else {
		client.apiDomain = K_ALI_PAY_SANDBOX_API_URL
	}
	return client
}

