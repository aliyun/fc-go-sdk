package fc

import (
	"net/url"
)

const (
	accountPath = "/account-settings"
)

type accountSettings struct {
	AvailableAZs []string `json:"availableAZs"`
}

// GetAccountSettingsInput defines get account settings input.
type GetAccountSettingsInput struct {
}

//noinspection GoUnusedExportedFunction
func NewGetAccountSettingsInput() *GetAccountSettingsInput {
	return new(GetAccountSettingsInput)
}

func (o GetAccountSettingsInput) GetQueryParams() url.Values {
	out := url.Values{}
	return out
}

func (o GetAccountSettingsInput) GetPath() string {
	return accountPath
}

func (o GetAccountSettingsInput) GetHeaders() Header {
	return make(Header, 0)
}

func (o GetAccountSettingsInput) GetPayload() interface{} {
	return o
}

func (o GetAccountSettingsInput) Validate() error {
	return nil
}

// GetAccountSettingsOutput defines get account settings output.
type GetAccountSettingsOutput struct {
	accountSettings
	outputDecorator
}
