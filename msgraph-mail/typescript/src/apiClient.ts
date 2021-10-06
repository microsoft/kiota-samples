import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';
import {RequestAdapter, registerDefaultSerializer, enableBackingStoreForSerializationWriterFactory, SerializationWriterFactoryRegistry, registerDefaultDeserializer, ParseNodeFactoryRegistry} from '@microsoft/kiota-abstractions';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class ApiClient {
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /** The http core service to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    public get users(): UsersRequestBuilder {
        return new UsersRequestBuilder(this.pathSegment, this.requestAdapter, false);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param requestAdapter The http core service to use to execute the requests.
     */
    public constructor(requestAdapter: RequestAdapter) {
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.pathSegment = "https://graph.microsoft.com/v1.0";
        this.requestAdapter = requestAdapter;
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item collection
     * @param id Unique identifier of the item
     * @returns a userRequestBuilder
     */
    public usersById(id: String) : UserRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new UserRequestBuilder(this.pathSegment + "/users/" + id, this.requestAdapter, false);
    };
}
