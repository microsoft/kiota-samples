import {MailFolderCollectionResponse} from '../../../models/';
import {createMailFolderCollectionResponseFromDiscriminatorValue} from '../../../models/createMailFolderCollectionResponseFromDiscriminatorValue';
import {createMailFolderFromDiscriminatorValue} from '../../../models/createMailFolderFromDiscriminatorValue';
import {deserializeIntoMailFolder} from '../../../models/deserializeIntoMailFolder';
import type {MailFolder} from '../../../models/mailFolder';
import {serializeMailFolder} from '../../../models/serializeMailFolder';
import {CountRequestBuilder} from './count/countRequestBuilder';
import {MailFolderItemRequestBuilder} from './item/mailFolderItemRequestBuilder';
import {MailFoldersRequestBuilderGetRequestConfiguration} from './mailFoldersRequestBuilderGetRequestConfiguration';
import {MailFoldersRequestBuilderPostRequestConfiguration} from './mailFoldersRequestBuilderPostRequestConfiguration';
import {BaseRequestBuilder, HttpMethod, RequestInformation, getPathParameters} from '@microsoft/kiota-abstractions';
import type {Parsable, ParsableFactory, RequestAdapter, RequestOption} from '@microsoft/kiota-abstractions';

/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders
 */
export class MailFoldersRequestBuilder extends BaseRequestBuilder {
    /**
     * The Count property
     */
    public get count(): CountRequestBuilder {
        return new CountRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item collection
     * @param mailFolderId The unique identifier of mailFolder
     * @returns a MailFolderItemRequestBuilder
     */
    public byMailFolderId(mailFolderId: string) : MailFolderItemRequestBuilder {
        if(!mailFolderId) throw new Error("mailFolderId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["mailFolder%2Did"] = mailFolderId
        return new MailFolderItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Instantiates a new MailFoldersRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders{?includeHiddenFolders,%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}");
    };
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MailFolderCollectionResponse
     * @see {@link https://learn.microsoft.com/graph/api/user-list-mailfolders?view=graph-rest-1.0|Find more info here}
     */
    public get(requestConfiguration?: MailFoldersRequestBuilderGetRequestConfiguration | undefined) : Promise<MailFolderCollectionResponse | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        return this.requestAdapter.sendAsync<MailFolderCollectionResponse>(requestInfo, createMailFolderCollectionResponseFromDiscriminatorValue, undefined);
    };
    /**
     * Use this API to create a new mail folder in the root folder of the user's mailbox. If you intend a new folder to be hidden, you must set the isHidden property to true on creation.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of MailFolder
     * @see {@link https://learn.microsoft.com/graph/api/user-post-mailfolders?view=graph-rest-1.0|Find more info here}
     */
    public post(body: MailFolder, requestConfiguration?: MailFoldersRequestBuilderPostRequestConfiguration | undefined) : Promise<MailFolder | undefined> {
        const requestInfo = this.toPostRequestInformation(
            body, requestConfiguration
        );
        return this.requestAdapter.sendAsync<MailFolder>(requestInfo, createMailFolderFromDiscriminatorValue, undefined);
    };
    /**
     * The user's mail folders. Read-only. Nullable.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: MailFoldersRequestBuilderGetRequestConfiguration | undefined) : RequestInformation {
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.GET;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.setQueryStringParametersFromRawObject(requestConfiguration.queryParameters);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        return requestInfo;
    };
    /**
     * Use this API to create a new mail folder in the root folder of the user's mailbox. If you intend a new folder to be hidden, you must set the isHidden property to true on creation.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPostRequestInformation(body: MailFolder, requestConfiguration?: MailFoldersRequestBuilderPostRequestConfiguration | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation();
        requestInfo.urlTemplate = this.urlTemplate;
        requestInfo.pathParameters = this.pathParameters;
        requestInfo.httpMethod = HttpMethod.POST;
        requestInfo.headers["Accept"] = ["application/json"];
        if (requestConfiguration) {
            requestInfo.addRequestHeaders(requestConfiguration.headers);
            requestInfo.addRequestOptions(requestConfiguration.options);
        }
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializeMailFolder);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a mailFoldersRequestBuilder
     */
    public withUrl(rawUrl: string) : MailFoldersRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new MailFoldersRequestBuilder(rawUrl, this.requestAdapter);
    };
}
