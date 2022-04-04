import {InferenceClassificationOverride} from '../../../../models/microsoft/graph/';
import {createInferenceClassificationOverrideFromDiscriminatorValue} from '../../../../models/microsoft/graph/createInferenceClassificationOverrideFromDiscriminatorValue';
import {createOverridesResponseFromDiscriminatorValue} from './createOverridesResponseFromDiscriminatorValue';
import {OverridesResponse} from './index';
import {getPathParameters, HttpMethod, Parsable, ParsableFactory, RequestAdapter, RequestInformation, RequestOption, ResponseHandler} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}/inferenceClassification/overrides  */
export class OverridesRequestBuilder {
    /** Path parameters for the request  */
    private readonly pathParameters: Record<string, unknown>;
    /** The request adapter to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private readonly urlTemplate: string;
    /**
     * Instantiates a new OverridesRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        if(!pathParameters) throw new Error("pathParameters cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.urlTemplate = "{+baseurl}/users/{user_id}/inferenceClassification/overrides{?top,skip,filter,count,orderby,select}";
        const urlTplParams = getPathParameters(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    };
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @returns a RequestInformation
     */
    public createGetRequestInformation(queryParameters?: {
                    count?: boolean,
                    filter?: string,
                    orderby?: string[],
                    select?: string[],
                    skip?: number,
                    top?: number
                    } | undefined, headers?: Record<string, string> | undefined, options?: RequestOption[] | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        if(headers) requestInfo.headers = headers;
        queryParameters && requestInfo.setQueryStringParametersFromRawObject(queryParameters);
        options && requestInfo.addRequestOptions(...options);
        return requestInfo;
    };
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @returns a RequestInformation
     */
    public createPostRequestInformation(body: InferenceClassificationOverride | undefined, headers?: Record<string, string> | undefined, options?: RequestOption[] | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.POST;
        if(headers) requestInfo.headers = headers;
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body);
        options && requestInfo.addRequestOptions(...options);
        return requestInfo;
    };
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of OverridesResponse
     */
    public get(queryParameters?: {
                    count?: boolean,
                    filter?: string,
                    orderby?: string[],
                    select?: string[],
                    skip?: number,
                    top?: number
                    } | undefined, headers?: Record<string, string> | undefined, options?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<OverridesResponse | undefined> {
        const requestInfo = this.createGetRequestInformation(
            queryParameters, headers, options
        );
        return this.requestAdapter?.sendAsync<OverridesResponse>(requestInfo, createOverridesResponseFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of InferenceClassificationOverride
     */
    public post(body: InferenceClassificationOverride | undefined, headers?: Record<string, string> | undefined, options?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassificationOverride | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPostRequestInformation(
            body, headers, options
        );
        return this.requestAdapter?.sendAsync<InferenceClassificationOverride>(requestInfo, createInferenceClassificationOverrideFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('http core is null'));
    };
}
