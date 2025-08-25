package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GenerateBackendAPIModelsResponse represents the GenerateBackendAPIModelsResponse schema from the OpenAPI specification
type GenerateBackendAPIModelsResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
}

// CreateBackendAuthRequestResourceConfig represents the CreateBackendAuthRequestResourceConfig schema from the OpenAPI specification
type CreateBackendAuthRequestResourceConfig struct {
	Service interface{} `json:"Service"`
	Userpoolconfigs CreateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs"`
	Authresources interface{} `json:"AuthResources"`
	Identitypoolconfigs CreateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
}

// BackendAPIResourceConfig represents the BackendAPIResourceConfig schema from the OpenAPI specification
type BackendAPIResourceConfig struct {
	Additionalauthtypes interface{} `json:"AdditionalAuthTypes,omitempty"`
	Apiname interface{} `json:"ApiName,omitempty"`
	Conflictresolution CreateBackendAPIrequestresourceConfigConflictResolution `json:"ConflictResolution,omitempty"`
	Defaultauthtype CreateBackendAPIrequestresourceConfigDefaultAuthType `json:"DefaultAuthType,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Transformschema interface{} `json:"TransformSchema,omitempty"`
}

// UpdateBackendAuthResponse represents the UpdateBackendAuthResponse schema from the OpenAPI specification
type UpdateBackendAuthResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// UpdateBackendAuthResourceConfig represents the UpdateBackendAuthResourceConfig schema from the OpenAPI specification
type UpdateBackendAuthResourceConfig struct {
	Service interface{} `json:"Service"`
	Userpoolconfigs UpdateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs"`
	Authresources interface{} `json:"AuthResources"`
	Identitypoolconfigs UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
}

// GetBackendrequest represents the GetBackendrequest schema from the OpenAPI specification
type GetBackendrequest struct {
	Backendenvironmentname string `json:"backendEnvironmentName,omitempty"` // The name of the backend environment.
}

// UpdateBackendAuthRequest represents the UpdateBackendAuthRequest schema from the OpenAPI specification
type UpdateBackendAuthRequest struct {
	Resourceconfig UpdateBackendAuthRequestResourceConfig `json:"ResourceConfig"`
	Resourcename interface{} `json:"ResourceName"`
}

// UpdateBackendAPIRequest represents the UpdateBackendAPIRequest schema from the OpenAPI specification
type UpdateBackendAPIRequest struct {
	Resourcename interface{} `json:"ResourceName"`
	Resourceconfig DeleteBackendAPIRequestResourceConfig `json:"ResourceConfig,omitempty"`
}

// CreateBackendConfigResponse represents the CreateBackendConfigResponse schema from the OpenAPI specification
type CreateBackendConfigResponse struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
}

// RemoveAllBackendsResponse represents the RemoveAllBackendsResponse schema from the OpenAPI specification
type RemoveAllBackendsResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetBackendAuthResponseResourceConfig represents the GetBackendAuthResponseResourceConfig schema from the OpenAPI specification
type GetBackendAuthResponseResourceConfig struct {
	Service interface{} `json:"Service"`
	Userpoolconfigs CreateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs"`
	Authresources interface{} `json:"AuthResources"`
	Identitypoolconfigs CreateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
}

// ListBackendJobsResponse represents the ListBackendJobsResponse schema from the OpenAPI specification
type ListBackendJobsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Jobs interface{} `json:"Jobs,omitempty"`
}

// CreateBackendStorageResourceConfig represents the CreateBackendStorageResourceConfig schema from the OpenAPI specification
type CreateBackendStorageResourceConfig struct {
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions"`
	Servicename interface{} `json:"ServiceName"`
	Bucketname interface{} `json:"BucketName,omitempty"`
}

// DeleteTokenRequest represents the DeleteTokenRequest schema from the OpenAPI specification
type DeleteTokenRequest struct {
}

// GetBackendStorageResponseResourceConfig represents the GetBackendStorageResponseResourceConfig schema from the OpenAPI specification
type GetBackendStorageResponseResourceConfig struct {
	Imported interface{} `json:"Imported"`
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions,omitempty"`
	Servicename interface{} `json:"ServiceName"`
	Bucketname interface{} `json:"BucketName,omitempty"`
}

// GetBackendResponse represents the GetBackendResponse schema from the OpenAPI specification
type GetBackendResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Appname interface{} `json:"AppName,omitempty"`
	Backendenvironmentlist interface{} `json:"BackendEnvironmentList,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Amplifyfeatureflags interface{} `json:"AmplifyFeatureFlags,omitempty"`
	Amplifymetaconfig interface{} `json:"AmplifyMetaConfig,omitempty"`
}

// RemoveAllBackendsrequest represents the RemoveAllBackendsrequest schema from the OpenAPI specification
type RemoveAllBackendsrequest struct {
	Cleanamplifyapp bool `json:"cleanAmplifyApp,omitempty"` // Cleans up the Amplify Console app if this value is set to true.
}

// GetBackendAPIResponseResourceConfig represents the GetBackendAPIResponseResourceConfig schema from the OpenAPI specification
type GetBackendAPIResponseResourceConfig struct {
	Additionalauthtypes interface{} `json:"AdditionalAuthTypes,omitempty"`
	Apiname interface{} `json:"ApiName,omitempty"`
	Conflictresolution CreateBackendAPIrequestresourceConfigConflictResolution `json:"ConflictResolution,omitempty"`
	Defaultauthtype CreateBackendAPIrequestresourceConfigDefaultAuthType `json:"DefaultAuthType,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Transformschema interface{} `json:"TransformSchema,omitempty"`
}

// CreateBackendAuthMFAConfigSettings represents the CreateBackendAuthMFAConfigSettings schema from the OpenAPI specification
type CreateBackendAuthMFAConfigSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
	Mfatypes interface{} `json:"MfaTypes,omitempty"`
}

// CreateBackendStorageRequestResourceConfig represents the CreateBackendStorageRequestResourceConfig schema from the OpenAPI specification
type CreateBackendStorageRequestResourceConfig struct {
	Bucketname interface{} `json:"BucketName,omitempty"`
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions"`
	Servicename interface{} `json:"ServiceName"`
}

// CreateBackendAuthForgotPasswordConfigSmsSettings represents the CreateBackendAuthForgotPasswordConfigSmsSettings schema from the OpenAPI specification
type CreateBackendAuthForgotPasswordConfigSmsSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
}

// CreateBackendStorageResponse represents the CreateBackendStorageResponse schema from the OpenAPI specification
type CreateBackendStorageResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
}

