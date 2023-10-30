// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { createMessageRuleFromDiscriminatorValue, deserializeIntoMessageRule, serializeMessageRule, type MessageRule } from '../../../../../../../../models/messageRule';
import { type ODataError } from '../../../../../../../../models/oDataErrors/';
import { createODataErrorFromDiscriminatorValue, deserializeIntoODataError, serializeODataError } from '../../../../../../../../models/oDataErrors/oDataError';
import { BaseRequestBuilder, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

export interface MessageRuleItemRequestBuilderDeleteRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]>;
    /**
     * Request options
     */
    options?: RequestOption[];
}
export interface MessageRuleItemRequestBuilderGetQueryParameters {
    /**
     * Select properties to be returned
     */
    select?: string[];
}
export interface MessageRuleItemRequestBuilderGetRequestConfiguration {
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
    queryParameters?: MessageRuleItemRequestBuilderGetQueryParameters;
}
export interface MessageRuleItemRequestBuilderPatchRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]>;
    /**
     * Request options
     */
    options?: RequestOption[];
}
/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1}/messageRules/{messageRule-id}
 */
export class MessageRuleItemRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new MessageRuleItemRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules/{messageRule%2Did}{?%24select}");
    };
    /**
     * Delete the specified messageRule object. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of ArrayBuffer
     * @see {@link https://learn.microsoft.com/graph/api/messagerule-delete?view=graph-rest-1.0|Find more info here}
     */
    public delete(requestConfiguration?: MessageRuleItemRequestBuilderDeleteRequestConfiguration | undefined) : Promise<ArrayBuffer | undefined> {
        const requestInfo = this.toDeleteRequestInformation(
            requestConfiguration
        );
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendPrimitiveAsync<ArrayBuffer>(requestInfo, "ArrayBuffer", errorMapping);
    };
    /**
     * Get the properties and relationships of a messageRule object. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MessageRule
     * @see {@link https://learn.microsoft.com/graph/api/messagerule-get?view=graph-rest-1.0|Find more info here}
     */
    public get(requestConfiguration?: MessageRuleItemRequestBuilderGetRequestConfiguration | undefined) : Promise<MessageRule | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendAsync<MessageRule>(requestInfo, createMessageRuleFromDiscriminatorValue, errorMapping);
    };
    /**
     * Change writable properties on a messageRule object and save the changes. This API is available in the following national cloud deployments.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MessageRule
     * @see {@link https://learn.microsoft.com/graph/api/messagerule-update?view=graph-rest-1.0|Find more info here}
     */
    public patch(body: MessageRule, requestConfiguration?: MessageRuleItemRequestBuilderPatchRequestConfiguration | undefined) : Promise<MessageRule | undefined> {
        const requestInfo = this.toPatchRequestInformation(
            body, requestConfiguration
        );
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendAsync<MessageRule>(requestInfo, createMessageRuleFromDiscriminatorValue, errorMapping);
    };
    /**
     * Delete the specified messageRule object. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toDeleteRequestInformation(requestConfiguration?: MessageRuleItemRequestBuilderDeleteRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.DELETE;
        requestInfo.tryAddRequestHeaders("Accept", "application/json");
        return requestInfo;
    };
    /**
     * Get the properties and relationships of a messageRule object. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: MessageRuleItemRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.tryAddRequestHeaders("Accept", "application/json;q=1");
        return requestInfo;
    };
    /**
     * Change writable properties on a messageRule object and save the changes. This API is available in the following national cloud deployments.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPatchRequestInformation(body: MessageRule, requestConfiguration?: MessageRuleItemRequestBuilderPatchRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.PATCH;
        requestInfo.tryAddRequestHeaders("Accept", "application/json;q=1");
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializeMessageRule);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a MessageRuleItemRequestBuilder
     */
    public withUrl(rawUrl: string) : MessageRuleItemRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new MessageRuleItemRequestBuilder(rawUrl, this.requestAdapter);
    };
}
// tslint:enable
// eslint-enable
