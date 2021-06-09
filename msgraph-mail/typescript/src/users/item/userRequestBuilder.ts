import {HttpCore, HttpMethod, RequestInfo, ResponseHandler, SerializationWriterFactory} from '@microsoft/kiota-abstractions';
import {InferenceClassificationRequestBuilder} from '../inferenceClassification/inferenceClassificationRequestBuilder';
import {MailFolderRequestBuilder} from '../mailFolders/item/mailFolderRequestBuilder';
import {MailFoldersRequestBuilder} from '../mailFolders/mailFoldersRequestBuilder';
import {MessageRequestBuilder} from '../messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../messages/messagesRequestBuilder';

/** Builds and executes requests for operations under /users/{user-id}  */
export class UserRequestBuilder {
    /** Current path for the request  */
    private _currentPath?: string | undefined;
    /** Core service to use to execute the requests  */
    private _httpCore?: HttpCore | undefined;
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
    private readonly _pathSegment: string;
    /** Factory to use to get a serializer for payload serialization  */
    private _serializerFactory?: SerializationWriterFactory | undefined;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     */
    public constructor() {
        this._pathSegment = "";
    };
    /**
     * Gets the currentPath property value. Current path for the request
     * @returns a string
     */
    public get currentPath() {
        return this._currentPath;
    };
    /**
     * Gets the httpCore property value. Core service to use to execute the requests
     * @returns a HttpCore
     */
    public get httpCore() {
        return this._httpCore;
    };
    /**
     * Gets the pathSegment property value. Path segment to use to build the URL for the current request builder
     * @returns a string
     */
    public get pathSegment() {
        return this._pathSegment;
    };
    /**
     * Gets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @returns a SerializationWriterFactory
     */
    public get serializerFactory() {
        return this._serializerFactory;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.mailFolders collection
     * @param id Unique identifier of the item
     * @returns a MailFolderRequestBuilder
     */
    public mailFoldersById(id: String) : MailFolderRequestBuilder {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/mailFolders/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.messages collection
     * @param id Unique identifier of the item
     * @returns a MessageRequestBuilder
     */
    public messagesById(id: String) : MessageRequestBuilder {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath ?? '') + this.pathSegment + "/messages/" + id;
        builder.httpCore = this.httpCore;
        builder.serializerFactory = this.serializerFactory;
        return builder;
    };
    /**
     * Sets the currentPath property value. Current path for the request
     * @param value Value to set for the currentPath property.
     */
    public set currentPath(value: string | undefined) {
        this._currentPath = value;
    };
    /**
     * Sets the httpCore property value. Core service to use to execute the requests
     * @param value Value to set for the httpCore property.
     */
    public set httpCore(value: HttpCore | undefined) {
        this._httpCore = value;
    };
    /**
     * Sets the serializerFactory property value. Factory to use to get a serializer for payload serialization
     * @param value Value to set for the serializerFactory property.
     */
    public set serializerFactory(value: SerializationWriterFactory | undefined) {
        this._serializerFactory = value;
    };
}