// CreateTokenResponse represents the CreateTokenResponse schema from the OpenAPI specification
type CreateTokenResponse struct {
	Ttl interface{} `json:"Ttl,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Challengecode interface{} `json:"ChallengeCode,omitempty"`
	Sessionid interface{} `json:"SessionId,omitempty"`
}

// GetBackendAPIModelsResponse represents the GetBackendAPIModelsResponse schema from the OpenAPI specification
type GetBackendAPIModelsResponse struct {
	Models interface{} `json:"Models,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Modelintrospectionschema interface{} `json:"ModelIntrospectionSchema,omitempty"`
}

// CloneBackendrequest represents the CloneBackendrequest schema from the OpenAPI specification
type CloneBackendrequest struct {
	Targetenvironmentname string `json:"targetEnvironmentName"` // The name of the destination backend environment to be created.
}

// GetTokenResponse represents the GetTokenResponse schema from the OpenAPI specification
type GetTokenResponse struct {
	Ttl interface{} `json:"Ttl,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Challengecode interface{} `json:"ChallengeCode,omitempty"`
	Sessionid interface{} `json:"SessionId,omitempty"`
}

// CreateBackendStoragerequestresourceConfig represents the CreateBackendStoragerequestresourceConfig schema from the OpenAPI specification
type CreateBackendStoragerequestresourceConfig struct {
	Servicename interface{} `json:"ServiceName,omitempty"`
	Bucketname interface{} `json:"BucketName,omitempty"`
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions,omitempty"`
}

// ListBackendJobsrequest represents the ListBackendJobsrequest schema from the OpenAPI specification
type ListBackendJobsrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // The token for the next set of results.
	Operation string `json:"operation,omitempty"` // Filters the list of response objects to include only those with the specified operation name.
	Status string `json:"status,omitempty"` // Filters the list of response objects to include only those with the specified status.
	Jobid string `json:"jobId,omitempty"` // The ID for the job.
	Maxresults int `json:"maxResults,omitempty"` // The maximum number of results that you want in the response.
}

// CreateBackendAuthVerificationMessageConfig represents the CreateBackendAuthVerificationMessageConfig schema from the OpenAPI specification
type CreateBackendAuthVerificationMessageConfig struct {
	Emailsettings CreateBackendAuthVerificationMessageConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthVerificationMessageConfigSmsSettings `json:"SmsSettings,omitempty"`
	Deliverymethod interface{} `json:"DeliveryMethod"`
}

// UpdateBackendAuthUserPoolConfigForgotPassword represents the UpdateBackendAuthUserPoolConfigForgotPassword schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfigForgotPassword struct {
	Smssettings UpdateBackendAuthForgotPasswordConfigSmsSettings `json:"SmsSettings,omitempty"`
	Deliverymethod interface{} `json:"DeliveryMethod,omitempty"`
	Emailsettings CreateBackendAuthForgotPasswordConfigEmailSettings `json:"EmailSettings,omitempty"`
}

// CreateBackendAPIrequestresourceConfigDefaultAuthType represents the CreateBackendAPIrequestresourceConfigDefaultAuthType schema from the OpenAPI specification
type CreateBackendAPIrequestresourceConfigDefaultAuthType struct {
	Mode interface{} `json:"Mode,omitempty"`
	Settings BackendAPIAuthTypeSettings `json:"Settings,omitempty"`
}

// CreateBackendStoragerequest represents the CreateBackendStoragerequest schema from the OpenAPI specification
type CreateBackendStoragerequest struct {
	Backendenvironmentname string `json:"backendEnvironmentName"` // The name of the backend environment.
	Resourceconfig CreateBackendStoragerequestresourceConfig `json:"resourceConfig"` // The resource configuration for creating backend storage.
	Resourcename string `json:"resourceName"` // The name of the storage resource.
}

// UpdateBackendConfigResponseLoginAuthConfig represents the UpdateBackendConfigResponseLoginAuthConfig schema from the OpenAPI specification
type UpdateBackendConfigResponseLoginAuthConfig struct {
	Awsuserpoolsid interface{} `json:"AwsUserPoolsId,omitempty"`
	Awsuserpoolswebclientid interface{} `json:"AwsUserPoolsWebClientId,omitempty"`
	Awscognitoidentitypoolid interface{} `json:"AwsCognitoIdentityPoolId,omitempty"`
	Awscognitoregion interface{} `json:"AwsCognitoRegion,omitempty"`
}

// UpdateBackendStorageRequest represents the UpdateBackendStorageRequest schema from the OpenAPI specification
type UpdateBackendStorageRequest struct {
	Resourceconfig UpdateBackendStorageRequestResourceConfig `json:"ResourceConfig"`
	Resourcename interface{} `json:"ResourceName"`
}

// UpdateBackendConfigRequest represents the UpdateBackendConfigRequest schema from the OpenAPI specification
type UpdateBackendConfigRequest struct {
	Loginauthconfig UpdateBackendConfigRequestLoginAuthConfig `json:"LoginAuthConfig,omitempty"`
}

// CreateBackendStorageRequest represents the CreateBackendStorageRequest schema from the OpenAPI specification
type CreateBackendStorageRequest struct {
	Resourcename interface{} `json:"ResourceName"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName"`
	Resourceconfig CreateBackendStorageRequestResourceConfig `json:"ResourceConfig"`
}

// Settings represents the Settings schema from the OpenAPI specification
type Settings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
	Mfatypes interface{} `json:"MfaTypes,omitempty"`
}

// BackendAuthSocialProviderConfig represents the BackendAuthSocialProviderConfig schema from the OpenAPI specification
type BackendAuthSocialProviderConfig struct {
	Clientid interface{} `json:"ClientId,omitempty"`
	Clientsecret interface{} `json:"ClientSecret,omitempty"`
}

// UpdateBackendStorageRequestResourceConfig represents the UpdateBackendStorageRequestResourceConfig schema from the OpenAPI specification
type UpdateBackendStorageRequestResourceConfig struct {
	Servicename interface{} `json:"ServiceName"`
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions"`
}

// CreateBackendAuthrequestresourceConfigIdentityPoolConfigs represents the CreateBackendAuthrequestresourceConfigIdentityPoolConfigs schema from the OpenAPI specification
type CreateBackendAuthrequestresourceConfigIdentityPoolConfigs struct {
	Unauthenticatedlogin interface{} `json:"UnauthenticatedLogin"`
	Identitypoolname interface{} `json:"IdentityPoolName"`
}

// UpdateBackendAuthForgotPasswordConfig represents the UpdateBackendAuthForgotPasswordConfig schema from the OpenAPI specification
type UpdateBackendAuthForgotPasswordConfig struct {
	Deliverymethod interface{} `json:"DeliveryMethod,omitempty"`
	Emailsettings CreateBackendAuthForgotPasswordConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings UpdateBackendAuthForgotPasswordConfigSmsSettings `json:"SmsSettings,omitempty"`
}

