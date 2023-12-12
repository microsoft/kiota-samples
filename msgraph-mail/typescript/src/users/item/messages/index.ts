// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { createMessageCollectionResponseFromDiscriminatorValue, createMessageFromDiscriminatorValue, deserializeIntoMessage, serializeMessage, type Message, type MessageCollectionResponse } from '../../../models/';
import { createODataErrorFromDiscriminatorValue, deserializeIntoODataError, serializeODataError, type ODataError } from '../../../models/oDataErrors/';
import { CountRequestBuilderRequestsMetadata, CountRequestBuilderUriTemplate, type CountRequestBuilder } from './count/';
import { MessageItemRequestBuilderNavigationMetadata, MessageItemRequestBuilderRequestsMetadata, MessageItemRequestBuilderUriTemplate, type MessageItemRequestBuilder } from './item/';
import { RequestInformation, type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata, type Parsable, type ParsableFactory, type RequestAdapter, type RequestConfiguration, type RequestMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users/{user-id}/messages
 */
export interface MessagesRequestBuilder extends BaseRequestBuilder<MessagesRequestBuilder> {
    /**
     * The Count property
     */
    get count(): CountRequestBuilder;
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item collection
     * @param messageId The unique identifier of message
     * @returns a MessageItemRequestBuilder
     */
     byMessageId(messageId: string) : MessageItemRequestBuilder;
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MessageCollectionResponse
     * @see {@link https://learn.microsoft.com/graph/api/opentypeextension-get?view=graph-rest-1.0|Find more info here}
     */
     get(requestConfiguration?: RequestConfiguration<MessagesRequestBuilderGetQueryParameters> | undefined) : Promise<MessageCollectionResponse | undefined>;
    /**
     * Create a draft of a new message in either JSON or MIME format. When using JSON format, you can:- Include an attachment to the message.- Update the draft later to add content to the body or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in the request body.- /* Add any attachments and S/MIME properties to the MIME content. By default, this operation saves the draft in the Drafts folder. Send the draft message in a subsequent operation. Alternatively, send a new message in a single operation, or create a draft to forward, reply and reply-all to an existing message.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Message
     * @see {@link https://learn.microsoft.com/graph/api/user-post-messages?view=graph-rest-1.0|Find more info here}
     */
     post(body: Message, requestConfiguration?: RequestConfiguration<object> | undefined) : Promise<Message | undefined>;
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<MessagesRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
    /**
     * Create a draft of a new message in either JSON or MIME format. When using JSON format, you can:- Include an attachment to the message.- Update the draft later to add content to the body or change other message properties. When using MIME format:- Provide the applicable Internet message headers and the MIME content, all encoded in base64 format in the request body.- /* Add any attachments and S/MIME properties to the MIME content. By default, this operation saves the draft in the Drafts folder. Send the draft message in a subsequent operation. Alternatively, send a new message in a single operation, or create a draft to forward, reply and reply-all to an existing message.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
     toPostRequestInformation(body: Message, requestConfiguration?: RequestConfiguration<object> | undefined) : RequestInformation;
}
/**
 * The messages in a mailbox or folder. Read-only. Nullable.
 */
export interface MessagesRequestBuilderGetQueryParameters {
    /**
     * Include count of items
     */
    count?: boolean;
    /**
     * Expand related entities
     */
    expand?: string[];
    /**
     * Filter items by property values
     */
    filter?: string;
    /**
     * Include Hidden Messages
     */
    includeHiddenMessages?: string;
    /**
     * Order items by property values
     */
    orderby?: string[];
    /**
     * Search items by search phrases
     */
    search?: string;
    /**
     * Select properties to be returned
     */
    select?: string[];
    /**
     * Skip the first n items
     */
    skip?: number;
    /**
     * Show only the first n items
     */
    top?: number;
}
const MessagesRequestBuilderGetQueryParametersMapper: Record<string, string> = {
    "count": "%24count",
    "expand": "%24expand",
    "filter": "%24filter",
    "orderby": "%24orderby",
    "search": "%24search",
    "select": "%24select",
    "skip": "%24skip",
    "top": "%24top",
};
export const MessagesRequestBuilderNavigationMetadata: Record<Exclude<keyof MessagesRequestBuilder, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    "byMessageId": {
        uriTemplate: MessageItemRequestBuilderUriTemplate,
        requestsMetadata: MessageItemRequestBuilderRequestsMetadata,
        navigationMetadata: MessageItemRequestBuilderNavigationMetadata,
        pathParametersMappings: ["message%2Did"],
    },
    "count": {
        uriTemplate: CountRequestBuilderUriTemplate,
        requestsMetadata: CountRequestBuilderRequestsMetadata,
    },
};
export const MessagesRequestBuilderRequestsMetadata: Record<string, RequestMetadata> = {
    "get": {
        errorMappings: {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>,
        adapterMethodName: "sendAsync",
        responseBodyFactory:  createMessageCollectionResponseFromDiscriminatorValue,
        queryParametersMapper: MessagesRequestBuilderGetQueryParametersMapper,
    },
    "post": {
        errorMappings: {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>,
        adapterMethodName: "sendAsync",
        responseBodyFactory:  createMessageFromDiscriminatorValue,
        requestBodyContentType: "application/json",
        requestBodySerializer: serializeMessage,
    },
};
export const MessagesRequestBuilderUriTemplate = "{+baseurl}/users/{user%2Did}/messages{?includeHiddenMessages,%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
// tslint:enable
// eslint-enable
