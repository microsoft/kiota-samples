import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {Message} from '../../message';
import {MessagesResponse} from './messagesResponse';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages  */
export class MessagesRequestBuilder {
    /** Current path for the request  */
    private _currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    private _httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly _pathSegment: string;
    /** Factory to use to get a serializer for payload serialization  */
    private _serializerFactory?: SerializationWriterFactory | undefined;
    /**
     * Instantiates a new MessagesRequestBuilder and sets the default values.
     */
    public constructor() {
        this._pathSegment = "/messages";
    };
    /**
     * Get messages from users
     * @param h Request headers
     * @param q Request query parameters
     * @returns a RequestInfo
     */
    public createGetRequestInfo(q?: {
                    count?: boolean,
                    expand?: string[],
                    filter?: string,
                    orderby?: string[],
                    search?: string,
                    select?: string[],
                    skip?: number,
                    top?: number
                    } | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        q && requestInfo.setQueryStringParametersFromRawObject(q);
        return requestInfo;
    };
    /**
     * Create new navigation property to messages for users
     * @param body 
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createPostRequestInfo(body: Message | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.POST,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setContentFromParsable(body, this.serializerFactory, "application/json");
        return requestInfo;
    };
    /**
     * Get messages from users
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MessagesResponse
     */
    public get(q?: {
                    count?: boolean,
                    expand?: string[],
                    filter?: string,
                    orderby?: string[],
                    search?: string,
                    select?: string[],
                    skip?: number,
                    top?: number
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MessagesResponse | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<MessagesResponse>(requestInfo, MessagesResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Gets the currentPath property value. Current path for the request
     * @returns a string
     */
    public get currentPath() {
        return this._currentPath;
    };
    /**
     * Gets the httpCore property value. Core service to use to execute the requests
     * @returns a HttpCore
     */
    public get httpCore() {
        return this._httpCore;
    };
    /**
     * Gets the pathSegment property value. Path segment to use to build the URL for the current request builder
     * @returns a string
     */
    public get pathSegment() {
        return this._pathSegment;
    };
    /**
     * Gets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @returns a SerializationWriterFactory
     */
    public get serializerFactory() {
        return this._serializerFactory;
    };
    /**
     * Create new navigation property to messages for users
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Message
     */
    public post(body: Message | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Message | undefined> {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<Message>(requestInfo, Message, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Sets the currentPath property value. Current path for the request
     * @param value Value to set for the currentPath property.
     */
    public set currentPath(value: string | undefined) {
        this._currentPath = value;
    };
    /**
     * Sets the httpCore property value. Core service to use to execute the requests
     * @param value Value to set for the httpCore property.
     */
    public set httpCore(value: HttpCore | undefined) {
        this._httpCore = value;
    };
    /**
     * Sets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @param value Value to set for the serializerFactory property.
     */
    public set serializerFactory(value: SerializationWriterFactory | undefined) {
        this._serializerFactory = value;
    };
}
