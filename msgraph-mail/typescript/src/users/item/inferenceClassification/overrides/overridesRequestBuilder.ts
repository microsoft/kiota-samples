import {InferenceClassificationOverride, InferenceClassificationOverrideCollectionResponse} from '../../../../models/';
import {createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue} from '../../../../models/createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue';
import {createInferenceClassificationOverrideFromDiscriminatorValue} from '../../../../models/createInferenceClassificationOverrideFromDiscriminatorValue';
import {InferenceClassificationOverrideItemRequestBuilder} from './item/inferenceClassificationOverrideItemRequestBuilder';
import {OverridesRequestBuilderGetRequestConfiguration} from './overridesRequestBuilderGetRequestConfiguration';
import {OverridesRequestBuilderPostRequestConfiguration} from './overridesRequestBuilderPostRequestConfiguration';
import {BaseRequestBuilder, getPathParameters, HttpMethod, Parsable, ParsableFactory, RequestAdapter, RequestInformation, RequestOption, ResponseHandler} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users/{user-id}/inferenceClassification/overrides
 */
export class OverridesRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new OverridesRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}");
    };
    /**
     * Get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of InferenceClassificationOverrideCollectionResponse
     * @see {@link https://docs.microsoft.com/graph/api/inferenceclassification-list-overrides?view=graph-rest-1.0|Find more info here}
     */
    public get(requestConfiguration?: OverridesRequestBuilderGetRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassificationOverrideCollectionResponse | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter?.sendAsync<InferenceClassificationOverrideCollectionResponse>(requestInfo, createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * Create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. **Note**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of InferenceClassificationOverride
     * @see {@link https://docs.microsoft.com/graph/api/inferenceclassification-post-overrides?view=graph-rest-1.0|Find more info here}
     */
    public post(body: InferenceClassificationOverride | undefined, requestConfiguration?: OverridesRequestBuilderPostRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<InferenceClassificationOverride | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.toPostRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter?.sendAsync<InferenceClassificationOverride>(requestInfo, createInferenceClassificationOverrideFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * Get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: OverridesRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * Create an override for a sender identified by an SMTP address. Future messages from that SMTP address will be consistently classifiedas specified in the override. **Note**
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPostRequestInformation(body: InferenceClassificationOverride | undefined, requestConfiguration?: OverridesRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body);
        return requestInfo;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.inferenceClassification.overrides.item collection
     * @param inferenceClassificationOverrideId Unique identifier of the item
     * @returns a InferenceClassificationOverrideItemRequestBuilder
     */
    public withInferenceClassificationOverrideId(inferenceClassificationOverrideId: string) : InferenceClassificationOverrideItemRequestBuilder {
        if(!inferenceClassificationOverrideId) throw new Error("inferenceClassificationOverrideId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["inferenceClassificationOverride%2Did"] = inferenceClassificationOverrideId
        return new InferenceClassificationOverrideItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
