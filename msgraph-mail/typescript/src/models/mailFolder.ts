<<<<<<< HEAD
import {Entity} from './entity';
=======
import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {createMessageFromDiscriminatorValue} from './createMessageFromDiscriminatorValue';
import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {EntityImpl, MessageImpl, MessageRuleImpl, MultiValueLegacyExtendedPropertyImpl, SingleValueLegacyExtendedPropertyImpl} from './index';
import {MailFolder} from './mailFolder';
>>>>>>> main
import {Message} from './message';
import {MessageRule} from './messageRule';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
<<<<<<< HEAD
import {Parsable} from '@microsoft/kiota-abstractions';

export interface MailFolder extends Entity, Partial<Parsable> {
    /** The number of immediate child mailFolders in the current mailFolder. */
    childFolderCount?: number | undefined;
    /** The collection of child folders in the mailFolder. */
    childFolders?: MailFolder[] | undefined;
    /** The mailFolder's display name. */
    displayName?: string | undefined;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. */
    isHidden?: boolean | undefined;
    /** The collection of rules that apply to the user's Inbox folder. */
    messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder. */
    messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. */
    multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder. */
    parentFolderId?: string | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. */
    singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The number of items in the mailFolder. */
    totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread. */
    unreadItemCount?: number | undefined;
=======
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MailFolderImpl extends EntityImpl implements MailFolder {
    /** The number of immediate child mailFolders in the current mailFolder. */
    private _childFolderCount?: number | undefined;
    /** The collection of child folders in the mailFolder. */
    private _childFolders?: MailFolder[] | undefined;
    /** The mailFolder's display name. */
    private _displayName?: string | undefined;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. */
    private _isHidden?: boolean | undefined;
    /** The collection of rules that apply to the user's Inbox folder. */
    private _messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder. */
    private _messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. */
    private _multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder. */
    private _parentFolderId?: string | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. */
    private _singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The number of items in the mailFolder. */
    private _totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread. */
    private _unreadItemCount?: number | undefined;
    /**
     * Gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
     * @returns a integer
     */
    public get childFolderCount() {
        return this._childFolderCount;
    };
    /**
     * Sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
     * @param value Value to set for the childFolderCount property.
     */
    public set childFolderCount(value: number | undefined) {
        if(value) {
            this._childFolderCount = value;
        }
    };
    /**
     * Gets the childFolders property value. The collection of child folders in the mailFolder.
     * @returns a MailFolderInterface
     */
    public get childFolders() {
        return this._childFolders;
    };
    /**
     * Sets the childFolders property value. The collection of child folders in the mailFolder.
     * @param value Value to set for the childFolders property.
     */
    public set childFolders(value: MailFolder[] | undefined) {
        if(value) {
            const childFoldersArrValue: MailFolderImpl[] = [];
            this.childFolders?.forEach(element => {
                childFoldersArrValue.push((element instanceof MailFolderImpl? element as MailFolderImpl:new MailFolderImpl(element)));
            });
            this._childFolders = childFoldersArrValue;
        }
    };
    /**
     * Instantiates a new mailFolder and sets the default values.
     * @param mailFolderParameterValue 
     */
    public constructor(mailFolderParameterValue?: MailFolder | undefined) {
        super(mailFolderParameterValue);
        this._childFolderCount = mailFolderParameterValue?.childFolderCount;
        this._childFolders = mailFolderParameterValue?.childFolders;
        this._displayName = mailFolderParameterValue?.displayName;
        this._isHidden = mailFolderParameterValue?.isHidden;
        this._messageRules = mailFolderParameterValue?.messageRules;
        this._messages = mailFolderParameterValue?.messages;
        this._multiValueExtendedProperties = mailFolderParameterValue?.multiValueExtendedProperties;
        this._parentFolderId = mailFolderParameterValue?.parentFolderId;
        this._singleValueExtendedProperties = mailFolderParameterValue?.singleValueExtendedProperties;
        this._totalItemCount = mailFolderParameterValue?.totalItemCount;
        this._unreadItemCount = mailFolderParameterValue?.unreadItemCount;
    };
    /**
     * Gets the displayName property value. The mailFolder's display name.
     * @returns a string
     */
    public get displayName() {
        return this._displayName;
    };
    /**
     * Sets the displayName property value. The mailFolder's display name.
     * @param value Value to set for the displayName property.
     */
    public set displayName(value: string | undefined) {
        if(value) {
            this._displayName = value;
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "childFolderCount": n => { this.childFolderCount = n.getNumberValue(); },
            "childFolders": n => { this.childFolders = n.getCollectionOfObjectValues<MailFolderImpl>(createMailFolderFromDiscriminatorValue); },
            "displayName": n => { this.displayName = n.getStringValue(); },
            "isHidden": n => { this.isHidden = n.getBooleanValue(); },
            "messageRules": n => { this.messageRules = n.getCollectionOfObjectValues<MessageRuleImpl>(createMessageRuleFromDiscriminatorValue); },
            "messages": n => { this.messages = n.getCollectionOfObjectValues<MessageImpl>(createMessageFromDiscriminatorValue); },
            "multiValueExtendedProperties": n => { this.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
            "parentFolderId": n => { this.parentFolderId = n.getStringValue(); },
            "singleValueExtendedProperties": n => { this.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>(createSingleValueLegacyExtendedPropertyFromDiscriminatorValue); },
            "totalItemCount": n => { this.totalItemCount = n.getNumberValue(); },
            "unreadItemCount": n => { this.unreadItemCount = n.getNumberValue(); },
        };
    };
    /**
     * Gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
     * @returns a boolean
     */
    public get isHidden() {
        return this._isHidden;
    };
    /**
     * Sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
     * @param value Value to set for the isHidden property.
     */
    public set isHidden(value: boolean | undefined) {
        if(value) {
            this._isHidden = value;
        }
    };
    /**
     * Gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
     * @returns a MessageRuleInterface
     */
    public get messageRules() {
        return this._messageRules;
    };
    /**
     * Sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
     * @param value Value to set for the messageRules property.
     */
    public set messageRules(value: MessageRule[] | undefined) {
        if(value) {
            const messageRulesArrValue: MessageRuleImpl[] = [];
            this.messageRules?.forEach(element => {
                messageRulesArrValue.push((element instanceof MessageRuleImpl? element as MessageRuleImpl:new MessageRuleImpl(element)));
            });
            this._messageRules = messageRulesArrValue;
        }
    };
    /**
     * Gets the messages property value. The collection of messages in the mailFolder.
     * @returns a MessageInterface
     */
    public get messages() {
        return this._messages;
    };
    /**
     * Sets the messages property value. The collection of messages in the mailFolder.
     * @param value Value to set for the messages property.
     */
    public set messages(value: Message[] | undefined) {
        if(value) {
            const messagesArrValue: MessageImpl[] = [];
            this.messages?.forEach(element => {
                messagesArrValue.push((element instanceof MessageImpl? element as MessageImpl:new MessageImpl(element)));
            });
            this._messages = messagesArrValue;
        }
    };
    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @returns a MultiValueLegacyExtendedPropertyInterface
     */
    public get multiValueExtendedProperties() {
        return this._multiValueExtendedProperties;
    };
    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param value Value to set for the multiValueExtendedProperties property.
     */
    public set multiValueExtendedProperties(value: MultiValueLegacyExtendedProperty[] | undefined) {
        if(value) {
            const multiValueExtendedPropertiesArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
            this.multiValueExtendedProperties?.forEach(element => {
                multiValueExtendedPropertiesArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
            });
            this._multiValueExtendedProperties = multiValueExtendedPropertiesArrValue;
        }
    };
    /**
     * Gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
     * @returns a string
     */
    public get parentFolderId() {
        return this._parentFolderId;
    };
    /**
     * Sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
     * @param value Value to set for the parentFolderId property.
     */
    public set parentFolderId(value: string | undefined) {
        if(value) {
            this._parentFolderId = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.childFolderCount){
            writer.writeNumberValue("childFolderCount", this.childFolderCount);
        }
        if(this.childFolders && this.childFolders.length != 0){        const childFoldersArrValue: MailFolderImpl[] = [];
        this.childFolders?.forEach(element => {
            childFoldersArrValue.push((element instanceof MailFolderImpl? element as MailFolderImpl:new MailFolderImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MailFolderImpl>("childFolders", childFoldersArrValue);
        }
        if(this.displayName){
            writer.writeStringValue("displayName", this.displayName);
        }
        if(this.isHidden){
            writer.writeBooleanValue("isHidden", this.isHidden);
        }
        if(this.messageRules && this.messageRules.length != 0){        const messageRulesArrValue: MessageRuleImpl[] = [];
        this.messageRules?.forEach(element => {
            messageRulesArrValue.push((element instanceof MessageRuleImpl? element as MessageRuleImpl:new MessageRuleImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MessageRuleImpl>("messageRules", messageRulesArrValue);
        }
        if(this.messages && this.messages.length != 0){        const messagesArrValue: MessageImpl[] = [];
        this.messages?.forEach(element => {
            messagesArrValue.push((element instanceof MessageImpl? element as MessageImpl:new MessageImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MessageImpl>("messages", messagesArrValue);
        }
        if(this.multiValueExtendedProperties && this.multiValueExtendedProperties.length != 0){        const multiValueExtendedPropertiesArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
        this.multiValueExtendedProperties?.forEach(element => {
            multiValueExtendedPropertiesArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>("multiValueExtendedProperties", multiValueExtendedPropertiesArrValue);
        }
        if(this.parentFolderId){
            writer.writeStringValue("parentFolderId", this.parentFolderId);
        }
        if(this.singleValueExtendedProperties && this.singleValueExtendedProperties.length != 0){        const singleValueExtendedPropertiesArrValue: SingleValueLegacyExtendedPropertyImpl[] = [];
        this.singleValueExtendedProperties?.forEach(element => {
            singleValueExtendedPropertiesArrValue.push((element instanceof SingleValueLegacyExtendedPropertyImpl? element as SingleValueLegacyExtendedPropertyImpl:new SingleValueLegacyExtendedPropertyImpl(element)));
        });
            writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>("singleValueExtendedProperties", singleValueExtendedPropertiesArrValue);
        }
        if(this.totalItemCount){
            writer.writeNumberValue("totalItemCount", this.totalItemCount);
        }
        if(this.unreadItemCount){
            writer.writeNumberValue("unreadItemCount", this.unreadItemCount);
        }
    };
    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @returns a SingleValueLegacyExtendedPropertyInterface
     */
    public get singleValueExtendedProperties() {
        return this._singleValueExtendedProperties;
    };
    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param value Value to set for the singleValueExtendedProperties property.
     */
    public set singleValueExtendedProperties(value: SingleValueLegacyExtendedProperty[] | undefined) {
        if(value) {
            const singleValueExtendedPropertiesArrValue: SingleValueLegacyExtendedPropertyImpl[] = [];
            this.singleValueExtendedProperties?.forEach(element => {
                singleValueExtendedPropertiesArrValue.push((element instanceof SingleValueLegacyExtendedPropertyImpl? element as SingleValueLegacyExtendedPropertyImpl:new SingleValueLegacyExtendedPropertyImpl(element)));
            });
            this._singleValueExtendedProperties = singleValueExtendedPropertiesArrValue;
        }
    };
    /**
     * Gets the totalItemCount property value. The number of items in the mailFolder.
     * @returns a integer
     */
    public get totalItemCount() {
        return this._totalItemCount;
    };
    /**
     * Sets the totalItemCount property value. The number of items in the mailFolder.
     * @param value Value to set for the totalItemCount property.
     */
    public set totalItemCount(value: number | undefined) {
        if(value) {
            this._totalItemCount = value;
        }
    };
    /**
     * Gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
     * @returns a integer
     */
    public get unreadItemCount() {
        return this._unreadItemCount;
    };
    /**
     * Sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
     * @param value Value to set for the unreadItemCount property.
     */
    public set unreadItemCount(value: number | undefined) {
        if(value) {
            this._unreadItemCount = value;
        }
    };
>>>>>>> main
}
