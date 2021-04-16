import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';

export class ContentRequestBuilder {
    public readonly get = (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<object>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createGetRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.GET,
        } as RequestInfo;
        return requestInfo;
    };
    public readonly put = (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createPutRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<object>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createPutRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.PUT,
        } as RequestInfo;
        return requestInfo;
    };
    private readonly pathSegment: string = "/$value";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
