import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MessageRule} from '../messageRule';
import {MessageRulesResponse} from '../messageRulesResponse';

export class MessageRulesRequestBuilder {
    public readonly get = async (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MessageRulesResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<MessageRulesResponse>(requestInfo, responseHandler);
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
    public readonly post = async (body: MessageRule, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MessageRule | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<MessageRule>(requestInfo, responseHandler);
    }
    public readonly createPostRequestInfo = (body: MessageRule, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "/messageRules";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
