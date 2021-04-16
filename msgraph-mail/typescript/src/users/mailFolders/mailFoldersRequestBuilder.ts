import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MailFoldersResponse} from '../mailFoldersResponse';
import {MailFolder} from './mailFolder';

export class MailFoldersRequestBuilder {
    public readonly get = async (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFoldersResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<MailFoldersResponse>(requestInfo, responseHandler);
    }
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
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return requestInfo;
    }
    public readonly post = async (body: MailFolder, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<MailFolder>(requestInfo, responseHandler);
    }
    public readonly createPostRequestInfo = (body: MailFolder, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "/mailFolders";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
