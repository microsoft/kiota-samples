import { createMailFolderFromDiscriminatorValue, serializeMailFolder, type MailFolder } from './mailFolder';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMailFolderCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMailFolderCollectionResponse;
}
export function deserializeIntoMailFolderCollectionResponse(mailFolderCollectionResponse: MailFolderCollectionResponse | undefined = {} as MailFolderCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { mailFolderCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { mailFolderCollectionResponse.value = n.getCollectionOfObjectValues<MailFolder>(createMailFolderFromDiscriminatorValue); },
    }
}
export interface MailFolderCollectionResponse extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The OdataNextLink property
     */
    odataNextLink?: string | undefined;
    /**
     * The value property
     */
    value?: MailFolder[] | undefined;
}
export function serializeMailFolderCollectionResponse(writer: SerializationWriter, mailFolderCollectionResponse: MailFolderCollectionResponse | undefined = {} as MailFolderCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", mailFolderCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<MailFolder>("value", mailFolderCollectionResponse.value, );
        writer.writeAdditionalData(mailFolderCollectionResponse.additionalData);
}
