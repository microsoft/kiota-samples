import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {InferenceClassification} from '../inferenceClassification';
import {InferenceClassificationOverrideRequestBuilder} from './overrides/item/inferenceClassificationOverrideRequestBuilder';
import {OverridesRequestBuilder} from './overrides/overridesRequestBuilder';

export class InferenceClassificationRequestBuilder {
    public get overrides(): OverridesRequestBuilder {
        const builder = new OverridesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public readonly get = async (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassification | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<InferenceClassification>(requestInfo, responseHandler);
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
    private readonly pathSegment: string = "/inferenceClassification";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public readonly overridesById = (id: String) : InferenceClassificationOverrideRequestBuilder => {
        const builder = new InferenceClassificationOverrideRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/overrides/" + id;
        return builder;
    }
}
