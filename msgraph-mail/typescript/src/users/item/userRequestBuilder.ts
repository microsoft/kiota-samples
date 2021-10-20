import {InferenceClassificationRequestBuilder} from './inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from './mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from './mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from './messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from './messages/messagesRequestBuilder';
import {getUrlTemplateParameters, RequestAdapter} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}  */
export class UserRequestBuilder {
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder(this.urlTemplateParameters, this.requestAdapter);
    }
    public get mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder(this.urlTemplateParameters, this.requestAdapter);
    }
    public get messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder(this.urlTemplateParameters, this.requestAdapter);
    }
    /** The request adapter to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private readonly urlTemplate: string;
    /** Url template parameters for the request  */
    private readonly urlTemplateParameters: Map<string, string>;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @param urlTemplateParameters The raw url or the Url template parameters for the request.
     */
    public constructor(urlTemplateParameters: Map<string, string> | string | undefined, requestAdapter: RequestAdapter) {
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        if(!urlTemplateParameters) throw new Error("urlTemplateParameters cannot be undefined");
        this.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}";
        const urlTplParams = getUrlTemplateParameters(urlTemplateParameters);
        this.urlTemplateParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @returns a mailFolderRequestBuilder
     */
    public mailFoldersById(id: string) : MailFolderRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getUrlTemplateParameters(this.urlTemplateParameters);
        id && urlTplParams.set("mailFolder_id", id);
        return new MailFolderRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @returns a messageRequestBuilder
     */
    public messagesById(id: string) : MessageRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        const urlTplParams = getUrlTemplateParameters(this.urlTemplateParameters);
        id && urlTplParams.set("message_id", id);
        return new MessageRequestBuilder(urlTplParams, this.requestAdapter);
    };
}
