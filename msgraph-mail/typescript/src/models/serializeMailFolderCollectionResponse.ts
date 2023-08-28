import type {MailFolder} from './mailFolder';
import type {MailFolderCollectionResponse} from './mailFolderCollectionResponse';
import {serializeMailFolder} from './serializeMailFolder';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMailFolderCollectionResponse(writer: SerializationWriter, mailFolderCollectionResponse: MailFolderCollectionResponse | undefined = {} as MailFolderCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", mailFolderCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<MailFolder>("value", mailFolderCollectionResponse.value, serializeMailFolder);
        writer.writeAdditionalData(mailFolderCollectionResponse.additionalData);
}
