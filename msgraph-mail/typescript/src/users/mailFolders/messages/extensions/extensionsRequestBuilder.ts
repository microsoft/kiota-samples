import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Extension} from '../../../extension';
import {ExtensionsResponse} from '../extensionsResponse';

export class ExtensionsRequestBuilder {
    public readonly get = async (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionsResponse | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<ExtensionsResponse>(requestInfo, responseHandler);
    }
    public readonly post = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Extension | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<Extension>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "/extensions";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
