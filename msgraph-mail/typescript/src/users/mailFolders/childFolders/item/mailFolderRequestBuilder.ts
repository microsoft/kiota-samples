import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {Entity} from '../../../entity';
import {MailFolder} from '../../mailFolder';

export class MailFolderRequestBuilder {
    public get (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<MailFolder>(requestInfo, MailFolder, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createGetRequestInfo (q?: {
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
    public patch (body: MailFolder, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Entity | undefined> {
        const requestInfo = this.createPatchRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<Entity>(requestInfo, Entity, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createPatchRequestInfo (body: MailFolder, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.PATCH,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setJsonContentFromParsable(body, this.serializerFactory);
        return requestInfo;
    };
    public delete (h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<Entity | undefined> {
        const requestInfo = this.createDeleteRequestInfo(
            h
        );
        return this.httpCore?.sendAsync<Entity>(requestInfo, Entity, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    public createDeleteRequestInfo (h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.DELETE,
        h && requestInfo.setHeadersFromRawObject(h);
        return requestInfo;
    };
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
}
