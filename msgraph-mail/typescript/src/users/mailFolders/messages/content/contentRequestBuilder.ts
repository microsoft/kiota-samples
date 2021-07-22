import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {ReadableStream} from 'web-streams-polyfill/es2018';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/$value  */
export class ContentRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    /** The http core service to use to execute the requests.  */
    private readonly httpCore: HttpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /**
     * Instantiates a new ContentRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     */
    public constructor(currentPath: string, httpCore: HttpCore) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!httpCore) throw new Error("httpCore cannot be undefined");
        this.pathSegment = "/$value";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
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
    public createPutRequestInfo(body: ReadableStream, h?: object | undefined) : RequestInfo {
        if(!body) throw new Error("body cannot be undefined");
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
     * Update media content for the navigation property messages in users
     * @param body Binary request body
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public put(body: ReadableStream, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPutRequestInfo(
            body, h
        );
        return this.httpCore?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
}
