import { type MessageRuleCollectionResponse } from '../../../../../../../models/';
import { createMessageRuleFromDiscriminatorValue, deserializeIntoMessageRule, serializeMessageRule, type MessageRule } from '../../../../../../../models/messageRule';
import { createMessageRuleCollectionResponseFromDiscriminatorValue } from '../../../../../../../models/messageRuleCollectionResponse';
import { CountRequestBuilder } from './count/countRequestBuilder';
import { MessageRuleItemRequestBuilder } from './item/messageRuleItemRequestBuilder';
import { BaseRequestBuilder, getPathParameters, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestOption } from '@microsoft/kiota-abstractions';

export interface MessageRulesRequestBuilderGetQueryParameters {
    /**
     * Include count of items
     */
    count?: boolean | undefined;
    /**
     * Filter items by property values
     */
    filter?: string | undefined;
    /**
     * Order items by property values
     */
    orderby?: string[] | undefined;
    /**
     * Select properties to be returned
     */
    select?: string[] | undefined;
    /**
     * Skip the first n items
     */
    skip?: number | undefined;
    /**
     * Show only the first n items
     */
    top?: number | undefined;
}
export interface MessageRulesRequestBuilderGetRequestConfiguration {
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
    queryParameters?: MessageRulesRequestBuilderGetQueryParameters | undefined;
}
export interface MessageRulesRequestBuilderPostRequestConfiguration {
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
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1}/messageRules
 */
export class MessageRulesRequestBuilder extends BaseRequestBuilder {
    /**
     * The Count property
     */
    public get count(): CountRequestBuilder {
        return new CountRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item.childFolders.item.messageRules.item collection
     * @param messageRuleId The unique identifier of messageRule
     * @returns a MessageRuleItemRequestBuilder
     */
    public byMessageRuleId(messageRuleId: string) : MessageRuleItemRequestBuilder {
        if(!messageRuleId) throw new Error("messageRuleId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["messageRule%2Did"] = messageRuleId
        return new MessageRuleItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Instantiates a new MessageRulesRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}");
    };
    /**
     * Get all the messageRule objects defined for the user's inbox.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MessageRuleCollectionResponse
     * @see {@link https://learn.microsoft.com/graph/api/mailfolder-list-messagerules?view=graph-rest-1.0|Find more info here}
     */
    public get(requestConfiguration?: MessageRulesRequestBuilderGetRequestConfiguration | undefined) : Promise<MessageRuleCollectionResponse | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter.sendAsync<MessageRuleCollectionResponse>(requestInfo, createMessageRuleCollectionResponseFromDiscriminatorValue, undefined);
    };
    /**
     * Create a messageRule object by specifying a set of conditions and actions. Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MessageRule
     * @see {@link https://learn.microsoft.com/graph/api/mailfolder-post-messagerules?view=graph-rest-1.0|Find more info here}
     */
    public post(body: MessageRule, requestConfiguration?: MessageRulesRequestBuilderPostRequestConfiguration | undefined) : Promise<MessageRule | undefined> {
        const requestInfo = this.toPostRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter.sendAsync<MessageRule>(requestInfo, createMessageRuleFromDiscriminatorValue, undefined);
    };
    /**
     * Get all the messageRule objects defined for the user's inbox.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: MessageRulesRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
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
     * Create a messageRule object by specifying a set of conditions and actions. Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPostRequestInformation(body: MessageRule, requestConfiguration?: MessageRulesRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
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
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializeMessageRule);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a messageRulesRequestBuilder
     */
    public withUrl(rawUrl: string) : MessageRulesRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new MessageRulesRequestBuilder(rawUrl, this.requestAdapter);
    };
}
