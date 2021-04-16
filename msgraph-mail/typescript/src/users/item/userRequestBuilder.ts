import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {InferenceClassificationRequestBuilder} from '../inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from '../mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from '../mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from '../mailFolders/messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../mailFolders/messages/messagesRequestBuilder';

export class UserRequestBuilder {
    public get inferenceClassification(): InferenceClassificationRequestBuilder {
        const builder = new InferenceClassificationRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get mailFolders(): MailFoldersRequestBuilder {
        const builder = new MailFoldersRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get messages(): MessagesRequestBuilder {
        const builder = new MessagesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public readonly mailFoldersById = (id: String) : MailFolderRequestBuilder => {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/mailFolders/" + id;
        return builder;
    }
    public readonly messagesById = (id: String) : MessageRequestBuilder => {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/messages/" + id;
        return builder;
    }
}
