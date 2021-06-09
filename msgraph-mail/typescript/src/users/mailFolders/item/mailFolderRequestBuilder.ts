import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {MailFolder} from '../../mailFolder';
import {ChildFoldersRequestBuilder} from '../childFolders/childFoldersRequestBuilder';
import {MessageRuleRequestBuilder} from '../messageRules/item/messageRuleRequestBuilder';
import {MessageRulesRequestBuilder} from '../messageRules/messageRulesRequestBuilder';
import {MessageRequestBuilder} from '../messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../messages/messagesRequestBuilder';
import {MultiValueLegacyExtendedPropertyRequestBuilder} from '../multiValueExtendedProperties/item/multiValueLegacyExtendedPropertyRequestBuilder';
import {MultiValueExtendedPropertiesRequestBuilder} from '../multiValueExtendedProperties/multiValueExtendedPropertiesRequestBuilder';
import {SingleValueLegacyExtendedPropertyRequestBuilder} from '../singleValueExtendedProperties/item/singleValueLegacyExtendedPropertyRequestBuilder';
import {SingleValueExtendedPropertiesRequestBuilder} from '../singleValueExtendedProperties/singleValueExtendedPropertiesRequestBuilder';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}  */
export class MailFolderRequestBuilder {
    public get childFolders(): ChildFoldersRequestBuilder {
        const builder = new ChildFoldersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /** Current path for the request  */
    private _currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    private _httpCore?: HttpCore | undefined;
    public get messageRules(): MessageRulesRequestBuilder {
        const builder = new MessageRulesRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    public get messages(): MessagesRequestBuilder {
        const builder = new MessagesRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    public get multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        const builder = new MultiValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly _pathSegment: string = "";
    /** Factory to use to get a serializer for payload serialization  */
    private _serializerFactory?: SerializationWriterFactory | undefined;
    public get singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        const builder = new SingleValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders.childFolders collection
     * @param id Unique identifier of the item
     * @returns a MailFolderRequestBuilder
     */
    public childFoldersById (id: String) : MailFolderRequestBuilder {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/childFolders/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Delete navigation property mailFolders for users
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createDeleteRequestInfo (h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.DELETE,
        h && requestInfo.setHeadersFromRawObject(h);
        return requestInfo;
    };
    /**
     * Get mailFolders from users
     * @param h Request headers
     * @param q Request query parameters
     * @returns a RequestInfo
     */
    public createGetRequestInfo (q?: {
                    expand?: string[],
                    select?: string[]
                    } | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        q && requestInfo.setQueryStringParametersFromRawObject(q);
        return requestInfo;
    };
    /**
     * Update the navigation property mailFolders in users
     * @param body 
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createPatchRequestInfo (body: MailFolder | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.PATCH,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setContentFromParsable(body, this.serializerFactory, "application/json");
        return requestInfo;
    };
    /**
     * Delete navigation property mailFolders for users
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public delete (h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        const requestInfo = this.createDeleteRequestInfo(
            h
        );
        return this.httpCore?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Get mailFolders from users
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MailFolder
     */
    public get (q?: {
                    expand?: string[],
                    select?: string[]
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<MailFolder>(requestInfo, MailFolder, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Gets the currentPath property value. Current path for the request
     * @returns a string
     */
    public get currentPath () {
        return this._currentPath;
    };
    /**
     * Gets the httpCore property value. Core service to use to execute the requests
     * @returns a HttpCore
     */
    public get httpCore () {
        return this._httpCore;
    };
    /**
     * Gets the pathSegment property value. Path segment to use to build the URL for the current request builder
     * @returns a string
     */
    public get pathSegment () {
        return this._pathSegment;
    };
    /**
     * Gets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @returns a SerializationWriterFactory
     */
    public get serializerFactory () {
        return this._serializerFactory;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders.messageRules collection
     * @param id Unique identifier of the item
     * @returns a MessageRuleRequestBuilder
     */
    public messageRulesById (id: String) : MessageRuleRequestBuilder {
        const builder = new MessageRuleRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/messageRules/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders.messages collection
     * @param id Unique identifier of the item
     * @returns a MessageRequestBuilder
     */
    public messagesById (id: String) : MessageRequestBuilder {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/messages/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders.multiValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @returns a MultiValueLegacyExtendedPropertyRequestBuilder
     */
    public multiValueExtendedPropertiesById (id: String) : MultiValueLegacyExtendedPropertyRequestBuilder {
        const builder = new MultiValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/multiValueExtendedProperties/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Update the navigation property mailFolders in users
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public patch (body: MailFolder | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        const requestInfo = this.createPatchRequestInfo(
            body, h
        );
        return this.httpCore?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Sets the currentPath property value. Current path for the request
     * @param value Value to set for the currentPath property.
     */
    public set currentPath (value: string | undefined) {
        this._currentPath = value;
    };
    /**
     * Sets the httpCore property value. Core service to use to execute the requests
     * @param value Value to set for the httpCore property.
     */
    public set httpCore (value: HttpCore | undefined) {
        this._httpCore = value;
    };
    /**
     * Sets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @param value Value to set for the serializerFactory property.
     */
    public set serializerFactory (value: SerializationWriterFactory | undefined) {
        this._serializerFactory = value;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders.singleValueExtendedProperties collection
     * @param id Unique identifier of the item
     * @returns a SingleValueLegacyExtendedPropertyRequestBuilder
     */
    public singleValueExtendedPropertiesById (id: String) : SingleValueLegacyExtendedPropertyRequestBuilder {
        const builder = new SingleValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/singleValueExtendedProperties/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
}
