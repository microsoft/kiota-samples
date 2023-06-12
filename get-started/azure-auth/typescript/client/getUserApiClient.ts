import {MeRequestBuilder} from './me/meRequestBuilder';
import {BaseRequestBuilder, enableBackingStoreForSerializationWriterFactory, ParseNodeFactoryRegistry, registerDefaultDeserializer, registerDefaultSerializer, RequestAdapter, SerializationWriterFactoryRegistry} from '@microsoft/kiota-abstractions';
import {FormParseNodeFactory, FormSerializationWriterFactory} from '@microsoft/kiota-serialization-form';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';
import {TextParseNodeFactory, TextSerializationWriterFactory} from '@microsoft/kiota-serialization-text';

/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export class GetUserApiClient extends BaseRequestBuilder {
    /** The me property */
    public get me(): MeRequestBuilder {
        return new MeRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new GetUserApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        super({}, requestAdapter, "{+baseurl}");
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultSerializer(TextSerializationWriterFactory);
        registerDefaultSerializer(FormSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
        registerDefaultDeserializer(TextParseNodeFactory);
        registerDefaultDeserializer(FormParseNodeFactory);
        if (requestAdapter.baseUrl === undefined || requestAdapter.baseUrl === "") {
            requestAdapter.baseUrl = "https://graph.microsoft.com/v1.0";
        }
        this.pathParameters["baseurl"] = requestAdapter.baseUrl;
    };
}
