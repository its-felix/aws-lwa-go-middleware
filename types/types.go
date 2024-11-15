package types

import "github.com/aws/aws-lambda-go/lambdacontext"

type CognitoIdentity struct {
	CognitoIdentityID     string `json:"cognito_identity_id"`
	CognitoIdentityPoolID string `json:"cognito_identity_pool_id"`
}

type EnvConfig struct {
	FunctionName string `json:"function_name"`
	Memory       int    `json:"memory"`
	Version      string `json:"version"`
	LogStream    string `json:"log_stream"`
	LogGroup     string `json:"log_group"`
}

type LambdaContext struct {
	RequestID          string                       `json:"request_id"`
	Deadline           int64                        `json:"deadline"`
	InvokedFunctionArn string                       `json:"invoked_function_arn"`
	XrayTraceID        string                       `json:"xray_trace_id"`
	ClientContext      *lambdacontext.ClientContext `json:"client_context,omitempty"`
	Identity           *CognitoIdentity             `json:"identity,omitempty"`
	EnvConfig          EnvConfig                    `json:"env_config"`
}
