import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users  */
export class UsersRequestBuilder {
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     */
    public constructor() {
        this.pathSegment = "/users";
    };
}
