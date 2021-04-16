import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Extension} from '../../../extension';
import {ExtensionsResponse} from '../extensionsResponse';

export class ExtensionsRequestBuilder {
    public get (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionsResponse | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<ExtensionsResponse>(requestInfo, ExtensionsResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createGetRequestInfo (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        q && requestInfo.setQueryStringParametersFromRawObject(q);
        return requestInfo;
    };
    public post (body: Extension, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Extension | undefined> {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<Extension>(requestInfo, Extension, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createPostRequestInfo (body: Extension, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.POST,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setJsonContentFromParsable(body, this.serializerFactory);
        return requestInfo;
    };
    private readonly pathSegment: string = "/extensions";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
}
