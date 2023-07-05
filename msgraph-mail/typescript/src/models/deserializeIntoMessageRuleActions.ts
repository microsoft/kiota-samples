import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {Importance} from './importance';
import {MessageRuleActions} from './messageRuleActions';
import {Recipient} from './recipient';
import {serializeRecipient} from './serializeRecipient';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

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
