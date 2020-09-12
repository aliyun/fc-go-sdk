package fc

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestCustomDomainStructs(t *testing.T) {
	suite.Run(t, new(CustomDomainStructsTestSuite))
}

type CustomDomainStructsTestSuite struct {
	suite.Suite
}

func (s *CustomDomainStructsTestSuite) TestCreateCustomDomain() {
	assert := s.Require()

	input := NewCreateCustomDomainInput()
	assert.NotNil(input)

	input.WithDomainName("my-app.com")
	assert.NotNil(input.DomainName)
	assert.Equal("my-app.com", *input.DomainName)

	input.WithProtocol("HTTP")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTP", *input.Protocol)

	input.WithProtocol("HTTPS")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTPS", *input.Protocol)

	input.WithProtocol("HTTP,HTTPS")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTP,HTTPS", *input.Protocol)

	cert := CertConfig{}
	cert.WithCertName("cert-name")
	cert.WithPrivateKey("server-privateKey")
	cert.WithCertificate("server-certificate")
	assert.Equal("server-privateKey", *cert.PrivateKey)
	assert.Equal("server-certificate", *cert.Certificate)

	input.WithCertConfig(&cert)
	assert.Equal("cert-name", *input.CertConfig.CertName)
	assert.Equal("server-privateKey", *input.CertConfig.PrivateKey)
	assert.Equal("server-certificate", *input.CertConfig.Certificate)

	input.WithRouteConfig(&RouteConfig{})
	assert.NotNil(input.RouteConfig)

	pathConfig := NewPathConfig()
	assert.NotNil(pathConfig)

	pathConfig.WithPath("/login")
	assert.NotNil(pathConfig.Path)
	assert.Equal("/login", *pathConfig.Path)

	pathConfig.WithServiceName("service0")
	assert.NotNil(pathConfig.ServiceName)
	assert.Equal("service0", *pathConfig.ServiceName)

	pathConfig.WithFunctionName("function0")
	assert.NotNil(pathConfig.FunctionName)
	assert.Equal("function0", *pathConfig.FunctionName)

	pathConfig.WithQualifier("v1")
	assert.NotNil(pathConfig.Qualifier)
	assert.Equal("v1", *pathConfig.Qualifier)

	pathConfig.WithMethods([]string{"GET", "POST"})
	assert.NotNil(pathConfig.Methods)
	assert.Equal([]string{"GET", "POST"}, pathConfig.Methods)

	routeConfig := NewRouteConfig()
	assert.NotNil(routeConfig)

	routeConfig.WithRoutes([]PathConfig{*pathConfig})
	assert.NotNil(routeConfig.Routes)
	assert.Equal(1, len(routeConfig.Routes))
}

func (s *CustomDomainStructsTestSuite) TestUpdateCustomDomain() {
	assert := s.Require()

	input := NewUpdateCustomDomainInput("my-app.com")
	assert.NotNil(input)
	assert.NotNil(input.DomainName)
	assert.Equal("my-app.com", *input.DomainName)

	input.WithProtocol("HTTP")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTP", *input.Protocol)

	input.WithProtocol("HTTPS")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTPS", *input.Protocol)

	input.WithProtocol("HTTP,HTTPS")
	assert.NotNil(input.Protocol)
	assert.Equal("HTTP,HTTPS", *input.Protocol)

	cert := CertConfig{}
	cert.WithCertName("cert-name")
	cert.WithPrivateKey("server-privateKey")
	cert.WithCertificate("server-certificate")
	assert.Equal("cert-name", *cert.CertName)
	assert.Equal("server-privateKey", *cert.PrivateKey)
	assert.Equal("server-certificate", *cert.Certificate)

	input.WithCertConfig(&cert)
	assert.Equal("cert-name", *cert.CertName)
	assert.Equal("server-privateKey", *input.CertConfig.PrivateKey)
	assert.Equal("server-certificate", *input.CertConfig.Certificate)

	input.WithRouteConfig(&RouteConfig{})
	assert.NotNil(input.RouteConfig)

	pathConfig := NewPathConfig()
	assert.NotNil(pathConfig)

	pathConfig.WithPath("/login")
	assert.NotNil(pathConfig.Path)
	assert.Equal("/login", *pathConfig.Path)

	pathConfig.WithServiceName("service0")
	assert.NotNil(pathConfig.ServiceName)
	assert.Equal("service0", *pathConfig.ServiceName)

	pathConfig.WithFunctionName("function0")
	assert.NotNil(pathConfig.FunctionName)
	assert.Equal("function0", *pathConfig.FunctionName)

	pathConfig.WithQualifier("v1")
	assert.NotNil(pathConfig.Qualifier)
	assert.Equal("v1", *pathConfig.Qualifier)

	pathConfig.WithMethods([]string{"GET", "POST"})
	assert.NotNil(pathConfig.Methods)
	assert.Equal([]string{"GET", "POST"}, pathConfig.Methods)

	routeConfig := NewRouteConfig()
	assert.NotNil(routeConfig)

	routeConfig.WithRoutes([]PathConfig{*pathConfig})
	assert.NotNil(routeConfig.Routes)
	assert.Equal(1, len(routeConfig.Routes))
}

func (s *CustomDomainStructsTestSuite) TestGetCustomDomain() {
	assert := s.Require()
	input := NewGetCustomDomainInput("my-app.com")
	assert.NotNil(input)
	assert.NotNil(input.DomainName)
	assert.Equal("my-app.com", *input.DomainName)
}

func (s *CustomDomainStructsTestSuite) TestDeleteCustomDomain() {
	assert := s.Require()
	input := NewDeleteCustomDomainInput("my-app.com")
	assert.NotNil(input)
	assert.NotNil(input.DomainName)
	assert.Equal("my-app.com", *input.DomainName)
}

func (s *CustomDomainStructsTestSuite) TestListCustomDomain() {
	assert := s.Require()
	input := NewListCustomDomainsInput()
	assert.NotNil(input)

	input.WithPrefix("app")
	assert.NotNil(input.Prefix)
	assert.Equal("app", *input.Prefix)

	input.WithNextToken("your-app.com")
	assert.NotNil(input.NextToken)
	assert.Equal("your-app.com", *input.NextToken)

	input.WithLimit(int32(10))
	assert.NotNil(input.Limit)
	assert.Equal(int32(10), *input.Limit)

	input.WithStartKey("my")
	assert.NotNil(input.StartKey)
	assert.Equal("my", *input.StartKey)
}
