import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, MiddlewareOption} from '@microsoft/kiota-abstractions';
import {ReadableStream} from 'web-streams-polyfill/es2018';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/$value  */
export class ContentRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    /** The http core service to use to execute the requests.  */
    private readonly httpCore: HttpCore;
    /** Whether the current path is a raw URL  */
    private readonly isRawUrl: boolean;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /**
     * Instantiates a new ContentRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @param isRawUrl Whether the current path is a raw URL
     */
    public constructor(currentPath: string, httpCore: HttpCore, isRawUrl: boolean = true) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!httpCore) throw new Error("httpCore cannot be undefined");
        this.pathSegment = "/$value";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    };
    /**
     * Get media content for the navigation property messages from users
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @returns a RequestInfo
     */
    public createGetRequestInfo(h?: object | undefined, o?: MiddlewareOption[] | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.GET;
        h && requestInfo.setHeadersFromRawObject(h);
        o && requestInfo.addMiddlewareOptions(...o);
        return requestInfo;
    };
    /**
     * Update media content for the navigation property messages in users
     * @param body Binary request body
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @returns a RequestInfo
     */
    public createPutRequestInfo(body: ReadableStream, h?: object | undefined, o?: MiddlewareOption[] | undefined) : RequestInfo {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInfo();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.PUT;
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setStreamContent(body);
        o && requestInfo.addMiddlewareOptions(...o);
        return requestInfo;
    };
    /**
     * Get media content for the navigation property messages from users
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of ReadableStream
     */
    public get(h?: object | undefined, o?: MiddlewareOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ReadableStream | undefined> {
        const requestInfo = this.createGetRequestInfo(
            h, o
        );
        return this.httpCore?.sendPrimitiveAsync<ReadableStream>(requestInfo, "ReadableStream", responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Update media content for the navigation property messages in users
     * @param body Binary request body
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     */
    public put(body: ReadableStream, h?: object | undefined, o?: MiddlewareOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<void> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPutRequestInfo(
            body, h, o
        );
        return this.httpCore?.sendNoResponseContentAsync(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
}
