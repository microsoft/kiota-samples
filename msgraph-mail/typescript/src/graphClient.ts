import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {UserRequestBuilder} from './users/item/userRequestBuilder';
import {UsersRequestBuilder} from './users/usersRequestBuilder';

export class GraphClient {
    public get users(): UsersRequestBuilder {
        const builder = new UsersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        return builder;
    }
    private readonly pathSegment: string = "https://graph.microsoft.com/v1.0";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public readonly usersById = (id: String) : UserRequestBuilder => {
        const builder = new UserRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/users/" + id;
        builder.httpCore = this.httpCore;
        return builder;
    };
}
