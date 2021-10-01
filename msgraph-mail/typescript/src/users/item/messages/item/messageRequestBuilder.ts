import {Message} from '../../../../models/microsoft/graph/message';
import {AttachmentsRequestBuilder} from './attachments/attachmentsRequestBuilder';
import {AttachmentRequestBuilder} from './attachments/item/attachmentRequestBuilder';
import {ExtensionsRequestBuilder} from './extensions/extensionsRequestBuilder';
import {ExtensionRequestBuilder} from './extensions/item/extensionRequestBuilder';
import {MultiValueLegacyExtendedPropertyRequestBuilder} from './multiValueExtendedProperties/item/multiValueLegacyExtendedPropertyRequestBuilder';
import {MultiValueExtendedPropertiesRequestBuilder} from './multiValueExtendedProperties/multiValueExtendedPropertiesRequestBuilder';
import {SingleValueLegacyExtendedPropertyRequestBuilder} from './singleValueExtendedProperties/item/singleValueLegacyExtendedPropertyRequestBuilder';
import {SingleValueExtendedPropertiesRequestBuilder} from './singleValueExtendedProperties/singleValueExtendedPropertiesRequestBuilder';
import {ContentRequestBuilder} from './value/contentRequestBuilder';
import {ResponseHandler, Parsable, HttpMethod, RequestInformation, RequestOption, RequestAdapter} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}  */
export class MessageRequestBuilder {
    public get attachments(): AttachmentsRequestBuilder {
        return new AttachmentsRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    public get content(): ContentRequestBuilder {
        return new ContentRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /** Current path for the request  */
    private readonly currentPath: string;
    public get extensions(): ExtensionsRequestBuilder {
        return new ExtensionsRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /** Whether the current path is a raw URL  */
    private readonly isRawUrl: boolean;
    public get multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /** The http core service to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    public get singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item.attachments.item collection
     * @param id Unique identifier of the item
     * @returns a attachmentRequestBuilder
     */
    public attachmentsById(id: String) : AttachmentRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new AttachmentRequestBuilder(this.currentPath + this.pathSegment + "/attachments/" + id, this.requestAdapter, false);
    };
    /**
     * Instantiates a new MessageRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param isRawUrl Whether the current path is a raw URL
     * @param requestAdapter The http core service to use to execute the requests.
     */
    public constructor(currentPath: string, requestAdapter: RequestAdapter, isRawUrl: boolean = true) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.pathSegment = "";
        this.requestAdapter = requestAdapter;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @returns a RequestInformation
     */
    public createDeleteRequestInformation(h?: object | undefined, o?: RequestOption[] | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.DELETE;
        h && requestInfo.setHeadersFromRawObject(h);
        o && requestInfo.addRequestOptions(...o);
        return requestInfo;
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @returns a RequestInformation
     */
    public createGetRequestInformation(q?: {
                    expand?: string[],
                    select?: string[]
                    } | undefined, h?: object | undefined, o?: RequestOption[] | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.GET;
        h && requestInfo.setHeadersFromRawObject(h);
        q && requestInfo.setQueryStringParametersFromRawObject(q);
        o && requestInfo.addRequestOptions(...o);
        return requestInfo;
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @returns a RequestInformation
     */
    public createPatchRequestInformation(body: Message | undefined, h?: object | undefined, o?: RequestOption[] | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.PATCH;
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body);
        o && requestInfo.addRequestOptions(...o);
        return requestInfo;
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public delete(h?: object | undefined, o?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        const requestInfo = this.createDeleteRequestInformation(
            h, o
        );
        return this.requestAdapter?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item.extensions.item collection
     * @param id Unique identifier of the item
     * @returns a extensionRequestBuilder
     */
    public extensionsById(id: String) : ExtensionRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new ExtensionRequestBuilder(this.currentPath + this.pathSegment + "/extensions/" + id, this.requestAdapter, false);
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Message
     */
    public get(q?: {
                    expand?: string[],
                    select?: string[]
                    } | undefined, h?: object | undefined, o?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Message | undefined> {
        const requestInfo = this.createGetRequestInformation(
            q, h, o
        );
        return this.requestAdapter?.sendAsync<Message>(requestInfo, Message, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item.multiValueExtendedProperties.item collection
     * @param id Unique identifier of the item
     * @returns a multiValueLegacyExtendedPropertyRequestBuilder
     */
    public multiValueExtendedPropertiesById(id: String) : MultiValueLegacyExtendedPropertyRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new MultiValueLegacyExtendedPropertyRequestBuilder(this.currentPath + this.pathSegment + "/multiValueExtendedProperties/" + id, this.requestAdapter, false);
    };
    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public patch(body: Message | undefined, h?: object | undefined, o?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPatchRequestInformation(
            body, h, o
        );
        return this.requestAdapter?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item.singleValueExtendedProperties.item collection
     * @param id Unique identifier of the item
     * @returns a singleValueLegacyExtendedPropertyRequestBuilder
     */
    public singleValueExtendedPropertiesById(id: String) : SingleValueLegacyExtendedPropertyRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new SingleValueLegacyExtendedPropertyRequestBuilder(this.currentPath + this.pathSegment + "/singleValueExtendedProperties/" + id, this.requestAdapter, false);
    };
}
