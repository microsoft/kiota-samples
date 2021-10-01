import {InferenceClassificationRequestBuilder} from './inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from './mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from './mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from './messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from './messages/messagesRequestBuilder';
import {RequestAdapter} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}  */
export class UserRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /** Whether the current path is a raw URL  */
    private readonly isRawUrl: boolean;
    public get mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    public get messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder(this.currentPath + this.pathSegment, this.requestAdapter, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /** The http core service to use to execute the requests.  */
    private readonly requestAdapter: RequestAdapter;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param isRawUrl Whether the current path is a raw URL
     * @param requestAdapter The http core service to use to execute the requests.
     */
    public constructor(currentPath: string, requestAdapter: RequestAdapter, isRawUrl: boolean = true) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!requestAdapter) throw new Error("requestAdapter cannot be undefined");
        this.pathSegment = "";
        this.requestAdapter = requestAdapter;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @returns a mailFolderRequestBuilder
     */
    public mailFoldersById(id: String) : MailFolderRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new MailFolderRequestBuilder(this.currentPath + this.pathSegment + "/mailFolders/" + id, this.requestAdapter, false);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @returns a messageRequestBuilder
     */
    public messagesById(id: String) : MessageRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new MessageRequestBuilder(this.currentPath + this.pathSegment + "/messages/" + id, this.requestAdapter, false);
    };
}
