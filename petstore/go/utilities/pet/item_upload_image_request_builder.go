package pet

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
)

// ItemUploadImageRequestBuilder builds and executes requests for operations under \pet\{petId}\uploadImage
type ItemUploadImageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemUploadImageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUploadImageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUploadImageRequestBuilderInternal instantiates a new UploadImageRequestBuilder and sets the default values.
func NewItemUploadImageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUploadImageRequestBuilder) {
    m := &ItemUploadImageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/pet/{petId}/uploadImage";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemUploadImageRequestBuilder instantiates a new UploadImageRequestBuilder and sets the default values.
func NewItemUploadImageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUploadImageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUploadImageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post uploads an image
func (m *ItemUploadImageRequestBuilder) Post(ctx context.Context, body []byte, requestConfiguration *ItemUploadImageRequestBuilderPostRequestConfiguration)(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.ApiResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.CreateApiResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.ApiResponseable), nil
}
func (m *ItemUploadImageRequestBuilder) ToPostRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemUploadImageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
