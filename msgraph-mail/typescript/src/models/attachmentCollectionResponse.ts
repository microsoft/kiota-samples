import { createAttachmentFromDiscriminatorValue, serializeAttachment, type Attachment } from './attachment';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export interface AttachmentCollectionResponse extends AdditionalDataHolder, Parsable {
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
    value?: Attachment[] | undefined;
}
export function createAttachmentCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoAttachmentCollectionResponse;
}
export function deserializeIntoAttachmentCollectionResponse(attachmentCollectionResponse: AttachmentCollectionResponse | undefined = {} as AttachmentCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { attachmentCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { attachmentCollectionResponse.value = n.getCollectionOfObjectValues<Attachment>(createAttachmentFromDiscriminatorValue); },
    }
}
export function serializeAttachmentCollectionResponse(writer: SerializationWriter, attachmentCollectionResponse: AttachmentCollectionResponse | undefined = {} as AttachmentCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", attachmentCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Attachment>("value", attachmentCollectionResponse.value, );
        writer.writeAdditionalData(attachmentCollectionResponse.additionalData);
}
