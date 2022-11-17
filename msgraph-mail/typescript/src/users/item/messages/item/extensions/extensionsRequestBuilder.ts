<<<<<<< HEAD
import {ExtensionCollectionResponseImpl, ExtensionImpl} from '../../../../../models/';
import {createExtensionCollectionResponseFromDiscriminatorValue} from '../../../../../models/createExtensionCollectionResponseFromDiscriminatorValue';
import {createExtensionFromDiscriminatorValue} from '../../../../../models/createExtensionFromDiscriminatorValue';
import {Extension} from '../../../../../models/extension';
=======
import {Extension, ExtensionCollectionResponse} from '../../../../../models/';
import {createExtensionCollectionResponseFromDiscriminatorValue} from '../../../../../models/createExtensionCollectionResponseFromDiscriminatorValue';
import {createExtensionFromDiscriminatorValue} from '../../../../../models/createExtensionFromDiscriminatorValue';
>>>>>>> main
import {ExtensionsRequestBuilderGetRequestConfiguration} from './extensionsRequestBuilderGetRequestConfiguration';
import {ExtensionsRequestBuilderPostRequestConfiguration} from './extensionsRequestBuilderPostRequestConfiguration';
import {getPathParameters, HttpMethod, Parsable, ParsableFactory, RequestAdapter, RequestInformation, RequestOption, ResponseHandler} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/extensions */
export class ExtensionsRequestBuilder {
    /** Path parameters for the request */
    private readonly pathParameters: Record<string, unknown>;
    /** The request adapter to use to execute the requests. */
    private readonly requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private readonly urlTemplate: string;
    /**
     * Instantiates a new ExtensionsRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        if(!pathParameters) throw new Error("pathParameters cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/extensions{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
        const urlTplParams = getPathParameters(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    };
    /**
     * Get extensions from users
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public createGetRequestInformation(requestConfiguration?: ExtensionsRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = "application/json";
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
<<<<<<< HEAD
     * Create new navigation property to extensions for users
=======
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
>>>>>>> main
     * @param body 
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public createPostRequestInformation(body: Extension | undefined, requestConfiguration?: ExtensionsRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.headers["Accept"] = "application/json";
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        const parsableBody = new ExtensionImpl(body)
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", parsableBody);
        return requestInfo;
    };
    /**
     * Get extensions from users
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of ExtensionCollectionResponse
     */
<<<<<<< HEAD
    public get(requestConfiguration?: ExtensionsRequestBuilderGetRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionCollectionResponseImpl | undefined> {
        const requestInfo = this.createGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter?.sendAsync<ExtensionCollectionResponseImpl>(requestInfo, createExtensionCollectionResponseFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('http core is null'));
    };
    /**
     * Create new navigation property to extensions for users
=======
    public get(requestConfiguration?: ExtensionsRequestBuilderGetRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionCollectionResponse | undefined> {
        const requestInfo = this.createGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter?.sendAsync<ExtensionCollectionResponse>(requestInfo, createExtensionCollectionResponseFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
    };
    /**
     * Create an open extension (openTypeExtension object) and add custom properties in a new or existing instance of a resource. You can create an open extension in a resource instance and store custom data to it all in the same operation, except for specific resources. See known limitations of open extensions for more information. The table in the Permissions section lists the resources that support open extensions.
>>>>>>> main
     * @param body 
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @returns a Promise of Extension
     */
    public post(body: Extension | undefined, requestConfiguration?: ExtensionsRequestBuilderPostRequestConfiguration | undefined, responseHandler?: ResponseHandler | undefined) : Promise<ExtensionImpl | undefined> {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = this.createPostRequestInformation(
            body, requestConfiguration
        );
<<<<<<< HEAD
        return this.requestAdapter?.sendAsync<ExtensionImpl>(requestInfo, createExtensionFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('http core is null'));
=======
        return this.requestAdapter?.sendAsync<Extension>(requestInfo, createExtensionFromDiscriminatorValue, responseHandler, undefined) ?? Promise.reject(new Error('request adapter is null'));
>>>>>>> main
    };
}
