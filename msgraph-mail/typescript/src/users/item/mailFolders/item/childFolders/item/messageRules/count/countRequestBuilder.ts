// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { type ODataError } from '../../../../../../../../models/oDataErrors/';
import { createODataErrorFromDiscriminatorValue, deserializeIntoODataError, serializeODataError } from '../../../../../../../../models/oDataErrors/oDataError';
import { BaseRequestBuilder, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

export interface CountRequestBuilderGetQueryParameters {
    /**
     * Filter items by property values
     */
    filter?: string;
}
export interface CountRequestBuilderGetRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]>;
    /**
     * Request options
     */
    options?: RequestOption[];
    /**
     * Request query parameters
     */
    queryParameters?: CountRequestBuilderGetQueryParameters;
}
/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1}/messageRules/$count
 */
export class CountRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new CountRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules/$count{?%24filter}");
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
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendPrimitiveAsync<number>(requestInfo, "number", errorMapping);
    };
    /**
     * Get the number of the resource
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: CountRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters, countRequestBuilderGetQueryParametersMapper);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.tryAddRequestHeaders("Accept", "text/plain;q=0.9");
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
const countRequestBuilderGetQueryParametersMapper: Record<string, string> = {
    "filter": "%24filter",
};
// tslint:enable
// eslint-enable
