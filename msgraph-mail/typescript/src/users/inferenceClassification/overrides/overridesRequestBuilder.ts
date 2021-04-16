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
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<OverridesResponse>(requestInfo, responseHandler);
    }
    public readonly post = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassificationOverride | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.POST,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<InferenceClassificationOverride>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "/overrides";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
