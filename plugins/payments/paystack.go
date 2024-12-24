package payments

import "github.com/neghi14/starter"

type PaystackConfig struct {
	key string
}

func NewPaystackConfig() *PaystackConfig {
	opts := &PaystackConfig{}
	return opts
}

func (p *PaystackConfig) SetKey(key string) *PaystackConfig {
	p.key = key
	return p
}

func Paystack(config *PaystackConfig) (*starter.PaymentAdapter, error) {

	return &starter.PaymentAdapter{
		Name: "paystack-payment",
	}, nil
}
