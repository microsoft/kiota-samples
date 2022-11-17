import {InferenceClassificationRequestBuilder} from './inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderItemRequestBuilder} from './mailFolders/item/mailFolderItemRequestBuilder';
import {MailFoldersRequestBuilder} from './mailFolders/mailFoldersRequestBuilder';
import {MessageItemRequestBuilder} from './messages/item/messageItemRequestBuilder';
import {MessagesRequestBuilder} from './messages/messagesRequestBuilder';
import {getPathParameters, RequestAdapter} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id} */
export class UserItemRequestBuilder {
    /** The inferenceClassification property */
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The mailFolders property */
    public get mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** The messages property */
    public get messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /** Path parameters for the request */
    private readonly pathParameters: Record<string, unknown>;
    /** The request adapter to use to execute the requests. */
    private readonly requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder */
    private readonly urlTemplate: string;
    /**
     * Instantiates a new UserItemRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        if(!pathParameters) throw new Error("pathParameters cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.urlTemplate = "{+baseurl}/users/{user%2Did}";
        const urlTplParams = getPathParameters(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    };
    /**
     * Gets an item from the ApiSdk.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @returns a MailFolderItemRequestBuilder
     */
    public mailFoldersById(id: string) : MailFolderItemRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["mailFolder%2Did"] = id
        return new MailFolderItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Gets an item from the ApiSdk.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @returns a MessageItemRequestBuilder
     */
    public messagesById(id: string) : MessageItemRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["message%2Did"] = id
        return new MessageItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
