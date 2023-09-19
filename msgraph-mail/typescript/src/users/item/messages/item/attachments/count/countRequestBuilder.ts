import { BaseRequestBuilder, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

export interface CountRequestBuilderGetQueryParameters {
    /**
     * Filter items by property values
     */
    filter?: string | undefined;
}
export interface CountRequestBuilderGetRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]> | undefined;
    /**
     * Request options
     */
    options?: RequestOption[] | undefined;
    /**
     * Request query parameters
     */
    queryParameters?: CountRequestBuilderGetQueryParameters | undefined;
}
/**
 * Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/attachments/$count
 */
export class CountRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new CountRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/attachments/$count{?%24filter}");
    };
    /**
     * Get the number of the resource
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Integer
     */
    public get(requestConfiguration?: CountRequestBuilderGetRequestConfiguration | undefined) : Promise<number | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter.sendPrimitiveAsync<number>(requestInfo, "number", undefined);
    };
    /**
     * Get the number of the resource
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: CountRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = ["text/plain"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a CountRequestBuilder
     */
    public withUrl(rawUrl: string) : CountRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new CountRequestBuilder(rawUrl, this.requestAdapter);
    };
}