// UpdateBackendAuthrequest represents the UpdateBackendAuthrequest schema from the OpenAPI specification
type UpdateBackendAuthrequest struct {
	Resourceconfig UpdateBackendAuthrequestresourceConfig `json:"resourceConfig"` // Defines the resource configuration when updating an authentication resource in your Amplify project.
	Resourcename string `json:"resourceName"` // The name of this resource.
}

// UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs represents the UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs schema from the OpenAPI specification
type UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs struct {
	Unauthenticatedlogin interface{} `json:"UnauthenticatedLogin,omitempty"`
}

// ListS3BucketsResponse represents the ListS3BucketsResponse schema from the OpenAPI specification
type ListS3BucketsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Buckets interface{} `json:"Buckets,omitempty"`
}

// UpdateBackendStoragerequestresourceConfig represents the UpdateBackendStoragerequestresourceConfig schema from the OpenAPI specification
type UpdateBackendStoragerequestresourceConfig struct {
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions,omitempty"`
	Servicename interface{} `json:"ServiceName,omitempty"`
}

// CreateBackendAuthVerificationMessageConfigSmsSettings represents the CreateBackendAuthVerificationMessageConfigSmsSettings schema from the OpenAPI specification
type CreateBackendAuthVerificationMessageConfigSmsSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
}

// CreateBackendAuthUserPoolConfig represents the CreateBackendAuthUserPoolConfig schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfig struct {
	Requiredsignupattributes interface{} `json:"RequiredSignUpAttributes"`
	Signinmethod interface{} `json:"SignInMethod"`
	Userpoolname interface{} `json:"UserPoolName"`
	Verificationmessage CreateBackendAuthUserPoolConfigVerificationMessage `json:"VerificationMessage,omitempty"`
	Forgotpassword CreateBackendAuthUserPoolConfigForgotPassword `json:"ForgotPassword,omitempty"`
	Mfa CreateBackendAuthUserPoolConfigMfa `json:"Mfa,omitempty"`
	Oauth CreateBackendAuthUserPoolConfigOAuth `json:"OAuth,omitempty"`
	Passwordpolicy CreateBackendAuthUserPoolConfigPasswordPolicy `json:"PasswordPolicy,omitempty"`
}

// UpdateBackendAPIResponse represents the UpdateBackendAPIResponse schema from the OpenAPI specification
type UpdateBackendAPIResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UpdateBackendStorageResourceConfig represents the UpdateBackendStorageResourceConfig schema from the OpenAPI specification
type UpdateBackendStorageResourceConfig struct {
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions"`
	Servicename interface{} `json:"ServiceName"`
}

