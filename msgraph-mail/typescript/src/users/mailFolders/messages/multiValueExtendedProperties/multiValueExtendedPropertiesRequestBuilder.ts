import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {MultiValueLegacyExtendedProperty} from '../../../multiValueLegacyExtendedProperty';
import {MultiValueExtendedPropertiesResponse} from './multiValueExtendedPropertiesResponse';

/** Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/messages/{message-id}/multiValueExtendedProperties  */
export class MultiValueExtendedPropertiesRequestBuilder {
    /**
     * Get multiValueExtendedProperties from users
     * @param q Request query parameters
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MultiValueExtendedPropertiesResponse
     */
    public get (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MultiValueExtendedPropertiesResponse | undefined> {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return this.httpCore?.sendAsync<MultiValueExtendedPropertiesResponse>(requestInfo, MultiValueExtendedPropertiesResponse, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Get multiValueExtendedProperties from users
     * @param q Request query parameters
     * @param h Request headers
     * @returns a RequestInfo
     */
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
    /**
     * Create new navigation property to multiValueExtendedProperties for users
     * @param body 
     * @param h Request headers
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of MultiValueLegacyExtendedProperty
     */
    public post (body: MultiValueLegacyExtendedProperty, h?: object | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MultiValueLegacyExtendedProperty | undefined> {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return this.httpCore?.sendAsync<MultiValueLegacyExtendedProperty>(requestInfo, MultiValueLegacyExtendedProperty, responseHandler) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Create new navigation property to multiValueExtendedProperties for users
     * @param body 
     * @param h Request headers
     * @returns a RequestInfo
     */
    public createPostRequestInfo (body: MultiValueLegacyExtendedProperty, h?: object | undefined) : RequestInfo {
        const requestInfo = new RequestInfo();
        requestInfo.URI = (this.currentPath ?? '') + this.pathSegment,
        requestInfo.httpMethod = HttpMethod.POST,
        h && requestInfo.setHeadersFromRawObject(h);
        requestInfo.setJsonContentFromParsable(body, this.serializerFactory);
        return requestInfo;
    };
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string = "/multiValueExtendedProperties";
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
}
