import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Attachment} from '../../../attachment';

export class AttachmentRequestBuilder {
    public readonly get = (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Attachment | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<Attachment>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createGetRequestInfo = (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return requestInfo;
    };
    public readonly patch = (body: Attachment, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createPatchRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<object>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createPatchRequestInfo = (body: Attachment, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.PATCH,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    };
    public readonly delete = (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createDeleteRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<object>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createDeleteRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return requestInfo;
    };
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