// ListS3BucketsRequest represents the ListS3BucketsRequest schema from the OpenAPI specification
type ListS3BucketsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteBackendAPIRequestResourceConfig represents the DeleteBackendAPIRequestResourceConfig schema from the OpenAPI specification
type DeleteBackendAPIRequestResourceConfig struct {
	Conflictresolution CreateBackendAPIrequestresourceConfigConflictResolution `json:"ConflictResolution,omitempty"`
	Defaultauthtype CreateBackendAPIrequestresourceConfigDefaultAuthType `json:"DefaultAuthType,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Transformschema interface{} `json:"TransformSchema,omitempty"`
	Additionalauthtypes interface{} `json:"AdditionalAuthTypes,omitempty"`
	Apiname interface{} `json:"ApiName,omitempty"`
}

// ImportBackendAuthrequest represents the ImportBackendAuthrequest schema from the OpenAPI specification
type ImportBackendAuthrequest struct {
	Identitypoolid string `json:"identityPoolId,omitempty"` // The ID of the Amazon Cognito identity pool.
	Nativeclientid string `json:"nativeClientId"` // The ID of the Amazon Cognito native client.
	Userpoolid string `json:"userPoolId"` // The ID of the Amazon Cognito user pool.
	Webclientid string `json:"webClientId"` // The ID of the Amazon Cognito web client.
}

// CreateBackendAuthUserPoolConfigOAuth represents the CreateBackendAuthUserPoolConfigOAuth schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfigOAuth struct {
	Redirectsigninuris interface{} `json:"RedirectSignInURIs"`
	Redirectsignouturis interface{} `json:"RedirectSignOutURIs"`
	Socialprovidersettings CreateBackendAuthOAuthConfigSocialProviderSettings `json:"SocialProviderSettings,omitempty"`
	Domainprefix interface{} `json:"DomainPrefix,omitempty"`
	Oauthgranttype interface{} `json:"OAuthGrantType"`
	Oauthscopes interface{} `json:"OAuthScopes"`
}

// UpdateBackendJobResponse represents the UpdateBackendJobResponse schema from the OpenAPI specification
type UpdateBackendJobResponse struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatetime interface{} `json:"UpdateTime,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// CreateBackendStoragerequestresourceConfigPermissions represents the CreateBackendStoragerequestresourceConfigPermissions schema from the OpenAPI specification
type CreateBackendStoragerequestresourceConfigPermissions struct {
	Authenticated interface{} `json:"Authenticated"`
	Unauthenticated interface{} `json:"UnAuthenticated,omitempty"`
}

// CreateBackendAuthIdentityPoolConfig represents the CreateBackendAuthIdentityPoolConfig schema from the OpenAPI specification
type CreateBackendAuthIdentityPoolConfig struct {
	Identitypoolname interface{} `json:"IdentityPoolName"`
	Unauthenticatedlogin interface{} `json:"UnauthenticatedLogin"`
}

// RemoveBackendConfigRequest represents the RemoveBackendConfigRequest schema from the OpenAPI specification
type RemoveBackendConfigRequest struct {
}

// CreateBackendAuthOAuthConfig represents the CreateBackendAuthOAuthConfig schema from the OpenAPI specification
type CreateBackendAuthOAuthConfig struct {
	Domainprefix interface{} `json:"DomainPrefix,omitempty"`
	Oauthgranttype interface{} `json:"OAuthGrantType"`
	Oauthscopes interface{} `json:"OAuthScopes"`
	Redirectsigninuris interface{} `json:"RedirectSignInURIs"`
	Redirectsignouturis interface{} `json:"RedirectSignOutURIs"`
	Socialprovidersettings CreateBackendAuthOAuthConfigSocialProviderSettings `json:"SocialProviderSettings,omitempty"`
}

// BackendAPIAppSyncAuthSettings represents the BackendAPIAppSyncAuthSettings schema from the OpenAPI specification
type BackendAPIAppSyncAuthSettings struct {
	Openidissueurl interface{} `json:"OpenIDIssueURL,omitempty"`
	Openidprovidername interface{} `json:"OpenIDProviderName,omitempty"`
	Cognitouserpoolid interface{} `json:"CognitoUserPoolId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Expirationtime interface{} `json:"ExpirationTime,omitempty"`
	Openidauthttl interface{} `json:"OpenIDAuthTTL,omitempty"`
	Openidclientid interface{} `json:"OpenIDClientId,omitempty"`
	Openidiatttl interface{} `json:"OpenIDIatTTL,omitempty"`
}

// BackendAPIAuthType represents the BackendAPIAuthType schema from the OpenAPI specification
type BackendAPIAuthType struct {
	Mode interface{} `json:"Mode,omitempty"`
	Settings BackendAPIAuthTypeSettings `json:"Settings,omitempty"`
}

// BackendJobRespObj represents the BackendJobRespObj schema from the OpenAPI specification
type BackendJobRespObj struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatetime interface{} `json:"UpdateTime,omitempty"`
	Appid interface{} `json:"AppId"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// SmsSettings represents the SmsSettings schema from the OpenAPI specification
type SmsSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
}

// CreateTokenRequest represents the CreateTokenRequest schema from the OpenAPI specification
type CreateTokenRequest struct {
}

// UpdateBackendAuthPasswordPolicyConfig represents the UpdateBackendAuthPasswordPolicyConfig schema from the OpenAPI specification
type UpdateBackendAuthPasswordPolicyConfig struct {
	Minimumlength interface{} `json:"MinimumLength,omitempty"`
	Additionalconstraints interface{} `json:"AdditionalConstraints,omitempty"`
}

// DeleteBackendStoragerequest represents the DeleteBackendStoragerequest schema from the OpenAPI specification
type DeleteBackendStoragerequest struct {
	Resourcename string `json:"resourceName"` // The name of the storage resource.
	Servicename string `json:"serviceName"` // The name of the storage service.
}

// UpdateBackendStoragerequest represents the UpdateBackendStoragerequest schema from the OpenAPI specification
type UpdateBackendStoragerequest struct {
	Resourceconfig UpdateBackendStoragerequestresourceConfig `json:"resourceConfig"` // The resource configuration for updating backend storage.
	Resourcename string `json:"resourceName"` // The name of the storage resource.
}

// CreateBackendAuthResourceConfig represents the CreateBackendAuthResourceConfig schema from the OpenAPI specification
type CreateBackendAuthResourceConfig struct {
	Authresources interface{} `json:"AuthResources"`
	Identitypoolconfigs CreateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
	Service interface{} `json:"Service"`
	Userpoolconfigs CreateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs"`
}

// CreateBackendAPIrequest represents the CreateBackendAPIrequest schema from the OpenAPI specification
type CreateBackendAPIrequest struct {
	Backendenvironmentname string `json:"backendEnvironmentName"` // The name of the backend environment.
	Resourceconfig CreateBackendAPIrequestresourceConfig `json:"resourceConfig"` // The resource config for the data model, configured as a part of the Amplify project.
	Resourcename string `json:"resourceName"` // The name of this resource.
}

// DeleteBackendStorageResponse represents the DeleteBackendStorageResponse schema from the OpenAPI specification
type DeleteBackendStorageResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UpdateBackendStorageResponse represents the UpdateBackendStorageResponse schema from the OpenAPI specification
type UpdateBackendStorageResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// GetBackendJobResponse represents the GetBackendJobResponse schema from the OpenAPI specification
type GetBackendJobResponse struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Updatetime interface{} `json:"UpdateTime,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// UpdateBackendConfigrequestloginAuthConfig represents the UpdateBackendConfigrequestloginAuthConfig schema from the OpenAPI specification
type UpdateBackendConfigrequestloginAuthConfig struct {
	Awscognitoidentitypoolid interface{} `json:"AwsCognitoIdentityPoolId,omitempty"`
	Awscognitoregion interface{} `json:"AwsCognitoRegion,omitempty"`
	Awsuserpoolsid interface{} `json:"AwsUserPoolsId,omitempty"`
	Awsuserpoolswebclientid interface{} `json:"AwsUserPoolsWebClientId,omitempty"`
}

// DeleteBackendAPIRequest represents the DeleteBackendAPIRequest schema from the OpenAPI specification
type DeleteBackendAPIRequest struct {
	Resourcename interface{} `json:"ResourceName"`
	Resourceconfig DeleteBackendAPIRequestResourceConfig `json:"ResourceConfig,omitempty"`
}

// UpdateBackendAuthOAuthConfigSocialProviderSettings represents the UpdateBackendAuthOAuthConfigSocialProviderSettings schema from the OpenAPI specification
type UpdateBackendAuthOAuthConfigSocialProviderSettings struct {
	Loginwithamazon BackendAuthSocialProviderConfig `json:"LoginWithAmazon,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Signinwithapple BackendAuthAppleProviderConfig `json:"SignInWithApple,omitempty"` // Describes Apple social federation configurations for allowing your app users to sign in using OAuth.
	Facebook BackendAuthSocialProviderConfig `json:"Facebook,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Google BackendAuthSocialProviderConfig `json:"Google,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
}

// GetBackendAPIResponse represents the GetBackendAPIResponse schema from the OpenAPI specification
type GetBackendAPIResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Resourceconfig GetBackendAPIResponseResourceConfig `json:"ResourceConfig,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
}

// CreateBackendAuthOAuthConfigSocialProviderSettings represents the CreateBackendAuthOAuthConfigSocialProviderSettings schema from the OpenAPI specification
type CreateBackendAuthOAuthConfigSocialProviderSettings struct {
	Loginwithamazon BackendAuthSocialProviderConfig `json:"LoginWithAmazon,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Signinwithapple BackendAuthAppleProviderConfig `json:"SignInWithApple,omitempty"` // Describes Apple social federation configurations for allowing your app users to sign in using OAuth.
	Facebook BackendAuthSocialProviderConfig `json:"Facebook,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Google BackendAuthSocialProviderConfig `json:"Google,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
}

// CreateBackendResponse represents the CreateBackendResponse schema from the OpenAPI specification
type CreateBackendResponse struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
}

// GetBackendStorageRequest represents the GetBackendStorageRequest schema from the OpenAPI specification
type GetBackendStorageRequest struct {
	Resourcename interface{} `json:"ResourceName"`
}

// S3BucketInfo represents the S3BucketInfo schema from the OpenAPI specification
type S3BucketInfo struct {
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// CreateBackendConfigrequest represents the CreateBackendConfigrequest schema from the OpenAPI specification
type CreateBackendConfigrequest struct {
	Backendmanagerappid string `json:"backendManagerAppId,omitempty"` // The app ID for the backend manager.
}

// BackendAPIConflictResolution represents the BackendAPIConflictResolution schema from the OpenAPI specification
type BackendAPIConflictResolution struct {
	Resolutionstrategy interface{} `json:"ResolutionStrategy,omitempty"`
}

// CreateBackendAPIRequest represents the CreateBackendAPIRequest schema from the OpenAPI specification
type CreateBackendAPIRequest struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName"`
	Resourceconfig CreateBackendAPIRequestResourceConfig `json:"ResourceConfig"`
	Resourcename interface{} `json:"ResourceName"`
}

