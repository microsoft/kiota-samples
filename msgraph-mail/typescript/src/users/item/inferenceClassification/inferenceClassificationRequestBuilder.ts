import { createInferenceClassificationFromDiscriminatorValue, deserializeIntoInferenceClassification, serializeInferenceClassification, type InferenceClassification } from '../../../models/inferenceClassification';
import { OverridesRequestBuilder } from './overrides/overridesRequestBuilder';
import { BaseRequestBuilder, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

export interface InferenceClassificationRequestBuilderGetQueryParameters {
    /**
     * Select properties to be returned
     */
    select?: string[] | undefined;
}
export interface InferenceClassificationRequestBuilderGetRequestConfiguration {
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
    queryParameters?: InferenceClassificationRequestBuilderGetQueryParameters | undefined;
}
export interface InferenceClassificationRequestBuilderPatchRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]> | undefined;
    /**
     * Request options
     */
    options?: RequestOption[] | undefined;
}
/**
 * Builds and executes requests for operations under /users/{user-id}/inferenceClassification
 */
export class InferenceClassificationRequestBuilder extends BaseRequestBuilder {
    /**
     * The overrides property
     */
    public get overrides(): OverridesRequestBuilder {
        return new OverridesRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24select}");
    };
    /**
     * Relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of InferenceClassification
     */
    public get(requestConfiguration?: InferenceClassificationRequestBuilderGetRequestConfiguration | undefined) : Promise<InferenceClassification | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter.sendAsync<InferenceClassification>(requestInfo, createInferenceClassificationFromDiscriminatorValue, undefined);
    };
    /**
     * Update the navigation property inferenceClassification in users
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of InferenceClassification
     */
    public patch(body: InferenceClassification, requestConfiguration?: InferenceClassificationRequestBuilderPatchRequestConfiguration | undefined) : Promise<InferenceClassification | undefined> {
        const requestInfo = this.toPatchRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter.sendAsync<InferenceClassification>(requestInfo, createInferenceClassificationFromDiscriminatorValue, undefined);
    };
    /**
     * Relevance classification of the user's messages based on explicit designations that override inferred relevance or importance.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: InferenceClassificationRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
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
     * Update the navigation property inferenceClassification in users
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPatchRequestInformation(body: InferenceClassification, requestConfiguration?: InferenceClassificationRequestBuilderPatchRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.PATCH;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializeInferenceClassification);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a inferenceClassificationRequestBuilder
     */
    public withUrl(rawUrl: string) : InferenceClassificationRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new InferenceClassificationRequestBuilder(rawUrl, this.requestAdapter);
    };
}
