import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Entity} from '../../entity';

export class ContentRequestBuilder {
    public get (h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Entity | undefined> {
        const requestInfo = this.createGetRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<Entity>(requestInfo, Entity, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createGetRequestInfo (h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.GET,
        h && requestInfo.setHeadersFromRawObject(h);
        return requestInfo;
    };
    public put (h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Entity | undefined> {
        const requestInfo = this.createPutRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<Entity>(requestInfo, Entity, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createPutRequestInfo (h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.PUT,
        h && requestInfo.setHeadersFromRawObject(h);
        return requestInfo;
    };
    private readonly pathSegment: string = "/$value";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
}
