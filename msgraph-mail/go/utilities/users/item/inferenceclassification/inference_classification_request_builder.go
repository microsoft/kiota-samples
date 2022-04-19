package inferenceclassification

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i3b892eb54cedbd9cc555b9f1a7958d6152b7730fa895edaed6eaa7e22b3c15ca "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/inferenceclassification/overrides"
    ie9af6222ea455c01969cfed303f450bfed10eddca9c64da588590323f4684264 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/inferenceclassification/overrides/item"
)

// InferenceClassificationRequestBuilder builds and executes requests for operations under \users\{user-id}\inferenceClassification
type InferenceClassificationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InferenceClassificationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InferenceClassificationRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// InferenceClassificationRequestBuilderGetQueryParameters relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
type InferenceClassificationRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InferenceClassificationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InferenceClassificationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InferenceClassificationRequestBuilderGetQueryParameters
}
// InferenceClassificationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InferenceClassificationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInferenceClassificationRequestBuilderInternal instantiates a new InferenceClassificationRequestBuilder and sets the default values.
func NewInferenceClassificationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InferenceClassificationRequestBuilder) {
    m := &InferenceClassificationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInferenceClassificationRequestBuilder instantiates a new InferenceClassificationRequestBuilder and sets the default values.
func NewInferenceClassificationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InferenceClassificationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInferenceClassificationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *InferenceClassificationRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *InferenceClassificationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreatePatchRequestInformation(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable, requestConfiguration *InferenceClassificationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) Delete(requestConfiguration *InferenceClassificationRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) DeleteWithResponseHandler(requestConfiguration *InferenceClassificationRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) Get(requestConfiguration *InferenceClassificationRequestBuilderGetRequestConfiguration)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) GetWithResponseHandler(requestConfiguration *InferenceClassificationRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateInferenceClassificationFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable), nil
}
// Overrides the overrides property
func (m *InferenceClassificationRequestBuilder) Overrides()(*i3b892eb54cedbd9cc555b9f1a7958d6152b7730fa895edaed6eaa7e22b3c15ca.OverridesRequestBuilder) {
    return i3b892eb54cedbd9cc555b9f1a7958d6152b7730fa895edaed6eaa7e22b3c15ca.NewOverridesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OverridesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.inferenceClassification.overrides.item collection
func (m *InferenceClassificationRequestBuilder) OverridesById(id string)(*ie9af6222ea455c01969cfed303f450bfed10eddca9c64da588590323f4684264.InferenceClassificationOverrideItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["inferenceClassificationOverride%2Did"] = id
    }
    return ie9af6222ea455c01969cfed303f450bfed10eddca9c64da588590323f4684264.NewInferenceClassificationOverrideItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) Patch(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable, requestConfiguration *InferenceClassificationRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
func (m *InferenceClassificationRequestBuilder) PatchWithResponseHandler(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.InferenceClassificationable, requestConfiguration *InferenceClassificationRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
