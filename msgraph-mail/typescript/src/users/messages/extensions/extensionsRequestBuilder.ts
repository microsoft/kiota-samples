import {Extension} from '../../extension';
import {ExtensionsResponse} from './extensionsResponse';
import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, MiddlewareOption} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/extensions  */
export class ExtensionsRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    /** The http core service to use to execute the requests.  */
    private readonly httpCore: HttpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /**
     * Instantiates a new ExtensionsRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     */
    public constructor(currentPath: string, httpCore: HttpCore) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!httpCore) throw new Error("httpCore cannot be undefined");
        this.pathSegment = "/extensions";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
    };
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
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
                    } | undefined, h?: object | undefined, o?: MiddlewareOption[] | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        q && requestInfo.setQueryStringParametersFromRawObject(q);
        o && requestInfo.addMiddlewareOptions(...o);
        return requestInfo;
    };
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @returns a RequestInfo
     */
    public createPostRequestInfo(body: Extension | undefined, h?: object | undefined, o?: MiddlewareOption[] | undefined) : RequestInfo {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.POST,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setContentFromParsable(body, this.httpCore, "application/json");
        o && requestInfo.addMiddlewareOptions(...o);
        return requestInfo;
    };
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of ExtensionsResponse
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
                    } | undefined, h?: object | undefined, o?: MiddlewareOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionsResponse | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h, o
        );
        return this.httpCore?.sendAsync<ExtensionsResponse>(requestInfo, ExtensionsResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * The collection of open extensions defined for the message. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options for HTTP middlewares
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Extension
     */
    public post(body: Extension | undefined, h?: object | undefined, o?: MiddlewareOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Extension | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPostRequestInfo(
            body, h, o
        );
        return this.httpCore?.sendAsync<Extension>(requestInfo, Extension, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
}
