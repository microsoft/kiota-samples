package pet

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
)

// FindByStatusRequestBuilder builds and executes requests for operations under \pet\findByStatus
type FindByStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FindByStatusRequestBuilderGetQueryParameters multiple status values can be provided with comma separated strings
type FindByStatusRequestBuilderGetQueryParameters struct {
    // Status values that need to be considered for filter
    Status []string
}
// FindByStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindByStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FindByStatusRequestBuilderGetQueryParameters
}
// NewFindByStatusRequestBuilderInternal instantiates a new FindByStatusRequestBuilder and sets the default values.
func NewFindByStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindByStatusRequestBuilder) {
    m := &FindByStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/pet/findByStatus{?status*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewFindByStatusRequestBuilder instantiates a new FindByStatusRequestBuilder and sets the default values.
func NewFindByStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindByStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindByStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get multiple status values can be provided with comma separated strings
func (m *FindByStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *FindByStatusRequestBuilderGetRequestConfiguration)([]idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.CreatePetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable, len(res))
    for i, v := range res {
        val[i] = v.(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable)
    }
    return val, nil
}
// ToGetRequestInformation multiple status values can be provided with comma separated strings
func (m *FindByStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindByStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
