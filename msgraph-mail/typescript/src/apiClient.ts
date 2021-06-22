import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';
import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class ApiClient {
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: SerializationWriterFactory | undefined;
    public get users(): UsersRequestBuilder {
        const builder = new UsersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /**
     * Instantiates a new ApiClient and sets the default values.
     */
    public constructor() {
        this.pathSegment = "https://graph.microsoft.com/v1.0";
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users collection
     * @param id Unique identifier of the item
     * @returns a UserRequestBuilder
     */
    public usersById(id: String) : UserRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const builder = new UserRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/users/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
}
