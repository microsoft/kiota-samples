import { Importance } from './importance';
import { createRecipientFromDiscriminatorValue, serializeRecipient, type Recipient } from './recipient';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageRuleActionsFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessageRuleActions;
}
export function deserializeIntoMessageRuleActions(messageRuleActions: MessageRuleActions | undefined = {} as MessageRuleActions) : Record<string, (node: ParseNode) => void> {
    return {
        "assignCategories": n => { messageRuleActions.assignCategories = n.getCollectionOfPrimitiveValues<string>(); },
        "copyToFolder": n => { messageRuleActions.copyToFolder = n.getStringValue(); },
        "delete": n => { messageRuleActions.delete = n.getBooleanValue(); },
        "forwardAsAttachmentTo": n => { messageRuleActions.forwardAsAttachmentTo = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "forwardTo": n => { messageRuleActions.forwardTo = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "markAsRead": n => { messageRuleActions.markAsRead = n.getBooleanValue(); },
        "markImportance": n => { messageRuleActions.markImportance = n.getEnumValue<Importance>(Importance); },
        "moveToFolder": n => { messageRuleActions.moveToFolder = n.getStringValue(); },
        "permanentDelete": n => { messageRuleActions.permanentDelete = n.getBooleanValue(); },
        "redirectTo": n => { messageRuleActions.redirectTo = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "stopProcessingRules": n => { messageRuleActions.stopProcessingRules = n.getBooleanValue(); },
    }
}
export interface MessageRuleActions extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * A list of categories to be assigned to a message.
     */
    assignCategories?: string[] | undefined;
    /**
     * The ID of a folder that a message is to be copied to.
     */
    copyToFolder?: string | undefined;
    /**
     * Indicates whether a message should be moved to the Deleted Items folder.
     */
    delete?: boolean | undefined;
    /**
     * The email addresses of the recipients to which a message should be forwarded as an attachment.
     */
    forwardAsAttachmentTo?: Recipient[] | undefined;
    /**
     * The email addresses of the recipients to which a message should be forwarded.
     */
    forwardTo?: Recipient[] | undefined;
    /**
     * Indicates whether a message should be marked as read.
     */
    markAsRead?: boolean | undefined;
    /**
     * The markImportance property
     */
    markImportance?: Importance | undefined;
    /**
     * The ID of the folder that a message will be moved to.
     */
    moveToFolder?: string | undefined;
    /**
     * Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     */
    permanentDelete?: boolean | undefined;
    /**
     * The email addresses to which a message should be redirected.
     */
    redirectTo?: Recipient[] | undefined;
    /**
     * Indicates whether subsequent rules should be evaluated.
     */
    stopProcessingRules?: boolean | undefined;
}
export function serializeMessageRuleActions(writer: SerializationWriter, messageRuleActions: MessageRuleActions | undefined = {} as MessageRuleActions) : void {
        writer.writeCollectionOfPrimitiveValues<string>("assignCategories", messageRuleActions.assignCategories);
        writer.writeStringValue("copyToFolder", messageRuleActions.copyToFolder);
        writer.writeBooleanValue("delete", messageRuleActions.delete);
        writer.writeCollectionOfObjectValues<Recipient>("forwardAsAttachmentTo", messageRuleActions.forwardAsAttachmentTo, );
        writer.writeCollectionOfObjectValues<Recipient>("forwardTo", messageRuleActions.forwardTo, );
        writer.writeBooleanValue("markAsRead", messageRuleActions.markAsRead);
        writer.writeEnumValue<Importance>("markImportance", messageRuleActions.markImportance);
        writer.writeStringValue("moveToFolder", messageRuleActions.moveToFolder);
        writer.writeBooleanValue("permanentDelete", messageRuleActions.permanentDelete);
        writer.writeCollectionOfObjectValues<Recipient>("redirectTo", messageRuleActions.redirectTo, );
        writer.writeBooleanValue("stopProcessingRules", messageRuleActions.stopProcessingRules);
        writer.writeAdditionalData(messageRuleActions.additionalData);
}