// UpdateBackendJobrequest represents the UpdateBackendJobrequest schema from the OpenAPI specification
type UpdateBackendJobrequest struct {
	Status string `json:"status,omitempty"` // Filters the list of response objects to include only those with the specified status.
	Operation string `json:"operation,omitempty"` // Filters the list of response objects to include only those with the specified operation name.
}

// UpdateBackendAuthUserPoolConfig represents the UpdateBackendAuthUserPoolConfig schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfig struct {
	Forgotpassword UpdateBackendAuthUserPoolConfigForgotPassword `json:"ForgotPassword,omitempty"`
	Mfa UpdateBackendAuthUserPoolConfigMfa `json:"Mfa,omitempty"`
	Oauth UpdateBackendAuthUserPoolConfigOAuth `json:"OAuth,omitempty"`
	Passwordpolicy UpdateBackendAuthUserPoolConfigPasswordPolicy `json:"PasswordPolicy,omitempty"`
	Verificationmessage UpdateBackendAuthUserPoolConfigVerificationMessage `json:"VerificationMessage,omitempty"`
}

// ImportBackendAuthResponse represents the ImportBackendAuthResponse schema from the OpenAPI specification
type ImportBackendAuthResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// BackendStoragePermissions represents the BackendStoragePermissions schema from the OpenAPI specification
type BackendStoragePermissions struct {
	Authenticated interface{} `json:"Authenticated"`
	Unauthenticated interface{} `json:"UnAuthenticated,omitempty"`
}

// GetBackendRequest represents the GetBackendRequest schema from the OpenAPI specification
type GetBackendRequest struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
}

// GetBackendAPIRequest represents the GetBackendAPIRequest schema from the OpenAPI specification
type GetBackendAPIRequest struct {
	Resourceconfig DeleteBackendAPIRequestResourceConfig `json:"ResourceConfig,omitempty"`
	Resourcename interface{} `json:"ResourceName"`
}

// DeleteBackendAPIResponse represents the DeleteBackendAPIResponse schema from the OpenAPI specification
type DeleteBackendAPIResponse struct {
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
}

// UpdateBackendAuthUserPoolConfigPasswordPolicy represents the UpdateBackendAuthUserPoolConfigPasswordPolicy schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfigPasswordPolicy struct {
	Additionalconstraints interface{} `json:"AdditionalConstraints,omitempty"`
	Minimumlength interface{} `json:"MinimumLength,omitempty"`
}

// UpdateBackendAuthrequestresourceConfig represents the UpdateBackendAuthrequestresourceConfig schema from the OpenAPI specification
type UpdateBackendAuthrequestresourceConfig struct {
	Authresources interface{} `json:"AuthResources,omitempty"`
	Identitypoolconfigs UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Userpoolconfigs UpdateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs,omitempty"`
}

// CreateBackendAuthUserPoolConfigMfa represents the CreateBackendAuthUserPoolConfigMfa schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfigMfa struct {
	Mfamode interface{} `json:"MFAMode"`
	Settings CreateBackendAuthMFAConfigSettings `json:"Settings,omitempty"`
}

// CreateBackendAuthForgotPasswordConfig represents the CreateBackendAuthForgotPasswordConfig schema from the OpenAPI specification
type CreateBackendAuthForgotPasswordConfig struct {
	Deliverymethod interface{} `json:"DeliveryMethod"`
	Emailsettings CreateBackendAuthForgotPasswordConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthForgotPasswordConfigSmsSettings `json:"SmsSettings,omitempty"`
}

// CloneBackendResponse represents the CloneBackendResponse schema from the OpenAPI specification
type CloneBackendResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetBackendStoragerequest represents the GetBackendStoragerequest schema from the OpenAPI specification
type GetBackendStoragerequest struct {
	Resourcename string `json:"resourceName"` // The name of the storage resource.
}

// BackendAPIAuthTypeSettings represents the BackendAPIAuthTypeSettings schema from the OpenAPI specification
type BackendAPIAuthTypeSettings struct {
	Openidauthttl interface{} `json:"OpenIDAuthTTL,omitempty"`
	Openidclientid interface{} `json:"OpenIDClientId,omitempty"`
	Openidiatttl interface{} `json:"OpenIDIatTTL,omitempty"`
	Openidissueurl interface{} `json:"OpenIDIssueURL,omitempty"`
	Openidprovidername interface{} `json:"OpenIDProviderName,omitempty"`
	Cognitouserpoolid interface{} `json:"CognitoUserPoolId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Expirationtime interface{} `json:"ExpirationTime,omitempty"`
}

// UpdateBackendAuthUserPoolConfigMfa represents the UpdateBackendAuthUserPoolConfigMfa schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfigMfa struct {
	Mfamode interface{} `json:"MFAMode,omitempty"`
	Settings UpdateBackendAuthMFAConfigSettings `json:"Settings,omitempty"`
}

// CreateBackendAuthResponse represents the CreateBackendAuthResponse schema from the OpenAPI specification
type CreateBackendAuthResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// CreateBackendAuthPasswordPolicyConfig represents the CreateBackendAuthPasswordPolicyConfig schema from the OpenAPI specification
type CreateBackendAuthPasswordPolicyConfig struct {
	Additionalconstraints interface{} `json:"AdditionalConstraints,omitempty"`
	Minimumlength interface{} `json:"MinimumLength"`
}

// UpdateBackendAuthVerificationMessageConfig represents the UpdateBackendAuthVerificationMessageConfig schema from the OpenAPI specification
type UpdateBackendAuthVerificationMessageConfig struct {
	Deliverymethod interface{} `json:"DeliveryMethod"`
	Emailsettings CreateBackendAuthVerificationMessageConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthVerificationMessageConfigSmsSettings `json:"SmsSettings,omitempty"`
}

// UpdateBackendConfigResponse represents the UpdateBackendConfigResponse schema from the OpenAPI specification
type UpdateBackendConfigResponse struct {
	Loginauthconfig UpdateBackendConfigResponseLoginAuthConfig `json:"LoginAuthConfig,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendmanagerappid interface{} `json:"BackendManagerAppId,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// ListBackendJobsRequest represents the ListBackendJobsRequest schema from the OpenAPI specification
type ListBackendJobsRequest struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateBackendAuthrequestresourceConfigUserPoolConfigs represents the UpdateBackendAuthrequestresourceConfigUserPoolConfigs schema from the OpenAPI specification
type UpdateBackendAuthrequestresourceConfigUserPoolConfigs struct {
	Verificationmessage UpdateBackendAuthUserPoolConfigVerificationMessage `json:"VerificationMessage,omitempty"`
	Forgotpassword UpdateBackendAuthUserPoolConfigForgotPassword `json:"ForgotPassword,omitempty"`
	Mfa UpdateBackendAuthUserPoolConfigMfa `json:"Mfa,omitempty"`
	Oauth UpdateBackendAuthUserPoolConfigOAuth `json:"OAuth,omitempty"`
	Passwordpolicy UpdateBackendAuthUserPoolConfigPasswordPolicy `json:"PasswordPolicy,omitempty"`
}

// ImportBackendAuthRequest represents the ImportBackendAuthRequest schema from the OpenAPI specification
type ImportBackendAuthRequest struct {
	Userpoolid interface{} `json:"UserPoolId"`
	Webclientid interface{} `json:"WebClientId"`
	Identitypoolid interface{} `json:"IdentityPoolId,omitempty"`
	Nativeclientid interface{} `json:"NativeClientId"`
}

// UpdateBackendAuthUserPoolConfigVerificationMessage represents the UpdateBackendAuthUserPoolConfigVerificationMessage schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfigVerificationMessage struct {
	Deliverymethod interface{} `json:"DeliveryMethod"`
	Emailsettings CreateBackendAuthVerificationMessageConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthVerificationMessageConfigSmsSettings `json:"SmsSettings,omitempty"`
}

