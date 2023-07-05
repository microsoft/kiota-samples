import {Importance} from './importance';
import {MessageRuleActions} from './messageRuleActions';
import {Recipient} from './recipient';
import {serializeRecipient} from './serializeRecipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMessageRuleActions(writer: SerializationWriter, messageRuleActions: MessageRuleActions | undefined = {} as MessageRuleActions) : void {
        writer.writeCollectionOfPrimitiveValues<string>("assignCategories", messageRuleActions.assignCategories);
        writer.writeStringValue("copyToFolder", messageRuleActions.copyToFolder);
        writer.writeBooleanValue("delete", messageRuleActions.delete);
        writer.writeCollectionOfObjectValues<Recipient>("forwardAsAttachmentTo", messageRuleActions.forwardAsAttachmentTo, serializeRecipient);
        writer.writeCollectionOfObjectValues<Recipient>("forwardTo", messageRuleActions.forwardTo, serializeRecipient);
        writer.writeBooleanValue("markAsRead", messageRuleActions.markAsRead);
        writer.writeEnumValue<Importance>("markImportance", messageRuleActions.markImportance);
        writer.writeStringValue("moveToFolder", messageRuleActions.moveToFolder);
        writer.writeBooleanValue("permanentDelete", messageRuleActions.permanentDelete);
        writer.writeCollectionOfObjectValues<Recipient>("redirectTo", messageRuleActions.redirectTo, serializeRecipient);
        writer.writeBooleanValue("stopProcessingRules", messageRuleActions.stopProcessingRules);
        writer.writeAdditionalData(messageRuleActions.additionalData);
}
