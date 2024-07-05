package pet

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
    i01fa6ea29b6cfe613045be406b47b662b5f8c510b65673a105a1b4b6eb08fe62 "github.com/microsoft/kiota-samples/petstore/go/utilities/pet/findbystatus"
)

// FindByStatusRequestBuilder builds and executes requests for operations under \pet\findByStatus
type FindByStatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FindByStatusRequestBuilderGetQueryParameters multiple status values can be provided with comma separated strings
type FindByStatusRequestBuilderGetQueryParameters struct {
    // Status values that need to be considered for filter
    // Deprecated: This property is deprecated, use StatusAsGetStatusQueryParameterType instead
    Status []string `uriparametername:"status"`
    // Status values that need to be considered for filter
    StatusAsGetStatusQueryParameterType []i01fa6ea29b6cfe613045be406b47b662b5f8c510b65673a105a1b4b6eb08fe62.GetStatusQueryParameterType `uriparametername:"status"`
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
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pet/findByStatus?status={status}", pathParameters),
    }
    return m
}
// NewFindByStatusRequestBuilder instantiates a new FindByStatusRequestBuilder and sets the default values.
func NewFindByStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindByStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindByStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Get multiple status values can be provided with comma separated strings
// returns a []Petable when successful
func (m *FindByStatusRequestBuilder) Get(ctx context.Context, requestConfiguration *FindByStatusRequestBuilderGetRequestConfiguration)([]idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.CreatePetFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable)
        }
    }
    return val, nil
}
// ToGetRequestInformation multiple status values can be provided with comma separated strings
// returns a *RequestInformation when successful
func (m *FindByStatusRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindByStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FindByStatusRequestBuilder when successful
func (m *FindByStatusRequestBuilder) WithUrl(rawUrl string)(*FindByStatusRequestBuilder) {
    return NewFindByStatusRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