// CreateBackendAuthUserPoolConfigVerificationMessage represents the CreateBackendAuthUserPoolConfigVerificationMessage schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfigVerificationMessage struct {
	Emailsettings CreateBackendAuthVerificationMessageConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthVerificationMessageConfigSmsSettings `json:"SmsSettings,omitempty"`
	Deliverymethod interface{} `json:"DeliveryMethod"`
}

// RemoveAllBackendsRequest represents the RemoveAllBackendsRequest schema from the OpenAPI specification
type RemoveAllBackendsRequest struct {
	Cleanamplifyapp interface{} `json:"CleanAmplifyApp,omitempty"`
}

// DeleteBackendStorageRequest represents the DeleteBackendStorageRequest schema from the OpenAPI specification
type DeleteBackendStorageRequest struct {
	Resourcename interface{} `json:"ResourceName"`
	Servicename interface{} `json:"ServiceName"`
}

// ImportBackendStorageRequest represents the ImportBackendStorageRequest schema from the OpenAPI specification
type ImportBackendStorageRequest struct {
	Servicename interface{} `json:"ServiceName"`
	Bucketname interface{} `json:"BucketName,omitempty"`
}

// CreateBackendAuthrequest represents the CreateBackendAuthrequest schema from the OpenAPI specification
type CreateBackendAuthrequest struct {
	Backendenvironmentname string `json:"backendEnvironmentName"` // The name of the backend environment.
	Resourceconfig CreateBackendAuthrequestresourceConfig `json:"resourceConfig"` // Defines the resource configuration when creating an auth resource in your Amplify project.
	Resourcename string `json:"resourceName"` // The name of this resource.
}

