import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {ReadableStream} from 'web-streams-polyfill/es2018';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/$value  */
export class ContentRequestBuilder {
    /** Current path for the request  */
    private _currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    private _httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly _pathSegment: string;
    /** Factory to use to get a serializer for payload serialization  */
    private _serializerFactory?: SerializationWriterFactory | undefined;
    /**
     * Instantiates a new ContentRequestBuilder and sets the default values.
     */
    public constructor() {
        this._pathSegment = "/$value";
    };
    /**
     * Get media content for the navigation property messages from users
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createGetRequestInfo(h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        return requestInfo;
    };
    /**
     * Update media content for the navigation property messages in users
     * @param body Binary request body
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createPutRequestInfo(body: ReadableStream | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.PUT,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setStreamContent(body);
        return requestInfo;
    };
    /**
     * Get media content for the navigation property messages from users
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of ReadableStream
     */
    public get(h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ReadableStream | undefined> {
        const requestInfo = this.createGetRequestInfo(
            h
        );
        return this.httpCore?.sendPrimitiveAsync<ReadableStream>(requestInfo, "ReadableStream", responseHandler) ?? Promise.reject(new Error('http core is null'));
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
     * Update media content for the navigation property messages in users
     * @param body Binary request body
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public put(body: ReadableStream | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        const requestInfo = this.createPutRequestInfo(
            body, h
        );
        return this.httpCore?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
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
