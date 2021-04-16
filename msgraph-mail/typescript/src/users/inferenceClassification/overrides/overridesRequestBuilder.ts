import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {InferenceClassificationOverride} from '../inferenceClassificationOverride';
import {OverridesResponse} from '../overridesResponse';

export class OverridesRequestBuilder {
    public readonly get = async (q?: {
                    top?: number,
                    skip?: number,
                    search?: string,
                    filter?: string,
                    count?: boolean,
                    orderby?: string[],
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<OverridesResponse | undefined> => {
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<OverridesResponse>(requestInfo, responseHandler);
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
    public readonly post = async (body: InferenceClassificationOverride, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassificationOverride | undefined> => {
        const requestInfo = this.createPostRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<InferenceClassificationOverride>(requestInfo, responseHandler);
    }
    public readonly createPostRequestInfo = (body: InferenceClassificationOverride, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    private readonly pathSegment: string = "/overrides";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
