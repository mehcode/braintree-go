package braintree

type PaymentMethodNonce struct {
	Type        string `xml:"type,omitempty"`
	Nonce       string `xml:"nonce,omitempty"`
	Description string `xml:"description,omitempty"`
	Consumed    bool   `xml:"consumed,omitempty"`
	Details     struct {
		Last2    string `xml:"last-two,omit-empty"`
		CardType string `xml:"card-type,omit-empty"`
	} `xml:"details"`
}
