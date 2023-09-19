import { createEmailAddressFromDiscriminatorValue, serializeEmailAddress, type EmailAddress } from './emailAddress';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createRecipientFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoRecipient;
}
export function deserializeIntoRecipient(recipient: Recipient | undefined = {} as Recipient) : Record<string, (node: ParseNode) => void> {
    return {
        "emailAddress": n => { recipient.emailAddress = n.getObjectValue<EmailAddress>(createEmailAddressFromDiscriminatorValue); },
    }
}
export interface Recipient extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The emailAddress property
     */
    emailAddress?: EmailAddress | undefined;
}
export function serializeRecipient(writer: SerializationWriter, recipient: Recipient | undefined = {} as Recipient) : void {
        writer.writeObjectValue<EmailAddress>("emailAddress", recipient.emailAddress, );
        writer.writeAdditionalData(recipient.additionalData);
}
