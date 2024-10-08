/* tslint:disable */
/* eslint-disable */
// Generated by Microsoft Kiota
// @ts-ignore
import { MeRequestBuilderRequestsMetadata, type MeRequestBuilder } from './me/index.js';
// @ts-ignore
import { apiClientProxifier, registerDefaultDeserializer, registerDefaultSerializer, type BaseRequestBuilder, type KeysToExcludeForNavigationMetadata, type NavigationMetadata, type RequestAdapter } from '@microsoft/kiota-abstractions';
// @ts-ignore
import { FormParseNodeFactory, FormSerializationWriterFactory } from '@microsoft/kiota-serialization-form';
// @ts-ignore
import { JsonParseNodeFactory, JsonSerializationWriterFactory } from '@microsoft/kiota-serialization-json';
// @ts-ignore
import { MultipartSerializationWriterFactory } from '@microsoft/kiota-serialization-multipart';
// @ts-ignore
import { TextParseNodeFactory, TextSerializationWriterFactory } from '@microsoft/kiota-serialization-text';

/**
 * Instantiates a new {@link GetUserApiClient} and sets the default values.
 * @param requestAdapter The request adapter to use to execute the requests.
 */
// @ts-ignore
export function createGetUserApiClient(requestAdapter: RequestAdapter) {
    registerDefaultSerializer(JsonSerializationWriterFactory);
    registerDefaultSerializer(TextSerializationWriterFactory);
    registerDefaultSerializer(FormSerializationWriterFactory);
    registerDefaultSerializer(MultipartSerializationWriterFactory);
    registerDefaultDeserializer(JsonParseNodeFactory);
    registerDefaultDeserializer(TextParseNodeFactory);
    registerDefaultDeserializer(FormParseNodeFactory);
    if (requestAdapter.baseUrl === undefined || requestAdapter.baseUrl === null || requestAdapter.baseUrl === "") {
        requestAdapter.baseUrl = "https://graph.microsoft.com/v1.0";
    }
    const pathParameters: Record<string, unknown> = {
        "baseurl": requestAdapter.baseUrl,
    };
    return apiClientProxifier<GetUserApiClient>(requestAdapter, pathParameters, GetUserApiClientNavigationMetadata, undefined);
}
/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export interface GetUserApiClient extends BaseRequestBuilder<GetUserApiClient> {
    /**
     * The me property
     */
    get me(): MeRequestBuilder;
}
/**
 * Uri template for the request builder.
 */
export const GetUserApiClientUriTemplate = "{+baseurl}";
/**
 * Metadata for all the navigation properties in the request builder.
 */
export const GetUserApiClientNavigationMetadata: Record<Exclude<keyof GetUserApiClient, KeysToExcludeForNavigationMetadata>, NavigationMetadata> = {
    me: {
        requestsMetadata: MeRequestBuilderRequestsMetadata,
    },
};
/* tslint:enable */
/* eslint-enable */
