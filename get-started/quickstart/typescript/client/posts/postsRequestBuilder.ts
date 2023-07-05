import {createPostFromDiscriminatorValue} from '../models/createPostFromDiscriminatorValue';
import {deserializeIntoPost} from '../models/deserializeIntoPost';
import {Post} from '../models/post';
import {serializePost} from '../models/serializePost';
import {PostItemRequestBuilder} from './item/postItemRequestBuilder';
import {PostsRequestBuilderGetRequestConfiguration} from './postsRequestBuilderGetRequestConfiguration';
import {PostsRequestBuilderPostRequestConfiguration} from './postsRequestBuilderPostRequestConfiguration';
import {BaseRequestBuilder, getPathParameters, HttpMethod, Parsable, ParsableFactory, RequestAdapter, RequestInformation, RequestOption} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /posts
 */
export class PostsRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the ApiSdk.posts.item collection
     * @param postId Unique identifier of the item
     * @returns a PostItemRequestBuilder
     */
    public byPostId(postId: string) : PostItemRequestBuilder {
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
    public post(body: Post | undefined, requestConfiguration?: PostsRequestBuilderPostRequestConfiguration | undefined) : Promise<Post | undefined> {
        if(!body) throw new Error("body cannot be undefined");
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
    public toPostRequestInformation(body: Post | undefined, requestConfiguration?: PostsRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
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
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body as any, serializePost);
        return requestInfo;
    };
}
