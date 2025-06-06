/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue, createInferenceClassificationOverrideFromDiscriminatorValue, serializeInferenceClassificationOverride, type InferenceClassificationOverride, type InferenceClassificationOverrideCollectionResponse } from '../../../../models/index.js';
// @ts-ignore
import { createODataErrorFromDiscriminatorValue, type ODataError } from '../../../../models/oDataErrors/index.js';
// @ts-ignore
import { CountRequestBuilderRequestsMetadata, type CountRequestBuilder } from './count/index.js';
// @ts-ignore
import { InferenceClassificationOverrideItemRequestBuilderRequestsMetadata, type InferenceClassificationOverrideItemRequestBuilder } from './item/index.js';
// @ts-ignore
import { type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users/{user-id}/inferenceClassification/overrides
 */
export interface OverridesRequestBuilder extends BaseRequestBuilder<OverridesRequestBuilder> {
    /**
     * The Count property
     */
    get count(): CountRequestBuilder;
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.inferenceClassification.overrides.item collection
     * @param inferenceClassificationOverrideId The unique identifier of inferenceClassificationOverride
     * @returns {InferenceClassificationOverrideItemRequestBuilder}
     */
     byInferenceClassificationOverrideId(inferenceClassificationOverrideId: string) : InferenceClassificationOverrideItemRequestBuilder;
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<InferenceClassificationOverrideCollectionResponse>}
     * @throws {ODataError} error when the service returns a 4XX or 5XX status code
     */
     get(requestConfiguration?: RequestConfiguration<OverridesRequestBuilderGetQueryParameters> | undefined) : Promise<InferenceClassificationOverrideCollectionResponse | undefined>;
    /**
     * Create new navigation property to overrides for users
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {Promise<InferenceClassificationOverride>}
     * @throws {ODataError} error when the service returns a 4XX or 5XX status code
     */
     post(body: InferenceClassificationOverride, requestConfiguration?: RequestConfiguration<object> | undefined) : Promise<InferenceClassificationOverride | undefined>;
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<OverridesRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
    /**
     * Create new navigation property to overrides for users
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns {RequestInformation}
     */
     toPostRequestInformation(body: InferenceClassificationOverride, requestConfiguration?: RequestConfiguration<object> | undefined) : RequestInformation;
}
/**
 * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
 */
export interface OverridesRequestBuilderGetQueryParameters {
    /**
     * Include count of items
     */
    count?: boolean;
    /**
     * Expand related entities
     */
    expand?: string[];
    /**
     * Filter items by property values
     */
    filter?: string;
    /**
     * Order items by property values
     */
    orderby?: string[];
    /**
     * Search items by search phrases
     */
    search?: string;
    /**
     * Select properties to be returned
     */
    select?: string[];
    /**
     * Skip the first n items
     */
    skip?: number;
    /**
     * Show only the first n items
     */
    top?: number;
}
/**
 * Uri template for the request builder.
 */
export const OverridesRequestBuilderUriTemplate = "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}";
/**
 * Mapper for query parameters from symbol name to serialization name represented as a constant.
 */
const OverridesRequestBuilderGetQueryParametersMapper: Record<string, string> = {
    "count": "%24count",
    "expand": "%24expand",
    "filter": "%24filter",
    "orderby": "%24orderby",
    "search": "%24search",
    "select": "%24select",
    "skip": "%24skip",
    "top": "%24top",
};
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const OverridesRequestBuilderNavigationMetadata: Record<Exclude<keyof OverridesRequestBuilder, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    byInferenceClassificationOverrideId: {
        requestsMetadata: InferenceClassificationOverrideItemRequestBuilderRequestsMetadata,
        pathParametersMappings: ["inferenceClassificationOverride%2Did"],
    },
    count: {
        requestsMetadata: CountRequestBuilderRequestsMetadata,
    },
};
/**
 * Metadata for all the requests in the request builder.
 */
export const OverridesRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        uriTemplate: OverridesRequestBuilderUriTemplate,
        responseBodyContentType: "application/json",
        errorMappings: {
            XXX: createODataErrorFromDiscriminatorValue as ParsableFactory<Parsable>,
        },
        adapterMethodName: "send",
        responseBodyFactory:  createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue,
        queryParametersMapper: OverridesRequestBuilderGetQueryParametersMapper,
    },
    post: {
        uriTemplate: OverridesRequestBuilderUriTemplate,
        responseBodyContentType: "application/json",
        errorMappings: {
            XXX: createODataErrorFromDiscriminatorValue as ParsableFactory<Parsable>,
        },
        adapterMethodName: "send",
        responseBodyFactory:  createInferenceClassificationOverrideFromDiscriminatorValue,
        requestBodyContentType: "application/json",
        requestBodySerializer: serializeInferenceClassificationOverride,
        requestInformationContentSetMethod: "setContentFromParsable",
    },
};
/* tslint:enable */
/* eslint-enable */
