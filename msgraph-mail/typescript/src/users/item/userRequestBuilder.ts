import {SerializationWriter, HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {InferenceClassificationRequestBuilder} from '../inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from '../mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from '../mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from '../mailFolders/messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../mailFolders/messages/messagesRequestBuilder';

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
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public serializerFactory?: ((mediaType: string) => SerializationWriter) | undefined;
    public mailFoldersById (id: String) : MailFolderRequestBuilder {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/mailFolders/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    public messagesById (id: String) : MessageRequestBuilder {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/messages/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
}
