import {EmailAddress} from './emailAddress';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeEmailAddress(writer: SerializationWriter, emailAddress: EmailAddress | undefined = {} as EmailAddress) : void {
        writer.writeStringValue("address", emailAddress.address);
        writer.writeStringValue("name", emailAddress.name);
        writer.writeAdditionalData(emailAddress.additionalData);
}
