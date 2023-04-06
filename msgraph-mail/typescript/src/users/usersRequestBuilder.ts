import {UserItemRequestBuilder} from './item/userItemRequestBuilder';
import {BaseRequestBuilder, getPathParameters, RequestAdapter} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users
 */
export class UsersRequestBuilder extends BaseRequestBuilder {
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users");
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item collection
     * @param userId Unique identifier of the item
     * @returns a UserItemRequestBuilder
     */
    public withUserId(userId: string) : UserItemRequestBuilder {
        if(!userId) throw new Error("userId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["user%2Did"] = userId
        return new UserItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
