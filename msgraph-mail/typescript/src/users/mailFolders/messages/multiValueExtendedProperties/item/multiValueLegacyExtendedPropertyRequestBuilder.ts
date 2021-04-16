import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MultiValueLegacyExtendedProperty} from '../../../../multiValueLegacyExtendedProperty';

export class MultiValueLegacyExtendedPropertyRequestBuilder {
    public readonly get = async (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MultiValueLegacyExtendedProperty | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<MultiValueLegacyExtendedProperty>(requestInfo, responseHandler);
    }
    public readonly patch = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PATCH,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly delete = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
