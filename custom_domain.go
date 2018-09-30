package fc

import (
	"fmt"
	"net/url"
	"strconv"
)

type PathConfig struct {
	Path         string `json:"path"`
	ServiceName  string `json:"serviceName"`
	FunctionName string `json:"functionName"`
	Qualifier    string `json:"qualifier"`
}

func NewPathConfig() *PathConfig {
	return &PathConfig{}
}

func (c *PathConfig) WithPath(path string) *PathConfig {
	c.Path = path
	return c
}

func (c *PathConfig) WithServiceName(serviceName string) *PathConfig {
	c.ServiceName = serviceName
	return c
}

func (c *PathConfig) WithFunctionName(functionName string) *PathConfig {
	c.FunctionName = functionName
	return c
}

func (c *PathConfig) WithQualifier(qualifier string) *PathConfig {
	c.Qualifier = qualifier
	return c
}

type RouteConfig struct {
	Routes []PathConfig `json:"routes"`
}

type customDomainMetadata struct {
	DomainName       *string      `json:"domainName"`
	AccountId        *string      `json:"accountId"`
	Protocol         *string      `json:"protocol"`
	ApiVersion       *string      `json:"apiVersion"`
	RouteConfig      *RouteConfig `json:"routeConfig"`
	CreatedTime      *string      `json:"createdTime"`
	LastModifiedTime *string      `json:"lastModifiedTime"`
}

type CreateCustomDomainInput struct {
	DomainName  *string      `json:"domainName"`
	Protocol    *string      `json:"protocol"`
	RouteConfig *RouteConfig `json:"routeConfig"`
}

func (i *CreateCustomDomainInput) WithDomainName(domainName string) *CreateCustomDomainInput {
	i.DomainName = &domainName
	return i
}

func (i *CreateCustomDomainInput) WithProtocol(protocol string) *CreateCustomDomainInput {
	i.Protocol = &protocol
	return i
}

func (i *CreateCustomDomainInput) WithRouteConfig(routeConfig RouteConfig) *CreateCustomDomainInput {
	i.RouteConfig = &routeConfig
	return i
}

//noinspection GoUnusedExportedFunction
func NewCreateCustomDomainInput() *CreateCustomDomainInput {
	return &CreateCustomDomainInput{}
}

func (i CreateCustomDomainInput) GetQueryParams() url.Values {
	return url.Values{}
}

func (i CreateCustomDomainInput) GetPath() string {
	return customDomainPath
}

func (i CreateCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i CreateCustomDomainInput) GetPayload() interface{} {
	return i
}

func (i CreateCustomDomainInput) Validate() error {
	return nil
}

type CreateCustomDomainOutput struct {
	customDomainMetadata
	outputDecorator
}

type CustomDomainUpdateObject struct {
	Protocol    *string      `json:"protocol"`
	RouteConfig *RouteConfig `json:"routeConfig"`
}

type UpdateCustomDomainInput struct {
	DomainName *string
	CustomDomainUpdateObject
}

//noinspection GoUnusedExportedFunction
func NewUpdateCustomDomainInput(domainName string) *UpdateCustomDomainInput {
	return &UpdateCustomDomainInput{
		DomainName: &domainName,
	}
}

func (i *UpdateCustomDomainInput) WithProtocol(protocol string) *UpdateCustomDomainInput {
	i.Protocol = &protocol
	return i
}

func (i *UpdateCustomDomainInput) WithRouteConfig(routeConfig RouteConfig) *UpdateCustomDomainInput {
	i.RouteConfig = &routeConfig
	return i
}

func (i *UpdateCustomDomainInput) GetQueryParams() url.Values {
	return url.Values{}
}

func (i *UpdateCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleCustomDomainPath, pathEscape(*i.DomainName))
}

func (i *UpdateCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *UpdateCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *UpdateCustomDomainInput) Validate() error {
	if IsBlank(i.DomainName) {
		return fmt.Errorf("domain name is required but not provided")
	}
	return nil
}

type UpdateCustomDomainOutput struct {
	customDomainMetadata
	outputDecorator
}

type GetCustomDomainInput struct {
	DomainName *string
}

func NewGetCustomDomainInput(domainName string) *GetCustomDomainInput {
	return &GetCustomDomainInput{
		DomainName: &domainName,
	}
}

func (i *GetCustomDomainInput) GetQueryParams() url.Values {
	return url.Values{}
}

func (i *GetCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleCustomDomainPath, pathEscape(*i.DomainName))
}

func (i *GetCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *GetCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *GetCustomDomainInput) Validate() error {
	if IsBlank(i.DomainName) {
		return fmt.Errorf("domain name is required but not provided")
	}
	return nil
}

type GetCustomDomainOutput struct {
	customDomainMetadata
	outputDecorator
}

type ListCustomDomainInput struct {
	Query
}

func NewListCustomDomainInput() *ListCustomDomainInput {
	return &ListCustomDomainInput{}
}

func (i *ListCustomDomainInput) WithPrefix(prefix string) *ListCustomDomainInput {
	i.Prefix = &prefix
	return i
}

func (i *ListCustomDomainInput) WithStartKey(startKey string) *ListCustomDomainInput {
	i.StartKey = &startKey
	return i
}

func (i *ListCustomDomainInput) WithNextToken(nextToken string) *ListCustomDomainInput {
	i.NextToken = &nextToken
	return i
}

func (i *ListCustomDomainInput) WithLimit(limit int32) *ListCustomDomainInput {
	i.Limit = &limit
	return i
}

func (i *ListCustomDomainInput) GetQueryParams() url.Values {
	out := url.Values{}
	if i.Prefix != nil {
		out.Set("prefix", *i.Prefix)
	}

	if i.StartKey != nil {
		out.Set("startKey", *i.StartKey)
	}

	if i.NextToken != nil {
		out.Set("nextToken", *i.NextToken)
	}

	if i.Limit != nil {
		out.Set("limit", strconv.FormatInt(int64(*i.Limit), 10))
	}

	return out
}

func (i *ListCustomDomainInput) GetPath() string {
	return fmt.Sprintf(customDomainPath)
}

func (i *ListCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *ListCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *ListCustomDomainInput) Validate() error {
	return nil
}

type ListCustomDomainOutput struct {
	CustomDomains []customDomainMetadata `json:"customDomains"`
	NextToken     *string                `json:"nextToken,omitempty"`
	outputDecorator
}

type DeleteCustomDomainInput struct {
	DomainName *string
}

func NewDeleteCustomDomainInput(domainName string) *DeleteCustomDomainInput {
	return &DeleteCustomDomainInput{
		DomainName: &domainName,
	}
}

func (i *DeleteCustomDomainInput) GetQueryParams() url.Values {
	return url.Values{}
}

func (i *DeleteCustomDomainInput) GetPath() string {
	return fmt.Sprintf(singleCustomDomainPath, pathEscape(*i.DomainName))
}

func (i *DeleteCustomDomainInput) GetHeaders() Header {
	return make(Header, 0)
}

func (i *DeleteCustomDomainInput) GetPayload() interface{} {
	return nil
}

func (i *DeleteCustomDomainInput) Validate() error {
	if IsBlank(i.DomainName) {
		return fmt.Errorf("domain name is required but not provided")
	}
	return nil
}

type DeleteCustomDomainOutput struct {
	outputDecorator
}
