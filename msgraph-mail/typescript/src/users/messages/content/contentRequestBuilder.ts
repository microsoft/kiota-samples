import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';

export class ContentRequestBuilder {
    public readonly get = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createGetRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
        } as RequestInfo;
        return requestInfo;
    }
    public readonly put = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createPutRequestInfo(
            h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createPutRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PUT,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "/$value";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
