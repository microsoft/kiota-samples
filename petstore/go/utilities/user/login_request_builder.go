package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LoginRequestBuilder builds and executes requests for operations under \user\login
type LoginRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LoginRequestBuilderGetQueryParameters logs user into the system
type LoginRequestBuilderGetQueryParameters struct {
    // The password for login in clear text
    Password *string `uriparametername:"password"`
    // The user name for login
    Username *string `uriparametername:"username"`
}
// LoginRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LoginRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LoginRequestBuilderGetQueryParameters
}
// NewLoginRequestBuilderInternal instantiates a new LoginRequestBuilder and sets the default values.
func NewLoginRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoginRequestBuilder) {
    m := &LoginRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/user/login?password={password}&username={username}", pathParameters),
    }
    return m
}
// NewLoginRequestBuilder instantiates a new LoginRequestBuilder and sets the default values.
func NewLoginRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LoginRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLoginRequestBuilderInternal(urlParams, requestAdapter)
}
// Get logs user into the system
// returns a *string when successful
func (m *LoginRequestBuilder) Get(ctx context.Context, requestConfiguration *LoginRequestBuilderGetRequestConfiguration)(*string, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*string), nil
}
// returns a *RequestInformation when successful
func (m *LoginRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LoginRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LoginRequestBuilder when successful
func (m *LoginRequestBuilder) WithUrl(rawUrl string)(*LoginRequestBuilder) {
    return NewLoginRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
