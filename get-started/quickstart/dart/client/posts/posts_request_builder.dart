// ignore_for_file: type=lint
import 'package:microsoft_kiota_abstractions/microsoft_kiota_abstractions.dart';
import '../models/post.dart';
import './item/post_item_request_builder.dart';
import './posts_request_builder_get_query_parameters.dart';

/// auto generated
/// Builds and executes requests for operations under \posts
class PostsRequestBuilder extends BaseRequestBuilder<PostsRequestBuilder> {
    /// Gets an item from the ApiSdk.posts.item collection
    ///  [postId] key: id of post
    PostItemRequestBuilder byPostId(int postId) {
        var urlTplParams = Map.of(pathParameters);
        urlTplParams.putIfAbsent('post%2Did', () => postId);
        return PostItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /// Clones the requestbuilder.
    @override
    PostsRequestBuilder clone() {
        return PostsRequestBuilder(pathParameters, requestAdapter);
    }
    /// Instantiates a new [PostsRequestBuilder] and sets the default values.
    ///  [pathParameters] Path parameters for the request
    ///  [requestAdapter] The request adapter to use to execute the requests.
    PostsRequestBuilder(Map<String, dynamic> pathParameters, RequestAdapter requestAdapter) : super(requestAdapter, "{+baseurl}/posts{?title*,userId*}", pathParameters) ;
    /// Instantiates a new [PostsRequestBuilder] and sets the default values.
    ///  [rawUrl] The raw URL to use for the request builder.
    ///  [requestAdapter] The request adapter to use to execute the requests.
    PostsRequestBuilder.withUrl(String rawUrl, RequestAdapter requestAdapter) : super(requestAdapter, "{+baseurl}/posts{?title*,userId*}", {RequestInformation.rawUrlKey : rawUrl}) ;
    /// Get posts
    ///  [requestConfiguration] Configuration for the request such as headers, query parameters, and middleware options.
    Future<Iterable<Post>?> getAsync([void Function(RequestConfiguration<PostsRequestBuilderGetQueryParameters>)? requestConfiguration]) async {
        var requestInfo = toGetRequestInformation(requestConfiguration);
        return await requestAdapter.sendCollection<Post>(requestInfo, Post.createFromDiscriminatorValue, {});
    }
    /// Create post
    ///  [body] The request body
    ///  [requestConfiguration] Configuration for the request such as headers, query parameters, and middleware options.
    Future<Post?> postAsync(Post body, [void Function(RequestConfiguration<DefaultQueryParameters>)? requestConfiguration]) async {
        var requestInfo = toPostRequestInformation(body, requestConfiguration);
        return await requestAdapter.send<Post>(requestInfo, Post.createFromDiscriminatorValue, {});
    }
    /// Get posts
    ///  [requestConfiguration] Configuration for the request such as headers, query parameters, and middleware options.
    RequestInformation toGetRequestInformation([void Function(RequestConfiguration<PostsRequestBuilderGetQueryParameters>)? requestConfiguration]) {
        var requestInfo = RequestInformation(httpMethod : HttpMethod.get, urlTemplate : urlTemplate, pathParameters :  pathParameters);
        requestInfo.configure<PostsRequestBuilderGetQueryParameters>(requestConfiguration, () => PostsRequestBuilderGetQueryParameters());
        requestInfo.headers.put('Accept', 'application/json');
        return requestInfo;
    }
    /// Create post
    ///  [body] The request body
    ///  [requestConfiguration] Configuration for the request such as headers, query parameters, and middleware options.
    RequestInformation toPostRequestInformation(Post body, [void Function(RequestConfiguration<DefaultQueryParameters>)? requestConfiguration]) {
        var requestInfo = RequestInformation(httpMethod : HttpMethod.post, urlTemplate : urlTemplate, pathParameters :  pathParameters);
        requestInfo.configure<DefaultQueryParameters>(requestConfiguration, () => DefaultQueryParameters());
        requestInfo.headers.put('Accept', 'application/json');
        requestInfo.setContentFromParsable(requestAdapter, 'application/json', body);
        return requestInfo;
    }
}
