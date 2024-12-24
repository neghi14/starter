package starter

type PaymentAdapter struct {
	Name   string
	Refund func()
	Pay    func()
}

type PaymentAdapterOptions struct{}

type PaymentConfig func(*PaymentAdapterOptions)
