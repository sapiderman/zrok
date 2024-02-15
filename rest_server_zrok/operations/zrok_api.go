// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/openziti/zrok/rest_model_zrok"
	"github.com/openziti/zrok/rest_server_zrok/operations/account"
	"github.com/openziti/zrok/rest_server_zrok/operations/admin"
	"github.com/openziti/zrok/rest_server_zrok/operations/environment"
	"github.com/openziti/zrok/rest_server_zrok/operations/metadata"
	"github.com/openziti/zrok/rest_server_zrok/operations/share"
)

// NewZrokAPI creates a new Zrok instance
func NewZrokAPI(spec *loads.Document) *ZrokAPI {
	return &ZrokAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ShareAccessHandler: share.AccessHandlerFunc(func(params share.AccessParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation share.Access has not yet been implemented")
		}),
		MetadataConfigurationHandler: metadata.ConfigurationHandlerFunc(func(params metadata.ConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.Configuration has not yet been implemented")
		}),
		AdminCreateFrontendHandler: admin.CreateFrontendHandlerFunc(func(params admin.CreateFrontendParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.CreateFrontend has not yet been implemented")
		}),
		AdminCreateIdentityHandler: admin.CreateIdentityHandlerFunc(func(params admin.CreateIdentityParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.CreateIdentity has not yet been implemented")
		}),
		AdminDeleteFrontendHandler: admin.DeleteFrontendHandlerFunc(func(params admin.DeleteFrontendParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.DeleteFrontend has not yet been implemented")
		}),
		EnvironmentDisableHandler: environment.DisableHandlerFunc(func(params environment.DisableParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation environment.Disable has not yet been implemented")
		}),
		EnvironmentEnableHandler: environment.EnableHandlerFunc(func(params environment.EnableParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation environment.Enable has not yet been implemented")
		}),
		MetadataGetAccountDetailHandler: metadata.GetAccountDetailHandlerFunc(func(params metadata.GetAccountDetailParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetAccountDetail has not yet been implemented")
		}),
		MetadataGetAccountMetricsHandler: metadata.GetAccountMetricsHandlerFunc(func(params metadata.GetAccountMetricsParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetAccountMetrics has not yet been implemented")
		}),
		MetadataGetEnvironmentDetailHandler: metadata.GetEnvironmentDetailHandlerFunc(func(params metadata.GetEnvironmentDetailParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetEnvironmentDetail has not yet been implemented")
		}),
		MetadataGetEnvironmentMetricsHandler: metadata.GetEnvironmentMetricsHandlerFunc(func(params metadata.GetEnvironmentMetricsParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetEnvironmentMetrics has not yet been implemented")
		}),
		MetadataGetFrontendDetailHandler: metadata.GetFrontendDetailHandlerFunc(func(params metadata.GetFrontendDetailParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetFrontendDetail has not yet been implemented")
		}),
		MetadataGetShareDetailHandler: metadata.GetShareDetailHandlerFunc(func(params metadata.GetShareDetailParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetShareDetail has not yet been implemented")
		}),
		MetadataGetShareMetricsHandler: metadata.GetShareMetricsHandlerFunc(func(params metadata.GetShareMetricsParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.GetShareMetrics has not yet been implemented")
		}),
		AccountInviteHandler: account.InviteHandlerFunc(func(params account.InviteParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Invite has not yet been implemented")
		}),
		AdminInviteTokenGenerateHandler: admin.InviteTokenGenerateHandlerFunc(func(params admin.InviteTokenGenerateParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.InviteTokenGenerate has not yet been implemented")
		}),
		AdminListFrontendsHandler: admin.ListFrontendsHandlerFunc(func(params admin.ListFrontendsParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.ListFrontends has not yet been implemented")
		}),
		AccountLoginHandler: account.LoginHandlerFunc(func(params account.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Login has not yet been implemented")
		}),
		MetadataOverviewHandler: metadata.OverviewHandlerFunc(func(params metadata.OverviewParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation metadata.Overview has not yet been implemented")
		}),
		AccountRegisterHandler: account.RegisterHandlerFunc(func(params account.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Register has not yet been implemented")
		}),
		AccountResetPasswordHandler: account.ResetPasswordHandlerFunc(func(params account.ResetPasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation account.ResetPassword has not yet been implemented")
		}),
		AccountResetPasswordRequestHandler: account.ResetPasswordRequestHandlerFunc(func(params account.ResetPasswordRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation account.ResetPasswordRequest has not yet been implemented")
		}),
		AccountResetTokenHandler: account.ResetTokenHandlerFunc(func(params account.ResetTokenParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ResetToken has not yet been implemented")
		}),
		ShareShareHandler: share.ShareHandlerFunc(func(params share.ShareParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation share.Share has not yet been implemented")
		}),
		ShareUnaccessHandler: share.UnaccessHandlerFunc(func(params share.UnaccessParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation share.Unaccess has not yet been implemented")
		}),
		ShareUnshareHandler: share.UnshareHandlerFunc(func(params share.UnshareParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation share.Unshare has not yet been implemented")
		}),
		AdminUpdateFrontendHandler: admin.UpdateFrontendHandlerFunc(func(params admin.UpdateFrontendParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation admin.UpdateFrontend has not yet been implemented")
		}),
		ShareUpdateShareHandler: share.UpdateShareHandlerFunc(func(params share.UpdateShareParams, principal *rest_model_zrok.Principal) middleware.Responder {
			return middleware.NotImplemented("operation share.UpdateShare has not yet been implemented")
		}),
		AccountVerifyHandler: account.VerifyHandlerFunc(func(params account.VerifyParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Verify has not yet been implemented")
		}),
		MetadataVersionHandler: metadata.VersionHandlerFunc(func(params metadata.VersionParams) middleware.Responder {
			return middleware.NotImplemented("operation metadata.Version has not yet been implemented")
		}),

		// Applies when the "x-token" header is set
		KeyAuth: func(token string) (*rest_model_zrok.Principal, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ZrokAPI zrok client access */
type ZrokAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/zrok.v1+json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/zrok.v1+json
	JSONProducer runtime.Producer

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x-token provided in the header
	KeyAuth func(string) (*rest_model_zrok.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ShareAccessHandler sets the operation handler for the access operation
	ShareAccessHandler share.AccessHandler
	// MetadataConfigurationHandler sets the operation handler for the configuration operation
	MetadataConfigurationHandler metadata.ConfigurationHandler
	// AdminCreateFrontendHandler sets the operation handler for the create frontend operation
	AdminCreateFrontendHandler admin.CreateFrontendHandler
	// AdminCreateIdentityHandler sets the operation handler for the create identity operation
	AdminCreateIdentityHandler admin.CreateIdentityHandler
	// AdminDeleteFrontendHandler sets the operation handler for the delete frontend operation
	AdminDeleteFrontendHandler admin.DeleteFrontendHandler
	// EnvironmentDisableHandler sets the operation handler for the disable operation
	EnvironmentDisableHandler environment.DisableHandler
	// EnvironmentEnableHandler sets the operation handler for the enable operation
	EnvironmentEnableHandler environment.EnableHandler
	// MetadataGetAccountDetailHandler sets the operation handler for the get account detail operation
	MetadataGetAccountDetailHandler metadata.GetAccountDetailHandler
	// MetadataGetAccountMetricsHandler sets the operation handler for the get account metrics operation
	MetadataGetAccountMetricsHandler metadata.GetAccountMetricsHandler
	// MetadataGetEnvironmentDetailHandler sets the operation handler for the get environment detail operation
	MetadataGetEnvironmentDetailHandler metadata.GetEnvironmentDetailHandler
	// MetadataGetEnvironmentMetricsHandler sets the operation handler for the get environment metrics operation
	MetadataGetEnvironmentMetricsHandler metadata.GetEnvironmentMetricsHandler
	// MetadataGetFrontendDetailHandler sets the operation handler for the get frontend detail operation
	MetadataGetFrontendDetailHandler metadata.GetFrontendDetailHandler
	// MetadataGetShareDetailHandler sets the operation handler for the get share detail operation
	MetadataGetShareDetailHandler metadata.GetShareDetailHandler
	// MetadataGetShareMetricsHandler sets the operation handler for the get share metrics operation
	MetadataGetShareMetricsHandler metadata.GetShareMetricsHandler
	// AccountInviteHandler sets the operation handler for the invite operation
	AccountInviteHandler account.InviteHandler
	// AdminInviteTokenGenerateHandler sets the operation handler for the invite token generate operation
	AdminInviteTokenGenerateHandler admin.InviteTokenGenerateHandler
	// AdminListFrontendsHandler sets the operation handler for the list frontends operation
	AdminListFrontendsHandler admin.ListFrontendsHandler
	// AccountLoginHandler sets the operation handler for the login operation
	AccountLoginHandler account.LoginHandler
	// MetadataOverviewHandler sets the operation handler for the overview operation
	MetadataOverviewHandler metadata.OverviewHandler
	// AccountRegisterHandler sets the operation handler for the register operation
	AccountRegisterHandler account.RegisterHandler
	// AccountResetPasswordHandler sets the operation handler for the reset password operation
	AccountResetPasswordHandler account.ResetPasswordHandler
	// AccountResetPasswordRequestHandler sets the operation handler for the reset password request operation
	AccountResetPasswordRequestHandler account.ResetPasswordRequestHandler
	// AccountResetTokenHandler sets the operation handler for the reset token operation
	AccountResetTokenHandler account.ResetTokenHandler
	// ShareShareHandler sets the operation handler for the share operation
	ShareShareHandler share.ShareHandler
	// ShareUnaccessHandler sets the operation handler for the unaccess operation
	ShareUnaccessHandler share.UnaccessHandler
	// ShareUnshareHandler sets the operation handler for the unshare operation
	ShareUnshareHandler share.UnshareHandler
	// AdminUpdateFrontendHandler sets the operation handler for the update frontend operation
	AdminUpdateFrontendHandler admin.UpdateFrontendHandler
	// ShareUpdateShareHandler sets the operation handler for the update share operation
	ShareUpdateShareHandler share.UpdateShareHandler
	// AccountVerifyHandler sets the operation handler for the verify operation
	AccountVerifyHandler account.VerifyHandler
	// MetadataVersionHandler sets the operation handler for the version operation
	MetadataVersionHandler metadata.VersionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ZrokAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ZrokAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ZrokAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ZrokAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ZrokAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ZrokAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ZrokAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ZrokAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ZrokAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ZrokAPI
func (o *ZrokAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "XTokenAuth")
	}

	if o.ShareAccessHandler == nil {
		unregistered = append(unregistered, "share.AccessHandler")
	}
	if o.MetadataConfigurationHandler == nil {
		unregistered = append(unregistered, "metadata.ConfigurationHandler")
	}
	if o.AdminCreateFrontendHandler == nil {
		unregistered = append(unregistered, "admin.CreateFrontendHandler")
	}
	if o.AdminCreateIdentityHandler == nil {
		unregistered = append(unregistered, "admin.CreateIdentityHandler")
	}
	if o.AdminDeleteFrontendHandler == nil {
		unregistered = append(unregistered, "admin.DeleteFrontendHandler")
	}
	if o.EnvironmentDisableHandler == nil {
		unregistered = append(unregistered, "environment.DisableHandler")
	}
	if o.EnvironmentEnableHandler == nil {
		unregistered = append(unregistered, "environment.EnableHandler")
	}
	if o.MetadataGetAccountDetailHandler == nil {
		unregistered = append(unregistered, "metadata.GetAccountDetailHandler")
	}
	if o.MetadataGetAccountMetricsHandler == nil {
		unregistered = append(unregistered, "metadata.GetAccountMetricsHandler")
	}
	if o.MetadataGetEnvironmentDetailHandler == nil {
		unregistered = append(unregistered, "metadata.GetEnvironmentDetailHandler")
	}
	if o.MetadataGetEnvironmentMetricsHandler == nil {
		unregistered = append(unregistered, "metadata.GetEnvironmentMetricsHandler")
	}
	if o.MetadataGetFrontendDetailHandler == nil {
		unregistered = append(unregistered, "metadata.GetFrontendDetailHandler")
	}
	if o.MetadataGetShareDetailHandler == nil {
		unregistered = append(unregistered, "metadata.GetShareDetailHandler")
	}
	if o.MetadataGetShareMetricsHandler == nil {
		unregistered = append(unregistered, "metadata.GetShareMetricsHandler")
	}
	if o.AccountInviteHandler == nil {
		unregistered = append(unregistered, "account.InviteHandler")
	}
	if o.AdminInviteTokenGenerateHandler == nil {
		unregistered = append(unregistered, "admin.InviteTokenGenerateHandler")
	}
	if o.AdminListFrontendsHandler == nil {
		unregistered = append(unregistered, "admin.ListFrontendsHandler")
	}
	if o.AccountLoginHandler == nil {
		unregistered = append(unregistered, "account.LoginHandler")
	}
	if o.MetadataOverviewHandler == nil {
		unregistered = append(unregistered, "metadata.OverviewHandler")
	}
	if o.AccountRegisterHandler == nil {
		unregistered = append(unregistered, "account.RegisterHandler")
	}
	if o.AccountResetPasswordHandler == nil {
		unregistered = append(unregistered, "account.ResetPasswordHandler")
	}
	if o.AccountResetPasswordRequestHandler == nil {
		unregistered = append(unregistered, "account.ResetPasswordRequestHandler")
	}
	if o.AccountResetTokenHandler == nil {
		unregistered = append(unregistered, "account.ResetTokenHandler")
	}
	if o.ShareShareHandler == nil {
		unregistered = append(unregistered, "share.ShareHandler")
	}
	if o.ShareUnaccessHandler == nil {
		unregistered = append(unregistered, "share.UnaccessHandler")
	}
	if o.ShareUnshareHandler == nil {
		unregistered = append(unregistered, "share.UnshareHandler")
	}
	if o.AdminUpdateFrontendHandler == nil {
		unregistered = append(unregistered, "admin.UpdateFrontendHandler")
	}
	if o.ShareUpdateShareHandler == nil {
		unregistered = append(unregistered, "share.UpdateShareHandler")
	}
	if o.AccountVerifyHandler == nil {
		unregistered = append(unregistered, "account.VerifyHandler")
	}
	if o.MetadataVersionHandler == nil {
		unregistered = append(unregistered, "metadata.VersionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ZrokAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ZrokAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.KeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *ZrokAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZrokAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/zrok.v1+json":
			result["application/zrok.v1+json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZrokAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/zrok.v1+json":
			result["application/zrok.v1+json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ZrokAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the zrok API
func (o *ZrokAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ZrokAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/access"] = share.NewAccess(o.context, o.ShareAccessHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/configuration"] = metadata.NewConfiguration(o.context, o.MetadataConfigurationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/frontend"] = admin.NewCreateFrontend(o.context, o.AdminCreateFrontendHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/identity"] = admin.NewCreateIdentity(o.context, o.AdminCreateIdentityHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/frontend"] = admin.NewDeleteFrontend(o.context, o.AdminDeleteFrontendHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/disable"] = environment.NewDisable(o.context, o.EnvironmentDisableHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/enable"] = environment.NewEnable(o.context, o.EnvironmentEnableHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/detail/account"] = metadata.NewGetAccountDetail(o.context, o.MetadataGetAccountDetailHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics/account"] = metadata.NewGetAccountMetrics(o.context, o.MetadataGetAccountMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/detail/environment/{envZId}"] = metadata.NewGetEnvironmentDetail(o.context, o.MetadataGetEnvironmentDetailHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics/environment/{envId}"] = metadata.NewGetEnvironmentMetrics(o.context, o.MetadataGetEnvironmentMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/detail/frontend/{feId}"] = metadata.NewGetFrontendDetail(o.context, o.MetadataGetFrontendDetailHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/detail/share/{shrToken}"] = metadata.NewGetShareDetail(o.context, o.MetadataGetShareDetailHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics/share/{shrToken}"] = metadata.NewGetShareMetrics(o.context, o.MetadataGetShareMetricsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/invite"] = account.NewInvite(o.context, o.AccountInviteHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/invite/token/generate"] = admin.NewInviteTokenGenerate(o.context, o.AdminInviteTokenGenerateHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/frontends"] = admin.NewListFrontends(o.context, o.AdminListFrontendsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = account.NewLogin(o.context, o.AccountLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/overview"] = metadata.NewOverview(o.context, o.MetadataOverviewHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/register"] = account.NewRegister(o.context, o.AccountRegisterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resetPassword"] = account.NewResetPassword(o.context, o.AccountResetPasswordHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resetPasswordRequest"] = account.NewResetPasswordRequest(o.context, o.AccountResetPasswordRequestHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resetToken"] = account.NewResetToken(o.context, o.AccountResetTokenHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/share"] = share.NewShare(o.context, o.ShareShareHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/unaccess"] = share.NewUnaccess(o.context, o.ShareUnaccessHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/unshare"] = share.NewUnshare(o.context, o.ShareUnshareHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/frontend"] = admin.NewUpdateFrontend(o.context, o.AdminUpdateFrontendHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/share"] = share.NewUpdateShare(o.context, o.ShareUpdateShareHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/verify"] = account.NewVerify(o.context, o.AccountVerifyHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/version"] = metadata.NewVersion(o.context, o.MetadataVersionHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ZrokAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ZrokAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ZrokAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ZrokAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ZrokAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
