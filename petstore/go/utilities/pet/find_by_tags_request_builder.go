package pet

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
)

// FindByTagsRequestBuilder builds and executes requests for operations under \pet\findByTags
type FindByTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FindByTagsRequestBuilderGetQueryParameters multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
type FindByTagsRequestBuilderGetQueryParameters struct {
    // Tags to filter by
    Tags []string `uriparametername:"tags"`
}
// FindByTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindByTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FindByTagsRequestBuilderGetQueryParameters
}
// NewFindByTagsRequestBuilderInternal instantiates a new FindByTagsRequestBuilder and sets the default values.
func NewFindByTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindByTagsRequestBuilder) {
    m := &FindByTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pet/findByTags?tags={tags}", pathParameters),
    }
    return m
}
// NewFindByTagsRequestBuilder instantiates a new FindByTagsRequestBuilder and sets the default values.
func NewFindByTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindByTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindByTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
// Deprecated: 
// returns a []Petable when successful
func (m *FindByTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *FindByTagsRequestBuilderGetRequestConfiguration)([]idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Petable, error) {
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
// ToGetRequestInformation multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
// Deprecated: 
// returns a *RequestInformation when successful
func (m *FindByTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindByTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: 
// returns a *FindByTagsRequestBuilder when successful
func (m *FindByTagsRequestBuilder) WithUrl(rawUrl string)(*FindByTagsRequestBuilder) {
    return NewFindByTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
