import {PostItemRequestBuilder} from './posts/item/postItemRequestBuilder';
import {PostsRequestBuilder} from './posts/postsRequestBuilder';
import {enableBackingStoreForSerializationWriterFactory, getPathParameters, ParseNodeFactoryRegistry, registerDefaultDeserializer, registerDefaultSerializer, RequestAdapter, SerializationWriterFactoryRegistry} from '@microsoft/kiota-abstractions';
import {FormParseNodeFactory, FormSerializationWriterFactory} from '@microsoft/kiota-serialization-form';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';
import {TextParseNodeFactory, TextSerializationWriterFactory} from '@microsoft/kiota-serialization-text';

/**
 * The main entry point of the SDK, exposes the configuration and the fluent API.
 */
export class PostsClient {
    /** Path parameters for the request */
    private pathParameters: Record<string, unknown>;
    /** The posts property */
    public get posts(): PostsRequestBuilder {
        return new PostsRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The request adapter to use to execute the requests. */
    private requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private urlTemplate: string;
    /**
     * Instantiates a new PostsClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.pathParameters = {};
        this.urlTemplate = "{+baseurl}";
        this.requestAdapter = requestAdapter;
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultSerializer(TextSerializationWriterFactory);
        registerDefaultSerializer(FormSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
        registerDefaultDeserializer(TextParseNodeFactory);
        registerDefaultDeserializer(FormParseNodeFactory);
        if (requestAdapter.baseUrl === undefined || requestAdapter.baseUrl === "") {
            requestAdapter.baseUrl = "https://jsonplaceholder.typicode.com";
        }
        this.pathParameters["baseurl"] = requestAdapter.baseUrl;
    };
    /**
     * Gets an item from the ApiSdk.posts.item collection
     * @param id Unique identifier of the item
     * @returns a PostItemRequestBuilder
     */
    public postsById(id: string) : PostItemRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["post%2Did"] = id
        return new PostItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
