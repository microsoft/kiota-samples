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
        const requestInfo = this.createGetRequestInfo(
            q, h
        );
        return await this.httpCore?.sendAsync<InferenceClassification>(requestInfo, responseHandler);
    }
    public readonly createGetRequestInfo = (q?: {
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
    public readonly patch = async (body: InferenceClassification, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createPatchRequestInfo(
            body, h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createPatchRequestInfo = (body: InferenceClassification, h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PATCH,
            content: body as unknown,
        } as RequestInfo;
        return requestInfo;
    }
    public readonly delete = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = this.createDeleteRequestInfo(
            h
        );
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly createDeleteRequestInfo = (h?: {} | undefined) : RequestInfo => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return requestInfo;
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
