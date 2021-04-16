import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Message} from '../../message';
import {MessagesResponse} from '../messagesResponse';

export class MessagesRequestBuilder {
    public readonly get = async (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MessagesResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<MessagesResponse>(requestInfo, responseHandler);
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
    public readonly post = async (body: Message, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Message | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<Message>(requestInfo, responseHandler);
    }
    public readonly createPostRequestInfo = (body: Message, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "/messages";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
