import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {Importance} from './importance';
import {RecipientImpl} from './index';
import {MessageRuleActions} from './messageRuleActions';
import {Recipient} from './recipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRuleActionsImpl implements MessageRuleActions {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** A list of categories to be assigned to a message. */
    public assignCategories?: string[] | undefined;
    /** The ID of a folder that a message is to be copied to. */
    public copyToFolder?: string | undefined;
    /** Indicates whether a message should be moved to the Deleted Items folder. */
    public delete?: boolean | undefined;
    /** The email addresses of the recipients to which a message should be forwarded as an attachment. */
    public forwardAsAttachmentTo?: Recipient[] | undefined;
    /** The email addresses of the recipients to which a message should be forwarded. */
    public forwardTo?: Recipient[] | undefined;
    /** Indicates whether a message should be marked as read. */
    public markAsRead?: boolean | undefined;
    /** The markImportance property */
    public markImportance?: Importance | undefined;
    /** The ID of the folder that a message will be moved to. */
    public moveToFolder?: string | undefined;
    /** Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder. */
    public permanentDelete?: boolean | undefined;
    /** The email address to which a message should be redirected. */
    public redirectTo?: Recipient[] | undefined;
    /** Indicates whether subsequent rules should be evaluated. */
    public stopProcessingRules?: boolean | undefined;
    /**
     * Instantiates a new messageRuleActions and sets the default values.
     * @param messageRuleActionsParameterValue 
     */
    public constructor(messageRuleActionsParameterValue?: MessageRuleActions | undefined) {
        this.additionalData = messageRuleActionsParameterValue?.additionalData ? messageRuleActionsParameterValue?.additionalData! : {};
        this.assignCategories = messageRuleActionsParameterValue?.assignCategories;
        this.copyToFolder = messageRuleActionsParameterValue?.copyToFolder;
        this.delete = messageRuleActionsParameterValue?.delete;
        this.forwardAsAttachmentTo = messageRuleActionsParameterValue?.forwardAsAttachmentTo;
        this.forwardTo = messageRuleActionsParameterValue?.forwardTo;
        this.markAsRead = messageRuleActionsParameterValue?.markAsRead;
        this.markImportance = messageRuleActionsParameterValue?.markImportance;
        this.moveToFolder = messageRuleActionsParameterValue?.moveToFolder;
        this.permanentDelete = messageRuleActionsParameterValue?.permanentDelete;
        this.redirectTo = messageRuleActionsParameterValue?.redirectTo;
        this.stopProcessingRules = messageRuleActionsParameterValue?.stopProcessingRules;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "assignCategories": n => { this.assignCategories = n.getCollectionOfPrimitiveValues<string>(); },
            "copyToFolder": n => { this.copyToFolder = n.getStringValue(); },
            "delete": n => { this.delete = n.getBooleanValue(); },
            "forwardAsAttachmentTo": n => { this.forwardAsAttachmentTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "forwardTo": n => { this.forwardTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "markAsRead": n => { this.markAsRead = n.getBooleanValue(); },
            "markImportance": n => { this.markImportance = n.getEnumValue<Importance>(Importance); },
            "moveToFolder": n => { this.moveToFolder = n.getStringValue(); },
            "permanentDelete": n => { this.permanentDelete = n.getBooleanValue(); },
            "redirectTo": n => { this.redirectTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "stopProcessingRules": n => { this.stopProcessingRules = n.getBooleanValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.assignCategories){
            writer.writeCollectionOfPrimitiveValues<string>("assignCategories", this.assignCategories);
        }
        if(this.copyToFolder){
            writer.writeStringValue("copyToFolder", this.copyToFolder);
        }
        if(this.delete){
            writer.writeBooleanValue("delete", this.delete);
        }
        if(this.forwardAsAttachmentTo && this.forwardAsAttachmentTo.length != 0){        const forwardAsAttachmentToArrValue: RecipientImpl[] = []; this.forwardAsAttachmentTo?.forEach(element => {forwardAsAttachmentToArrValue.push(new RecipientImpl(element));});
            writer.writeCollectionOfObjectValues<RecipientImpl>("forwardAsAttachmentTo", forwardAsAttachmentToArrValue);
        }
        if(this.forwardTo && this.forwardTo.length != 0){        const forwardToArrValue: RecipientImpl[] = []; this.forwardTo?.forEach(element => {forwardToArrValue.push(new RecipientImpl(element));});
            writer.writeCollectionOfObjectValues<RecipientImpl>("forwardTo", forwardToArrValue);
        }
        if(this.markAsRead){
            writer.writeBooleanValue("markAsRead", this.markAsRead);
        }
        if(this.markImportance){
            writer.writeEnumValue<Importance>("markImportance", this.markImportance);
        }
        if(this.moveToFolder){
            writer.writeStringValue("moveToFolder", this.moveToFolder);
        }
        if(this.permanentDelete){
            writer.writeBooleanValue("permanentDelete", this.permanentDelete);
        }
        if(this.redirectTo && this.redirectTo.length != 0){        const redirectToArrValue: RecipientImpl[] = []; this.redirectTo?.forEach(element => {redirectToArrValue.push(new RecipientImpl(element));});
            writer.writeCollectionOfObjectValues<RecipientImpl>("redirectTo", redirectToArrValue);
        }
        if(this.stopProcessingRules){
            writer.writeBooleanValue("stopProcessingRules", this.stopProcessingRules);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
