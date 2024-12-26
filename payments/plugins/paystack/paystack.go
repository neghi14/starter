package paystack

import "github.com/neghi14/starter/payments"

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

func New(config *PaystackConfig) (*payments.PaymentAdapter, error) {

	return &payments.PaymentAdapter{
		Name: "paystack-payment",
	}, nil
}
