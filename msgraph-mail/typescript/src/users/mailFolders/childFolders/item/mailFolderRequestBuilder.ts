import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MailFolder} from '../../mailFolder';

export class MailFolderRequestBuilder {
    public readonly get = async (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<MailFolder>(requestInfo, responseHandler);
    }
    public readonly createGetRequestInfo = (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return requestInfo;
    }
    public readonly patch = async (body: MailFolder, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createPatchRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createPatchRequestInfo = (body: MailFolder, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PATCH,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    public readonly delete = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createDeleteRequestInfo(
            h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createDeleteRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
