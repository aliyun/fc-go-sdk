package fc

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-resty/resty"
)

// Client defines fc client
type Client struct {
	Config  *Config
	Connect *Connection
}

// NewClient new fc client
func NewClient(endpoint, apiVersion, accessKeyID, accessKeySecret string, opts ...ClientOption) (*Client, error) {
	config := NewConfig()
	config.APIVersion = apiVersion
	config.AccessKeyID = accessKeyID
	config.AccessKeySecret = accessKeySecret
	config.Endpoint, config.host = GetAccessPoint(endpoint)
	connect := NewConnection()
	client := &Client{config, connect}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// GetAccountSettings returns account settings from fc
func (c *Client) GetAccountSettings(input *GetAccountSettingsInput) (*GetAccountSettingsOutput, error) {
	var output = new(GetAccountSettingsOutput)
	if input == nil {
		input = new(GetAccountSettingsInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetService returns service metadata from fc
func (c *Client) GetService(input *GetServiceInput) (*GetServiceOutput, error) {
	var output = new(GetServiceOutput)
	if input == nil {
		input = new(GetServiceInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// ListServices returns list of services from fc
func (c *Client) ListServices(input *ListServicesInput) (*ListServicesOutput, error) {
	var output = new(ListServicesOutput)
	if input == nil {
		input = new(ListServicesInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// UpdateService updates service
func (c *Client) UpdateService(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	var output = new(UpdateServiceOutput)
	if input == nil {
		input = new(UpdateServiceInput)
	}

	err := c.sendRequestHelper(input, http.MethodPut, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// CreateService creates service
func (c *Client) CreateService(input *CreateServiceInput) (*CreateServiceOutput, error) {
	var output = new(CreateServiceOutput)
	if input == nil {
		input = new(CreateServiceInput)
	}

	err := c.sendRequestHelper(input, http.MethodPost, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// DeleteService deletes service
func (c *Client) DeleteService(input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	var output = new(DeleteServiceOutput)
	if input == nil {
		input = new(DeleteServiceInput)
	}

	err := c.sendRequestHelper(input, http.MethodDelete, output, *deleteRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) CreateCustomDomain(input *CreateCustomDomainInput) (*CreateCustomDomainOutput, error) {
	var output = new(CreateCustomDomainOutput)
	if input == nil {
		input = new(CreateCustomDomainInput)
	}

	err := c.sendRequestHelper(input, http.MethodPost, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) GetCustomDomain(input *GetCustomDomainInput) (*GetCustomDomainOutput, error) {
	var output = new(GetCustomDomainOutput)
	if input == nil {
		input = new(GetCustomDomainInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}
func (c *Client) UpdateCustomDomain(input *UpdateCustomDomainInput) (*UpdateCustomDomainOutput, error) {
	var output = new(UpdateCustomDomainOutput)
	if input == nil {
		input = new(UpdateCustomDomainInput)
	}

	err := c.sendRequestHelper(input, http.MethodPut, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil

}

func (c *Client) ListCustomDomain(input *ListCustomDomainInput) (*ListCustomDomainOutput, error) {
	var output = new(ListCustomDomainOutput)
	if input == nil {
		input = new(ListCustomDomainInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (c *Client) DeleteCustomDomain(input *DeleteCustomDomainInput) (*DeleteFunctionOutput, error) {
	var output = new(DeleteFunctionOutput)
	if input == nil {
		input = new(DeleteCustomDomainInput)
	}

	err := c.sendRequestHelper(input, http.MethodDelete, output, *deleteRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// CreateFunction creates function
func (c *Client) CreateFunction(input *CreateFunctionInput) (*CreateFunctionOutput, error) {
	var output = new(CreateFunctionOutput)
	if input == nil {
		input = new(CreateFunctionInput)
	}

	err := c.sendRequestHelper(input, http.MethodPost, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// DeleteFunction deletes function from service
func (c *Client) DeleteFunction(input *DeleteFunctionInput) (*DeleteFunctionOutput, error) {
	var output = new(DeleteFunctionOutput)
	if input == nil {
		input = new(DeleteFunctionInput)
	}

	err := c.sendRequestHelper(input, http.MethodDelete, output, *deleteRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetFunction returns function metadata from service
func (c *Client) GetFunction(input *GetFunctionInput) (*GetFunctionOutput, error) {
	var output = new(GetFunctionOutput)
	if input == nil {
		input = new(GetFunctionInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetFunctionCode returns function code
func (c *Client) GetFunctionCode(input *GetFunctionCodeInput) (*GetFunctionCodeOutput, error) {
	var output = new(GetFunctionCodeOutput)
	if input == nil {
		input = new(GetFunctionCodeInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// ListFunctions returns list of functions
func (c *Client) ListFunctions(input *ListFunctionsInput) (*ListFunctionsOutput, error) {
	var output = new(ListFunctionsOutput)
	if input == nil {
		input = new(ListFunctionsInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// UpdateFunction updates function
func (c *Client) UpdateFunction(input *UpdateFunctionInput) (*UpdateFunctionOutput, error) {
	var output = new(UpdateFunctionOutput)
	if input == nil {
		input = new(UpdateFunctionInput)
	}

	err := c.sendRequestHelper(input, http.MethodPut, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// CreateTrigger creates trigger
func (c *Client) CreateTrigger(input *CreateTriggerInput) (*CreateTriggerOutput, error) {
	var output = new(CreateTriggerOutput)
	if input == nil {
		input = new(CreateTriggerInput)
	}

	err := c.sendRequestHelper(input, http.MethodPost, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GetTrigger returns trigger metadata
func (c *Client) GetTrigger(input *GetTriggerInput) (*GetTriggerOutput, error) {
	var output = new(GetTriggerOutput)
	if input == nil {
		input = new(GetTriggerInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// UpdateTrigger updates trigger
func (c *Client) UpdateTrigger(input *UpdateTriggerInput) (*UpdateTriggerOutput, error) {
	var output = new(UpdateTriggerOutput)
	if input == nil {
		input = new(UpdateTriggerInput)
	}

	err := c.sendRequestHelper(input, http.MethodPut, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// DeleteTrigger deletes trigger
func (c *Client) DeleteTrigger(input *DeleteTriggerInput) (*DeleteTriggerOutput, error) {
	var output = new(DeleteTriggerOutput)
	if input == nil {
		input = new(DeleteTriggerInput)
	}

	err := c.sendRequestHelper(input, http.MethodDelete, output, *deleteRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// ListTriggers returns list of triggers
func (c *Client) ListTriggers(input *ListTriggersInput) (*ListTriggersOutput, error) {
	var output = new(ListTriggersOutput)
	if input == nil {
		input = new(ListTriggersInput)
	}

	err := c.sendRequestHelper(input, http.MethodGet, output, *defaultRequestOption)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// InvokeFunction : invoke function in fc
func (c *Client) InvokeFunction(input *InvokeFunctionInput) (*InvokeFunctionOutput, error) {
	if input == nil {
		input = new(InvokeFunctionInput)
	}

	var output = new(InvokeFunctionOutput)
	httpResponse, err := c.sendRequest(input, http.MethodPost)
	if err != nil {
		return nil, err
	}
	output.Header = httpResponse.Header()
	output.Payload = httpResponse.Body()

	return output, nil
}

type requestOption struct {
	DoSetHeader     bool
	DoUnmarshalBody bool
}

var (
	defaultRequestOption = &requestOption{
		DoSetHeader:     true,
		DoUnmarshalBody: true,
	}
	deleteRequestOption = &requestOption{
		DoSetHeader:     true,
		DoUnmarshalBody: false,
	}
)

func (c *Client) sendRequestHelper(
	input ServiceInput,
	httpMethod string,
	output ServiceOutput,
	option requestOption,
) error {
	httpResponse, err := c.sendRequest(input, httpMethod)
	if err != nil {
		return err
	}

	if option.DoSetHeader {
		output.SetHeader(httpResponse.Header())
	}

	if option.DoUnmarshalBody {
		json.Unmarshal(httpResponse.Body(), output)
	}

	return nil
}

func (c *Client) sendRequest(input ServiceInput, httpMethod string) (*resty.Response, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	var serviceError = new(ServiceError)
	path := "/" + c.Config.APIVersion + input.GetPath()

	headerParams := make(Header)
	for k, v := range input.GetHeaders() {
		headerParams[k] = v
	}
	headerParams["Host"] = c.Config.host
	headerParams[HTTPHeaderAccountID] = c.Config.AccountID
	headerParams[HTTPHeaderUserAgent] = c.Config.UserAgent
	headerParams["Accept"] = "application/json"
	// Caution: should not declare this as byte[] whose zero value is an empty byte array
	// if input has no payload, the http body should not be populated at all.
	var rawBody interface{}
	if input.GetPayload() != nil {
		switch input.GetPayload().(type) {
		case *[]byte:
			headerParams["Content-Type"] = "application/octet-stream"
			b := input.GetPayload().(*[]byte)
			headerParams["Content-MD5"] = MD5(*b)
			rawBody = *b
		default:
			headerParams["Content-Type"] = "application/json"
			b, err := json.Marshal(input.GetPayload())
			if err != nil {
				// TODO: return client side error
				return nil, nil
			}
			headerParams["Content-MD5"] = MD5(b)
			rawBody = b
		}
	}
	headerParams["Date"] = time.Now().UTC().Format(http.TimeFormat)
	if c.Config.SecurityToken != "" {
		headerParams[HTTPHeaderSecurityToken] = c.Config.SecurityToken
	}
	headerParams["Authorization"] = GetAuthStr(
		c.Config.AccessKeyID,
		c.Config.AccessKeySecret,
		httpMethod,
		headerParams,
		path,
	)
	resp, err := c.Connect.SendRequest(
		c.Config.Endpoint+path,
		httpMethod,
		rawBody,
		headerParams,
		input.GetQueryParams(),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() >= 300 {
		serviceError.RequestID = resp.Header().Get(HTTPHeaderRequestID)
		serviceError.HTTPStatus = resp.StatusCode()
		json.Unmarshal(resp.Body(), &serviceError)
		return nil, serviceError
	}
	return resp, nil
}