// UpdateBackendJobRequest represents the UpdateBackendJobRequest schema from the OpenAPI specification
type UpdateBackendJobRequest struct {
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetTokenRequest represents the GetTokenRequest schema from the OpenAPI specification
type GetTokenRequest struct {
}

// RemoveBackendConfigResponse represents the RemoveBackendConfigResponse schema from the OpenAPI specification
type RemoveBackendConfigResponse struct {
	ErrorField interface{} `json:"Error,omitempty"`
}

// CreateBackendAPIrequestresourceConfig represents the CreateBackendAPIrequestresourceConfig schema from the OpenAPI specification
type CreateBackendAPIrequestresourceConfig struct {
	Additionalauthtypes interface{} `json:"AdditionalAuthTypes,omitempty"`
	Apiname interface{} `json:"ApiName,omitempty"`
	Conflictresolution CreateBackendAPIrequestresourceConfigConflictResolution `json:"ConflictResolution,omitempty"`
	Defaultauthtype CreateBackendAPIrequestresourceConfigDefaultAuthType `json:"DefaultAuthType,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Transformschema interface{} `json:"TransformSchema,omitempty"`
}

// GetBackendStorageResourceConfig represents the GetBackendStorageResourceConfig schema from the OpenAPI specification
type GetBackendStorageResourceConfig struct {
	Servicename interface{} `json:"ServiceName"`
	Bucketname interface{} `json:"BucketName,omitempty"`
	Imported interface{} `json:"Imported"`
	Permissions CreateBackendStoragerequestresourceConfigPermissions `json:"Permissions,omitempty"`
}

// CreateBackendAPIRequestResourceConfig represents the CreateBackendAPIRequestResourceConfig schema from the OpenAPI specification
type CreateBackendAPIRequestResourceConfig struct {
	Defaultauthtype CreateBackendAPIrequestresourceConfigDefaultAuthType `json:"DefaultAuthType,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Transformschema interface{} `json:"TransformSchema,omitempty"`
	Additionalauthtypes interface{} `json:"AdditionalAuthTypes,omitempty"`
	Apiname interface{} `json:"ApiName,omitempty"`
	Conflictresolution CreateBackendAPIrequestresourceConfigConflictResolution `json:"ConflictResolution,omitempty"`
}

// UpdateBackendConfigRequestLoginAuthConfig represents the UpdateBackendConfigRequestLoginAuthConfig schema from the OpenAPI specification
type UpdateBackendConfigRequestLoginAuthConfig struct {
	Awsuserpoolswebclientid interface{} `json:"AwsUserPoolsWebClientId,omitempty"`
	Awscognitoidentitypoolid interface{} `json:"AwsCognitoIdentityPoolId,omitempty"`
	Awscognitoregion interface{} `json:"AwsCognitoRegion,omitempty"`
	Awsuserpoolsid interface{} `json:"AwsUserPoolsId,omitempty"`
}

// CreateBackendRequest represents the CreateBackendRequest schema from the OpenAPI specification
type CreateBackendRequest struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName"`
	Resourceconfig interface{} `json:"ResourceConfig,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
	Appid interface{} `json:"AppId"`
	Appname interface{} `json:"AppName"`
}

// CloneBackendRequest represents the CloneBackendRequest schema from the OpenAPI specification
type CloneBackendRequest struct {
	Targetenvironmentname interface{} `json:"TargetEnvironmentName"`
}

// GetBackendAuthRequest represents the GetBackendAuthRequest schema from the OpenAPI specification
type GetBackendAuthRequest struct {
	Resourcename interface{} `json:"ResourceName"`
}

// CreateBackendAuthUserPoolConfigPasswordPolicy represents the CreateBackendAuthUserPoolConfigPasswordPolicy schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfigPasswordPolicy struct {
	Additionalconstraints interface{} `json:"AdditionalConstraints,omitempty"`
	Minimumlength interface{} `json:"MinimumLength"`
}

// GetBackendJobRequest represents the GetBackendJobRequest schema from the OpenAPI specification
type GetBackendJobRequest struct {
}

// CreateBackendAuthrequestresourceConfigUserPoolConfigs represents the CreateBackendAuthrequestresourceConfigUserPoolConfigs schema from the OpenAPI specification
type CreateBackendAuthrequestresourceConfigUserPoolConfigs struct {
	Forgotpassword CreateBackendAuthUserPoolConfigForgotPassword `json:"ForgotPassword,omitempty"`
	Mfa CreateBackendAuthUserPoolConfigMfa `json:"Mfa,omitempty"`
	Oauth CreateBackendAuthUserPoolConfigOAuth `json:"OAuth,omitempty"`
	Passwordpolicy CreateBackendAuthUserPoolConfigPasswordPolicy `json:"PasswordPolicy,omitempty"`
	Requiredsignupattributes interface{} `json:"RequiredSignUpAttributes"`
	Signinmethod interface{} `json:"SignInMethod"`
	Userpoolname interface{} `json:"UserPoolName"`
	Verificationmessage CreateBackendAuthUserPoolConfigVerificationMessage `json:"VerificationMessage,omitempty"`
}

// GetBackendAuthResponse represents the GetBackendAuthResponse schema from the OpenAPI specification
type GetBackendAuthResponse struct {
	Resourceconfig GetBackendAuthResponseResourceConfig `json:"ResourceConfig,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// DeleteTokenResponse represents the DeleteTokenResponse schema from the OpenAPI specification
type DeleteTokenResponse struct {
	Issuccess interface{} `json:"IsSuccess,omitempty"`
}

// CreateBackendAuthMFAConfig represents the CreateBackendAuthMFAConfig schema from the OpenAPI specification
type CreateBackendAuthMFAConfig struct {
	Mfamode interface{} `json:"MFAMode"`
	Settings CreateBackendAuthMFAConfigSettings `json:"Settings,omitempty"`
}

// ResourceConfig represents the ResourceConfig schema from the OpenAPI specification
type ResourceConfig struct {
}

// GenerateBackendAPIModelsRequest represents the GenerateBackendAPIModelsRequest schema from the OpenAPI specification
type GenerateBackendAPIModelsRequest struct {
	Resourcename interface{} `json:"ResourceName"`
}

// CreateBackendAPIResponse represents the CreateBackendAPIResponse schema from the OpenAPI specification
type CreateBackendAPIResponse struct {
	ErrorField interface{} `json:"Error,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
}

// LoginAuthConfigReqObj represents the LoginAuthConfigReqObj schema from the OpenAPI specification
type LoginAuthConfigReqObj struct {
	Awsuserpoolsid interface{} `json:"AwsUserPoolsId,omitempty"`
	Awsuserpoolswebclientid interface{} `json:"AwsUserPoolsWebClientId,omitempty"`
	Awscognitoidentitypoolid interface{} `json:"AwsCognitoIdentityPoolId,omitempty"`
	Awscognitoregion interface{} `json:"AwsCognitoRegion,omitempty"`
}

// UpdateBackendAuthIdentityPoolConfig represents the UpdateBackendAuthIdentityPoolConfig schema from the OpenAPI specification
type UpdateBackendAuthIdentityPoolConfig struct {
	Unauthenticatedlogin interface{} `json:"UnauthenticatedLogin,omitempty"`
}

// DeleteBackendAuthRequest represents the DeleteBackendAuthRequest schema from the OpenAPI specification
type DeleteBackendAuthRequest struct {
	Resourcename interface{} `json:"ResourceName"`
}

// GetBackendStorageResponse represents the GetBackendStorageResponse schema from the OpenAPI specification
type GetBackendStorageResponse struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Resourceconfig GetBackendStorageResponseResourceConfig `json:"ResourceConfig,omitempty"`
	Resourcename interface{} `json:"ResourceName,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
}

// CreateBackendAuthForgotPasswordConfigEmailSettings represents the CreateBackendAuthForgotPasswordConfigEmailSettings schema from the OpenAPI specification
type CreateBackendAuthForgotPasswordConfigEmailSettings struct {
	Emailmessage interface{} `json:"EmailMessage,omitempty"`
	Emailsubject interface{} `json:"EmailSubject,omitempty"`
}

// EmailSettings represents the EmailSettings schema from the OpenAPI specification
type EmailSettings struct {
	Emailmessage interface{} `json:"EmailMessage,omitempty"`
	Emailsubject interface{} `json:"EmailSubject,omitempty"`
}

// CreateBackendrequest represents the CreateBackendrequest schema from the OpenAPI specification
type CreateBackendrequest struct {
	Appid string `json:"appId"` // The app ID.
	Appname string `json:"appName"` // The name of the app.
	Backendenvironmentname string `json:"backendEnvironmentName"` // The name of the backend environment.
	Resourceconfig map[string]interface{} `json:"resourceConfig,omitempty"` // Defines the resource configuration for the data model in your Amplify project.
	Resourcename string `json:"resourceName,omitempty"` // The name of the resource.
}

// CreateBackendAPIrequestresourceConfigConflictResolution represents the CreateBackendAPIrequestresourceConfigConflictResolution schema from the OpenAPI specification
type CreateBackendAPIrequestresourceConfigConflictResolution struct {
	Resolutionstrategy interface{} `json:"ResolutionStrategy,omitempty"`
}

// GetBackendAPIModelsRequest represents the GetBackendAPIModelsRequest schema from the OpenAPI specification
type GetBackendAPIModelsRequest struct {
	Resourcename interface{} `json:"ResourceName"`
}

// DeleteBackendAuthrequest represents the DeleteBackendAuthrequest schema from the OpenAPI specification
type DeleteBackendAuthrequest struct {
	Resourcename string `json:"resourceName"` // The name of this resource.
}

// CreateBackendAuthrequestresourceConfig represents the CreateBackendAuthrequestresourceConfig schema from the OpenAPI specification
type CreateBackendAuthrequestresourceConfig struct {
	Authresources interface{} `json:"AuthResources,omitempty"`
	Identitypoolconfigs CreateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
	Service interface{} `json:"Service,omitempty"`
	Userpoolconfigs CreateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs,omitempty"`
}

// CreateBackendConfigRequest represents the CreateBackendConfigRequest schema from the OpenAPI specification
type CreateBackendConfigRequest struct {
	Backendmanagerappid interface{} `json:"BackendManagerAppId,omitempty"`
}

// SocialProviderSettings represents the SocialProviderSettings schema from the OpenAPI specification
type SocialProviderSettings struct {
	Facebook BackendAuthSocialProviderConfig `json:"Facebook,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Google BackendAuthSocialProviderConfig `json:"Google,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Loginwithamazon BackendAuthSocialProviderConfig `json:"LoginWithAmazon,omitempty"` // Describes third-party social federation configurations for allowing your app users to sign in using OAuth.
	Signinwithapple BackendAuthAppleProviderConfig `json:"SignInWithApple,omitempty"` // Describes Apple social federation configurations for allowing your app users to sign in using OAuth.
}

// ImportBackendStoragerequest represents the ImportBackendStoragerequest schema from the OpenAPI specification
type ImportBackendStoragerequest struct {
	Bucketname string `json:"bucketName,omitempty"` // The name of the S3 bucket.
	Servicename string `json:"serviceName"` // The name of the storage service.
}

// BackendAuthAppleProviderConfig represents the BackendAuthAppleProviderConfig schema from the OpenAPI specification
type BackendAuthAppleProviderConfig struct {
	Clientid interface{} `json:"ClientId,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
	Privatekey interface{} `json:"PrivateKey,omitempty"`
	Teamid interface{} `json:"TeamId,omitempty"`
}

// DeleteBackendResponse represents the DeleteBackendResponse schema from the OpenAPI specification
type DeleteBackendResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// CreateBackendAuthUserPoolConfigForgotPassword represents the CreateBackendAuthUserPoolConfigForgotPassword schema from the OpenAPI specification
type CreateBackendAuthUserPoolConfigForgotPassword struct {
	Deliverymethod interface{} `json:"DeliveryMethod"`
	Emailsettings CreateBackendAuthForgotPasswordConfigEmailSettings `json:"EmailSettings,omitempty"`
	Smssettings CreateBackendAuthForgotPasswordConfigSmsSettings `json:"SmsSettings,omitempty"`
}

// UpdateBackendAuthForgotPasswordConfigSmsSettings represents the UpdateBackendAuthForgotPasswordConfigSmsSettings schema from the OpenAPI specification
type UpdateBackendAuthForgotPasswordConfigSmsSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
}

// ImportBackendStorageResponse represents the ImportBackendStorageResponse schema from the OpenAPI specification
type ImportBackendStorageResponse struct {
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	Jobid interface{} `json:"JobId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// UpdateBackendAuthMFAConfigSettings represents the UpdateBackendAuthMFAConfigSettings schema from the OpenAPI specification
type UpdateBackendAuthMFAConfigSettings struct {
	Smsmessage interface{} `json:"SmsMessage,omitempty"`
	Mfatypes interface{} `json:"MfaTypes,omitempty"`
}

// UpdateBackendAuthUserPoolConfigOAuth represents the UpdateBackendAuthUserPoolConfigOAuth schema from the OpenAPI specification
type UpdateBackendAuthUserPoolConfigOAuth struct {
	Socialprovidersettings UpdateBackendAuthOAuthConfigSocialProviderSettings `json:"SocialProviderSettings,omitempty"`
	Domainprefix interface{} `json:"DomainPrefix,omitempty"`
	Oauthgranttype interface{} `json:"OAuthGrantType,omitempty"`
	Oauthscopes interface{} `json:"OAuthScopes,omitempty"`
	Redirectsigninuris interface{} `json:"RedirectSignInURIs,omitempty"`
	Redirectsignouturis interface{} `json:"RedirectSignOutURIs,omitempty"`
}

// DeleteBackendAPIrequest represents the DeleteBackendAPIrequest schema from the OpenAPI specification
type DeleteBackendAPIrequest struct {
	Resourceconfig CreateBackendAPIrequestresourceConfig `json:"resourceConfig,omitempty"` // The resource config for the data model, configured as a part of the Amplify project.
	Resourcename string `json:"resourceName"` // The name of this resource.
}

// ListS3Bucketsrequest represents the ListS3Bucketsrequest schema from the OpenAPI specification
type ListS3Bucketsrequest struct {
	Nexttoken string `json:"nextToken,omitempty"` // Reserved for future use.
}

// UpdateBackendConfigrequest represents the UpdateBackendConfigrequest schema from the OpenAPI specification
type UpdateBackendConfigrequest struct {
	Loginauthconfig UpdateBackendConfigrequestloginAuthConfig `json:"loginAuthConfig,omitempty"` // The request object for this operation.
}

// UpdateBackendAuthOAuthConfig represents the UpdateBackendAuthOAuthConfig schema from the OpenAPI specification
type UpdateBackendAuthOAuthConfig struct {
	Socialprovidersettings UpdateBackendAuthOAuthConfigSocialProviderSettings `json:"SocialProviderSettings,omitempty"`
	Domainprefix interface{} `json:"DomainPrefix,omitempty"`
	Oauthgranttype interface{} `json:"OAuthGrantType,omitempty"`
	Oauthscopes interface{} `json:"OAuthScopes,omitempty"`
	Redirectsigninuris interface{} `json:"RedirectSignInURIs,omitempty"`
	Redirectsignouturis interface{} `json:"RedirectSignOutURIs,omitempty"`
}

// DeleteBackendRequest represents the DeleteBackendRequest schema from the OpenAPI specification
type DeleteBackendRequest struct {
}

// DeleteBackendAuthResponse represents the DeleteBackendAuthResponse schema from the OpenAPI specification
type DeleteBackendAuthResponse struct {
	Jobid interface{} `json:"JobId,omitempty"`
	Operation interface{} `json:"Operation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Appid interface{} `json:"AppId,omitempty"`
	Backendenvironmentname interface{} `json:"BackendEnvironmentName,omitempty"`
	ErrorField interface{} `json:"Error,omitempty"`
}

// UpdateBackendAuthMFAConfig represents the UpdateBackendAuthMFAConfig schema from the OpenAPI specification
type UpdateBackendAuthMFAConfig struct {
	Settings UpdateBackendAuthMFAConfigSettings `json:"Settings,omitempty"`
	Mfamode interface{} `json:"MFAMode,omitempty"`
}

// UpdateBackendAuthRequestResourceConfig represents the UpdateBackendAuthRequestResourceConfig schema from the OpenAPI specification
type UpdateBackendAuthRequestResourceConfig struct {
	Service interface{} `json:"Service"`
	Userpoolconfigs UpdateBackendAuthrequestresourceConfigUserPoolConfigs `json:"UserPoolConfigs"`
	Authresources interface{} `json:"AuthResources"`
	Identitypoolconfigs UpdateBackendAuthrequestresourceConfigIdentityPoolConfigs `json:"IdentityPoolConfigs,omitempty"`
}

// CreateBackendAuthRequest represents the CreateBackendAuthRequest schema from the OpenAPI specification
type CreateBackendAuthRequest struct {
	Backendenvironmentname interface{} `json:"BackendEnvironmentName"`
	Resourceconfig CreateBackendAuthRequestResourceConfig `json:"ResourceConfig"`
	Resourcename interface{} `json:"ResourceName"`
}

// CreateBackendAuthVerificationMessageConfigEmailSettings represents the CreateBackendAuthVerificationMessageConfigEmailSettings schema from the OpenAPI specification
type CreateBackendAuthVerificationMessageConfigEmailSettings struct {
	Emailsubject interface{} `json:"EmailSubject,omitempty"`
	Emailmessage interface{} `json:"EmailMessage,omitempty"`
}
