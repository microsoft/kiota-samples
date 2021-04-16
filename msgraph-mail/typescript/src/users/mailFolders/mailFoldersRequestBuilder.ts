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
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<MailFoldersResponse>(requestInfo, responseHandler);
    }
    public readonly post = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<MailFolder>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "/mailFolders";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
