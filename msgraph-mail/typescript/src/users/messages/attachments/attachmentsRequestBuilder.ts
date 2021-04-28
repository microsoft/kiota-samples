import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {Attachment} from '../../attachment';
import {AttachmentsResponse} from './attachmentsResponse';

/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/attachments  */
export class AttachmentsRequestBuilder {
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string = "/attachments";
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: SerializationWriterFactory | undefined;
    /**
     * Get attachments from users
     * @param h Request headers
     * @param q Request query parameters
     * @returns a RequestInfo
     */
    public createGetRequestInfo (q?: {
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
     * Create new navigation property to attachments for users
     * @param body 
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createPostRequestInfo (body: Attachment, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.POST,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setJsonContentFromParsable(body, this.serializerFactory);
        return requestInfo;
    };
    /**
     * Get attachments from users
     * @param h Request headers
     * @param q Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of AttachmentsResponse
     */
    public get (q?: {
                    count?: boolean,
                    expand?: string[],
                    filter?: string,
                    orderby?: string[],
                    search?: string,
                    select?: string[],
                    skip?: number,
                    top?: number
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<AttachmentsResponse | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<AttachmentsResponse>(requestInfo, AttachmentsResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Create new navigation property to attachments for users
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Attachment
     */
    public post (body: Attachment, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Attachment | undefined> {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<Attachment>(requestInfo, Attachment, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
}
