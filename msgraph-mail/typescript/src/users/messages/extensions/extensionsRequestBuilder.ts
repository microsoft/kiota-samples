import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Extension} from '../../extension';
import {ExtensionsResponse} from '../extensionsResponse';

export class ExtensionsRequestBuilder {
    public readonly get = (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionsResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<ExtensionsResponse>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createGetRequestInfo = (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
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
    public readonly post = (body: Extension, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Extension | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<Extension>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createPostRequestInfo = (body: Extension, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    };
    private readonly pathSegment: string = "/extensions";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
