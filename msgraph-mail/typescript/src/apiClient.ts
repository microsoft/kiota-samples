import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';
import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, MiddlewareOption, registerDefaultSerializer, registerDefaultDeserializer, enableBackingStoreForSerializationWriterFactory, SerializationWriterFactoryRegistry, ParseNodeFactoryRegistry} from '@microsoft/kiota-abstractions';
import {JsonParseNodeFactory, JsonSerializationWriterFactory} from '@microsoft/kiota-serialization-json';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class ApiClient {
    /** The http core service to use to execute the requests.  */
    private readonly httpCore: HttpCore;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    public get users(): UsersRequestBuilder {
        return new UsersRequestBuilder(this.pathSegment, this.httpCore);
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param httpCore The http core service to use to execute the requests.
     */
    public constructor(httpCore: HttpCore) {
        if(!httpCore) throw new Error("httpCore cannot be undefined");
        this.pathSegment = "https://graph.microsoft.com/v1.0";
        this.httpCore = httpCore;
        registerDefaultSerializer(JsonSerializationWriterFactory);
        registerDefaultDeserializer(JsonParseNodeFactory);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users collection
     * @param id Unique identifier of the item
     * @returns a UserRequestBuilder
     */
    public usersById(id: String) : UserRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new UserRequestBuilder(this.pathSegment + "/users/" + id, this.httpCore);
    };
}
