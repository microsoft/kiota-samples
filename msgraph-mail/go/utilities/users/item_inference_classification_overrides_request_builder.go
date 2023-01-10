package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

// ItemInferenceClassificationOverridesRequestBuilder builds and executes requests for operations under \users\{user-id}\inferenceClassification\overrides
type ItemInferenceClassificationOverridesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
type ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInferenceClassificationOverridesRequestBuilderGetQueryParameters
}
// ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemInferenceClassificationOverridesRequestBuilderInternal instantiates a new OverridesRequestBuilder and sets the default values.
func NewItemInferenceClassificationOverridesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceClassificationOverridesRequestBuilder) {
    m := &ItemInferenceClassificationOverridesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemInferenceClassificationOverridesRequestBuilder instantiates a new OverridesRequestBuilder and sets the default values.
func NewItemInferenceClassificationOverridesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInferenceClassificationOverridesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInferenceClassificationOverridesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/inferenceclassification-list-overrides?view=graph-rest-1.0
func (m *ItemInferenceClassificationOverridesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideCollectionResponseable), nil
}
// Post create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. **Note**
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/inferenceclassification-post-overrides?view=graph-rest-1.0
func (m *ItemInferenceClassificationOverridesRequestBuilder) Post(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideable, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateInferenceClassificationOverrideFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideable), nil
}
// ToGetRequestInformation get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
func (m *ItemInferenceClassificationOverridesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. **Note**
func (m *ItemInferenceClassificationOverridesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverrideable, requestConfiguration *ItemInferenceClassificationOverridesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
