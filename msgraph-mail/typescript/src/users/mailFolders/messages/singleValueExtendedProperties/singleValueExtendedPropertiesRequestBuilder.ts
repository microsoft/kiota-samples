import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {SingleValueLegacyExtendedProperty} from '../../../singleValueLegacyExtendedProperty';
import {SingleValueExtendedPropertiesResponse} from '../singleValueExtendedPropertiesResponse';

export class SingleValueExtendedPropertiesRequestBuilder {
    public readonly get = (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<SingleValueExtendedPropertiesResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<SingleValueExtendedPropertiesResponse>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
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
    public readonly post = (body: SingleValueLegacyExtendedProperty, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<SingleValueLegacyExtendedProperty | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<SingleValueLegacyExtendedProperty>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createPostRequestInfo = (body: SingleValueLegacyExtendedProperty, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    };
    private readonly pathSegment: string = "/singleValueExtendedProperties";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
