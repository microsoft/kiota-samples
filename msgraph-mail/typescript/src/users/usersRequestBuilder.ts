import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';

export class UsersRequestBuilder {
    private readonly pathSegment: string = "/users";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
}
