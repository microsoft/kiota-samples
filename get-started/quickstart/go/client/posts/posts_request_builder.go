package posts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620 "kiota_posts/client/models"
)

// PostsRequestBuilder builds and executes requests for operations under \posts
type PostsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PostsRequestBuilderGetQueryParameters get posts
type PostsRequestBuilderGetQueryParameters struct {
    // Filter results by title
    Title *string
    // Filter results by user ID
    UserId *int32
}
// PostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PostsRequestBuilderGetQueryParameters
}
// PostsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPostsRequestBuilderInternal instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    m := &PostsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/posts{?userId*,title*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewPostsRequestBuilder instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get posts
func (m *PostsRequestBuilder) Get(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)([]i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollection(ctx, requestInfo, i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.CreatePostFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, len(res))
    for i, v := range res {
        val[i] = v.(i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable)
    }
    return val, nil
}
// Post create post
func (m *PostsRequestBuilder) Post(ctx context.Context, body i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)(i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.CreatePostFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable), nil
}
// ToGetRequestInformation get posts
func (m *PostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create post
func (m *PostsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i5dbc5a8abf7315a9f71dbdc1d57fa9606d698daab75f49bdbbd6d7a719d6e620.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
