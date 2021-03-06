package paypalsdk

import "fmt"

// CreateSinglePayout submits a payout with an asynchronous API call, which immediately returns the results of a PayPal payment.
// For email payout set RecipientType: "EMAIL" and receiver email into Receiver
// Endpoint: POST /v1/payments/payouts
func (c *Client) CreateSinglePayout(p Payout) (*PayoutResponse, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/payments/payouts"), p)
	if err != nil {
		return &PayoutResponse{}, err
	}

	response := &PayoutResponse{}

	err = c.SendWithAuth(req, response)
	if err != nil {
		return response, err
	}

	return response, nil
}
