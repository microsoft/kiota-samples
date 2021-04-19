import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {InferenceClassificationRequestBuilder} from '../inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from '../mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from '../mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from '../mailFolders/messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../mailFolders/messages/messagesRequestBuilder';

/** Builds and executes requests for operations under /users/{user-id}  */
export class UserRequestBuilder {
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        const builder = new InferenceClassificationRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    public get mailFolders(): MailFoldersRequestBuilder {
        const builder = new MailFoldersRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    public get messages(): MessagesRequestBuilder {
        const builder = new MessagesRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    }
    /** Path segment to use to build the URL for the current request builder  */
    private readonly pathSegment: string = "";
    /** Current path for the request  */
    public currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    public httpCore?: HttpCore | undefined;
    /** Factory to use to get a serializer for payload serialization  */
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
    /**
     * Gets an item from the users.mailFolders collection
     * @param id Unique identifier of the item
     * @returns a MailFolderRequestBuilder
     */
    public mailFoldersById (id: String) : MailFolderRequestBuilder {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/mailFolders/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Gets an item from the users.mailFolders.messages collection
     * @param id Unique identifier of the item
     * @returns a MessageRequestBuilder
     */
    public messagesById (id: String) : MessageRequestBuilder {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/messages/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
}
