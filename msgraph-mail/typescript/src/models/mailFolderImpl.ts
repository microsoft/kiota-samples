import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {createMessageFromDiscriminatorValue} from './createMessageFromDiscriminatorValue';
import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {EntityImpl, MessageImpl, MessageRuleImpl, MultiValueLegacyExtendedPropertyImpl, SingleValueLegacyExtendedPropertyImpl} from './index';
import {MailFolder} from './mailFolder';
import {Message} from './message';
import {MessageRule} from './messageRule';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MailFolderImpl extends EntityImpl implements MailFolder {
    /** The number of immediate child mailFolders in the current mailFolder. */
    public childFolderCount?: number | undefined;
    /** The collection of child folders in the mailFolder. */
    public childFolders?: MailFolder[] | undefined;
    /** The mailFolder's display name. */
    public displayName?: string | undefined;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. */
    public isHidden?: boolean | undefined;
    /** The collection of rules that apply to the user's Inbox folder. */
    public messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder. */
    public messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. */
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder. */
    public parentFolderId?: string | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. */
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The number of items in the mailFolder. */
    public totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread. */
    public unreadItemCount?: number | undefined;
    /**
     * Instantiates a new mailFolder and sets the default values.
     * @param mailFolderParameterValue 
     */
    public constructor(mailFolderParameterValue?: MailFolder | undefined) {
        super(mailFolderParameterValue);
        this.childFolderCount = mailFolderParameterValue?.childFolderCount;
        this.childFolders = mailFolderParameterValue?.childFolders;
        this.displayName = mailFolderParameterValue?.displayName;
        this.isHidden = mailFolderParameterValue?.isHidden;
        this.messageRules = mailFolderParameterValue?.messageRules;
        this.messages = mailFolderParameterValue?.messages;
        this.multiValueExtendedProperties = mailFolderParameterValue?.multiValueExtendedProperties;
        this.parentFolderId = mailFolderParameterValue?.parentFolderId;
        this.singleValueExtendedProperties = mailFolderParameterValue?.singleValueExtendedProperties;
        this.totalItemCount = mailFolderParameterValue?.totalItemCount;
        this.unreadItemCount = mailFolderParameterValue?.unreadItemCount;
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
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.childFolderCount){
            writer.writeNumberValue("childFolderCount", this.childFolderCount);
        }
        if(this.childFolders && this.childFolders.length != 0){        const childFoldersArrValue: MailFolderImpl[] = []; this.childFolders?.forEach(element => {childFoldersArrValue.push(new MailFolderImpl(element));});
            writer.writeCollectionOfObjectValues<MailFolderImpl>("childFolders", childFoldersArrValue);
        }
        if(this.displayName){
            writer.writeStringValue("displayName", this.displayName);
        }
        if(this.isHidden){
            writer.writeBooleanValue("isHidden", this.isHidden);
        }
        if(this.messageRules && this.messageRules.length != 0){        const messageRulesArrValue: MessageRuleImpl[] = []; this.messageRules?.forEach(element => {messageRulesArrValue.push(new MessageRuleImpl(element));});
            writer.writeCollectionOfObjectValues<MessageRuleImpl>("messageRules", messageRulesArrValue);
        }
        if(this.messages && this.messages.length != 0){        const messagesArrValue: MessageImpl[] = []; this.messages?.forEach(element => {messagesArrValue.push(new MessageImpl(element));});
            writer.writeCollectionOfObjectValues<MessageImpl>("messages", messagesArrValue);
        }
        if(this.multiValueExtendedProperties && this.multiValueExtendedProperties.length != 0){        const multiValueExtendedPropertiesArrValue: MultiValueLegacyExtendedPropertyImpl[] = []; this.multiValueExtendedProperties?.forEach(element => {multiValueExtendedPropertiesArrValue.push(new MultiValueLegacyExtendedPropertyImpl(element));});
            writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>("multiValueExtendedProperties", multiValueExtendedPropertiesArrValue);
        }
        if(this.parentFolderId){
            writer.writeStringValue("parentFolderId", this.parentFolderId);
        }
        if(this.singleValueExtendedProperties && this.singleValueExtendedProperties.length != 0){        const singleValueExtendedPropertiesArrValue: SingleValueLegacyExtendedPropertyImpl[] = []; this.singleValueExtendedProperties?.forEach(element => {singleValueExtendedPropertiesArrValue.push(new SingleValueLegacyExtendedPropertyImpl(element));});
            writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>("singleValueExtendedProperties", singleValueExtendedPropertiesArrValue);
        }
        if(this.totalItemCount){
            writer.writeNumberValue("totalItemCount", this.totalItemCount);
        }
        if(this.unreadItemCount){
            writer.writeNumberValue("unreadItemCount", this.unreadItemCount);
        }
    };
}
