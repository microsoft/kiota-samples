import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';

/** The main entry point of the SDK, exposes the configuration and the fluent API.  */
export class GraphClient {
    public get users(): UsersRequestBuilder {
        const builder = new UsersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string = "https://graph.microsoft.com/v1.0";
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
    /**
     * Gets an item from the users collection
     * @param id Unique identifier of the item
     * @returns a UserRequestBuilder
     */
    public usersById (id: String) : UserRequestBuilder {
        const builder = new UserRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/users/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
}
