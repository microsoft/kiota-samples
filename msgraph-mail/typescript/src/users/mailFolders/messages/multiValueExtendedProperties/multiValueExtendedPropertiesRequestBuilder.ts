import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MultiValueLegacyExtendedProperty} from '../../../multiValueLegacyExtendedProperty';
import {MultiValueExtendedPropertiesResponse} from '../multiValueExtendedPropertiesResponse';

export class MultiValueExtendedPropertiesRequestBuilder {
    public readonly get = (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MultiValueExtendedPropertiesResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<MultiValueExtendedPropertiesResponse>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
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
    public readonly post = (body: MultiValueLegacyExtendedProperty, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MultiValueLegacyExtendedProperty | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<MultiValueLegacyExtendedProperty>(requestInfo, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public readonly createPostRequestInfo = (body: MultiValueLegacyExtendedProperty, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: (this.currentPath ?? '') + this.pathSegment,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    };
    private readonly pathSegment: string = "/multiValueExtendedProperties";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
