import {InferenceClassificationRequestBuilder} from '../inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from '../mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from '../mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from '../messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../messages/messagesRequestBuilder';
import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, MiddlewareOption} from '@microsoft/kiota-abstractions';

/** Builds and executes requests for operations under /users/{user-id}  */
export class UserRequestBuilder {
    /** Current path for the request  */
    private readonly currentPath: string;
    /** The http core service to use to execute the requests.  */
    private readonly httpCore: HttpCore;
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder(this.currentPath + this.pathSegment, this.httpCore, false);
    }
    /** Whether the current path is a raw URL  */
    private readonly isRawUrl: boolean;
    public get mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder(this.currentPath + this.pathSegment, this.httpCore, false);
    }
    public get messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder(this.currentPath + this.pathSegment, this.httpCore, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param httpCore The http core service to use to execute the requests.
     * @param isRawUrl Whether the current path is a raw URL
     */
    public constructor(currentPath: string, httpCore: HttpCore, isRawUrl: boolean = true) {
        if(!currentPath) throw new Error("currentPath cannot be undefined");
        if(!httpCore) throw new Error("httpCore cannot be undefined");
        this.pathSegment = "";
        this.httpCore = httpCore;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders collection
     * @param id Unique identifier of the item
     * @returns a MailFolderRequestBuilder
     */
    public mailFoldersById(id: String) : MailFolderRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new MailFolderRequestBuilder(this.currentPath + this.pathSegment + "/mailFolders/" + id, this.httpCore, false);
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.messages collection
     * @param id Unique identifier of the item
     * @returns a MessageRequestBuilder
     */
    public messagesById(id: String) : MessageRequestBuilder {
        if(!id) throw new Error("id cannot be undefined");
        return new MessageRequestBuilder(this.currentPath + this.pathSegment + "/messages/" + id, this.httpCore, false);
    };
}
