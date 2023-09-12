import { createPostFromDiscriminatorValue } from '../models/createPostFromDiscriminatorValue';
import { deserializeIntoPost } from '../models/deserializeIntoPost';
import { type Post } from '../models/post';
import { serializePost } from '../models/serializePost';
import { PostItemRequestBuilder } from './item/postItemRequestBuilder';
import { type PostsRequestBuilderGetRequestConfiguration } from './postsRequestBuilderGetRequestConfiguration';
import { type PostsRequestBuilderPostRequestConfiguration } from './postsRequestBuilderPostRequestConfiguration';
import { BaseRequestBuilder, getPathParameters, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /posts
 */
export class PostsRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the ApiSdk.posts.item collection
     * @param postId key: id of post
     * @returns a PostItemRequestBuilder
     */
    public byPostId(postId: number) : PostItemRequestBuilder {
        if(!postId) throw new Error("postId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["post%2Did"] = postId
        return new PostItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Instantiates a new PostsRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/posts{?userId*,title*}");
    };
    /**
     * Get posts
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Post
     */
    public get(requestConfiguration?: PostsRequestBuilderGetRequestConfiguration | undefined) : Promise<Post[] | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter.sendCollectionAsync<Post>(requestInfo, createPostFromDiscriminatorValue, undefined);
    };
    /**
     * Create post
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Post
     */
    public post(body: Post, requestConfiguration?: PostsRequestBuilderPostRequestConfiguration | undefined) : Promise<Post | undefined> {
        const requestInfo = this.toPostRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter.sendAsync<Post>(requestInfo, createPostFromDiscriminatorValue, undefined);
    };
    /**
     * Get posts
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: PostsRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * Create post
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPostRequestInformation(body: Post, requestConfiguration?: PostsRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializePost);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a postsRequestBuilder
     */
    public withUrl(rawUrl: string) : PostsRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new PostsRequestBuilder(rawUrl, this.requestAdapter);
    };
}
