import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';
import {enableBackingStoreForSerializationWriterFactory, getUrlTemplateParameters, ParseNodeFactoryRegistry, registerDefaultDeserializer, registerDefaultSerializer, RequestAdapter, SerializationWriterFactoryRegistry} from '@microsoft/kiota-abstractions';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class ApiClient {
    /** The request adapter to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private readonly urlTemplate: string;
    /** Url template parameters for the request  */
    private readonly urlTemplateParameters: Map<string, string>;
    public get users(): UsersRequestBuilder {
        return new UsersRequestBuilder(this.urlTemplateParameters, this.requestAdapter);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.urlTemplateParameters = new Map<string, string>();
        this.urlTemplate = "https://graph.microsoft.com/v1.0";
        this.requestAdapter = requestAdapter;
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item collection
     * @param id Unique identifier of the item
     * @returns a userRequestBuilder
     */
    public usersById(id: string) : UserRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getUrlTemplateParameters(this.urlTemplateParameters);
        id && urlTplParams.set("user_id", id);
        return new UserRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
