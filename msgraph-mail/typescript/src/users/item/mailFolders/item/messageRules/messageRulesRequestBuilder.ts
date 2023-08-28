import {MessageRuleCollectionResponse} from '../../../../../models/';
import {createMessageRuleCollectionResponseFromDiscriminatorValue} from '../../../../../models/createMessageRuleCollectionResponseFromDiscriminatorValue';
import {createMessageRuleFromDiscriminatorValue} from '../../../../../models/createMessageRuleFromDiscriminatorValue';
import {deserializeIntoMessageRule} from '../../../../../models/deserializeIntoMessageRule';
import type {MessageRule} from '../../../../../models/messageRule';
import {serializeMessageRule} from '../../../../../models/serializeMessageRule';
import {CountRequestBuilder} from './count/countRequestBuilder';
import {MessageRuleItemRequestBuilder} from './item/messageRuleItemRequestBuilder';
import {MessageRulesRequestBuilderGetRequestConfiguration} from './messageRulesRequestBuilderGetRequestConfiguration';
import {MessageRulesRequestBuilderPostRequestConfiguration} from './messageRulesRequestBuilderPostRequestConfiguration';
import {BaseRequestBuilder, HttpMethod, RequestInformation, getPathParameters} from '@microsoft/kiota-abstractions';
import type {Parsable, ParsableFactory, RequestAdapter, RequestOption} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messageRules
 */
export class MessageRulesRequestBuilder extends BaseRequestBuilder {
    /**
     * The Count property
     */
    public get count(): CountRequestBuilder {
        return new CountRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item.messageRules.item collection
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
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}");
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
     * Create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
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
     * Create a messageRule object by specifying a set of conditions and actions.  Outlook carries out those actions if an incoming message in the user's Inbox meets the specified conditions.
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
