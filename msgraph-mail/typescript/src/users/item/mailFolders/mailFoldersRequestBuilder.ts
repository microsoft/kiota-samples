import {MailFolder} from '../../../models/microsoft/graph/mailFolder';
import {MailFoldersResponse} from './mailFoldersResponse';
import {HttpMethod, RequestInformation, RequestOption, RequestAdapter, ResponseHandler, Parsable} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders  */
export class MailFoldersRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    /** Whether the current path is a raw URL  */
    private readonly isRawUrl: boolean;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /** The http core service to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    /**
     * Instantiates a new MailFoldersRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param isRawUrl Whether the current path is a raw URL
     * @param requestAdapter The http core service to use to execute the requests.
     */
    public constructor(currentPath: string, requestAdapter: RequestAdapter, isRawUrl: boolean = true) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.pathSegment = "/mailFolders";
        this.requestAdapter = requestAdapter;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    };
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @returns a RequestInformation
     */
    public createGetRequestInformation(q?: {
                    count?: boolean,
                    expand?: string[],
                    filter?: string,
                    orderby?: string[],
                    search?: string,
                    select?: string[],
                    skip?: number,
                    top?: number
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
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @returns a RequestInformation
     */
    public createPostRequestInformation(body: MailFolder | undefined, h?: object | undefined, o?: RequestOption[] | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.setUri(this.currentPath, this.pathSegment, this.isRawUrl);
        requestInfo.httpMethod = HttpMethod.POST;
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body);
        o && requestInfo.addRequestOptions(...o);
        return requestInfo;
    };
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param h Request headers
     * @param o Request options
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MailFoldersResponse
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
                    } | undefined, h?: object | undefined, o?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFoldersResponse | undefined> {
        const requestInfo = this.createGetRequestInformation(
            q, h, o
        );
        return this.requestAdapter?.sendAsync<MailFoldersResponse>(requestInfo, MailFoldersResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param body 
     * @param h Request headers
     * @param o Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MailFolder
     */
    public post(body: MailFolder | undefined, h?: object | undefined, o?: RequestOption[] | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPostRequestInformation(
            body, h, o
        );
        return this.requestAdapter?.sendAsync<MailFolder>(requestInfo, MailFolder, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
}
