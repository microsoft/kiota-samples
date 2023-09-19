import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createEmailAddressFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoEmailAddress;
}
export function deserializeIntoEmailAddress(emailAddress: EmailAddress | undefined = {} as EmailAddress) : Record<string, (node: ParseNode) => void> {
    return {
        "address": n => { emailAddress.address = n.getStringValue(); },
        "name": n => { emailAddress.name = n.getStringValue(); },
    }
}
export interface EmailAddress extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The email address of the person or entity.
     */
    address?: string | undefined;
    /**
     * The display name of the person or entity.
     */
    name?: string | undefined;
}
export function serializeEmailAddress(writer: SerializationWriter, emailAddress: EmailAddress | undefined = {} as EmailAddress) : void {
        writer.writeStringValue("address", emailAddress.address);
        writer.writeStringValue("name", emailAddress.name);
        writer.writeAdditionalData(emailAddress.additionalData);
}
