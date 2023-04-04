package user

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
)

// CreateWithListRequestBuilder builds and executes requests for operations under \user\createWithList
type CreateWithListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CreateWithListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CreateWithListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCreateWithListRequestBuilderInternal instantiates a new CreateWithListRequestBuilder and sets the default values.
func NewCreateWithListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateWithListRequestBuilder) {
    m := &CreateWithListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/user/createWithList";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCreateWithListRequestBuilder instantiates a new CreateWithListRequestBuilder and sets the default values.
func NewCreateWithListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CreateWithListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateWithListRequestBuilderInternal(urlParams, requestAdapter)
}
// Post creates list of users with given input array
func (m *CreateWithListRequestBuilder) Post(ctx context.Context, body []idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Userable, requestConfiguration *CreateWithListRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
func (m *CreateWithListRequestBuilder) ToPostRequestInformation(ctx context.Context, body []idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Userable, requestConfiguration *CreateWithListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(body))
    for i, v := range body {
        cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
    }
    err := requestInfo.SetContentFromParsableCollection(ctx, m.requestAdapter, "application/json", cast)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
