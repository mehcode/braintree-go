package braintree

type PaymentMethodNonceGateway struct {
	*Braintree
}

// Find returns the plan with the specified id, or nil
func (g *PaymentMethodNonceGateway) Find(token string) (*PaymentMethodNonce, error) {
	resp, err := g.execute("GET", "payment_method_nonces/"+token, nil)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case 200:
		return resp.paymentMethodNonce()
	}
	return nil, &invalidResponseError{resp}
}
