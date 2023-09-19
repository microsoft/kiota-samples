import { UsersRequestBuilder } from './users/usersRequestBuilder';
import { BaseRequestBuilder, enableBackingStoreForSerializationWriterFactory, ParseNodeFactoryRegistry, registerDefaultDeserializer, registerDefaultSerializer, SerializationWriterFactoryRegistry, type RequestAdapter } from '@microsoft/kiota-abstractions';
import { FormParseNodeFactory, FormSerializationWriterFactory } from '@microsoft/kiota-serialization-form';
import { JsonParseNodeFactory, JsonSerializationWriterFactory } from '@microsoft/kiota-serialization-json';
import { MultipartSerializationWriterFactory } from '@microsoft/kiota-serialization-multipart';
import { TextParseNodeFactory, TextSerializationWriterFactory } from '@microsoft/kiota-serialization-text';

/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export class ApiClient extends BaseRequestBuilder {
    /**
     * The users property
     */
    public get users(): UsersRequestBuilder {
        return new UsersRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        super({}, requestAdapter, "{+baseurl}");
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultSerializer(TextSerializationWriterFactory);
        registerDefaultSerializer(FormSerializationWriterFactory);
        registerDefaultSerializer(MultipartSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
        registerDefaultDeserializer(TextParseNodeFactory);
        registerDefaultDeserializer(FormParseNodeFactory);
        if (requestAdapter.baseUrl === undefined || requestAdapter.baseUrl === "") {
            requestAdapter.baseUrl = "https://graph.microsoft.com/v1.0";
        }
        this.pathParameters["baseurl"] = requestAdapter.baseUrl;
    };
}
