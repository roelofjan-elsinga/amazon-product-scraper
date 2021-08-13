package models

type GatewayRequest struct {
	Resource   string `json:"resource"`
	Path       string `json:"path"`
	HTTPMethod string `json:"httpMethod"`
	Headers    struct {
		Header1 string `json:"header1"`
		Header2 string `json:"header2"`
	} `json:"headers"`
	MultiValueHeaders struct {
		Header1 []string `json:"header1"`
		Header2 []string `json:"header2"`
	} `json:"multiValueHeaders"`
	QueryStringParameters struct {
		Parameter1 string `json:"parameter1"`
		Parameter2 string `json:"parameter2"`
	} `json:"queryStringParameters"`
	MultiValueQueryStringParameters struct {
		Parameter1 []string `json:"parameter1"`
		Parameter2 []string `json:"parameter2"`
	} `json:"multiValueQueryStringParameters"`
	RequestContext struct {
		AccountID  string `json:"accountId"`
		APIID      string `json:"apiId"`
		Authorizer struct {
			Claims interface{} `json:"claims"`
			Scopes interface{} `json:"scopes"`
		} `json:"authorizer"`
		DomainName        string `json:"domainName"`
		DomainPrefix      string `json:"domainPrefix"`
		ExtendedRequestID string `json:"extendedRequestId"`
		HTTPMethod        string `json:"httpMethod"`
		Identity          struct {
			AccessKey                     interface{} `json:"accessKey"`
			AccountID                     interface{} `json:"accountId"`
			Caller                        interface{} `json:"caller"`
			CognitoAuthenticationProvider interface{} `json:"cognitoAuthenticationProvider"`
			CognitoAuthenticationType     interface{} `json:"cognitoAuthenticationType"`
			CognitoIdentityID             interface{} `json:"cognitoIdentityId"`
			CognitoIdentityPoolID         interface{} `json:"cognitoIdentityPoolId"`
			PrincipalOrgID                interface{} `json:"principalOrgId"`
			SourceIP                      string      `json:"sourceIp"`
			User                          interface{} `json:"user"`
			UserAgent                     string      `json:"userAgent"`
			UserArn                       interface{} `json:"userArn"`
			ClientCert                    struct {
				ClientCertPem string `json:"clientCertPem"`
				SubjectDN     string `json:"subjectDN"`
				IssuerDN      string `json:"issuerDN"`
				SerialNumber  string `json:"serialNumber"`
				Validity      struct {
					NotBefore string `json:"notBefore"`
					NotAfter  string `json:"notAfter"`
				} `json:"validity"`
			} `json:"clientCert"`
		} `json:"identity"`
		Path             string      `json:"path"`
		Protocol         string      `json:"protocol"`
		RequestID        string      `json:"requestId"`
		RequestTime      string      `json:"requestTime"`
		RequestTimeEpoch int64       `json:"requestTimeEpoch"`
		ResourceID       interface{} `json:"resourceId"`
		ResourcePath     string      `json:"resourcePath"`
		Stage            string      `json:"stage"`
	} `json:"requestContext"`
	PathParameters  interface{} `json:"pathParameters"`
	StageVariables  interface{} `json:"stageVariables"`
	Body            string      `json:"body"`
	IsBase64Encoded bool        `json:"isBase64Encoded"`
}

type GatewayResponse struct {
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
}
