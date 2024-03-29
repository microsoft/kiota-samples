/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
import { createPetFromDiscriminatorValue, type Pet } from '../../models/';
import { type BaseRequestBuilder, type Parsable, type ParsableFactory, type RequestConfiguration, type RequestInformation, type RequestsMetadata } from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /pet/findByTags
 */
export interface FindByTagsRequestBuilder extends BaseRequestBuilder<FindByTagsRequestBuilder> {
    /**
     * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Pet
     * @deprecated 
     */
     get(requestConfiguration?: RequestConfiguration<FindByTagsRequestBuilderGetQueryParameters> | undefined) : Promise<Pet[] | undefined>;
    /**
     * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     * @deprecated 
     */
     toGetRequestInformation(requestConfiguration?: RequestConfiguration<FindByTagsRequestBuilderGetQueryParameters> | undefined) : RequestInformation;
}
/**
 * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
 */
export interface FindByTagsRequestBuilderGetQueryParameters {
    /**
     * Tags to filter by
     */
    tags?: string[];
}
/**
 * Metadata for all the requests in the request builder.
 */
export const FindByTagsRequestBuilderRequestsMetadata: RequestsMetadata = {
    get: {
        responseBodyContentType: "application/json",
        adapterMethodName: "sendCollectionAsync",
        responseBodyFactory:  createPetFromDiscriminatorValue,
    },
};
/**
 * Uri template for the request builder.
 */
export const FindByTagsRequestBuilderUriTemplate = "{+baseurl}/pet/findByTags?tags={tags}";
/* tslint:enable */
/* eslint-enable */
