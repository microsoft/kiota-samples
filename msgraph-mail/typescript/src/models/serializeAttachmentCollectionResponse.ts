import {Attachment} from './attachment';
import {AttachmentCollectionResponse} from './attachmentCollectionResponse';
import {serializeAttachment} from './serializeAttachment';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeAttachmentCollectionResponse(writer: SerializationWriter, attachmentCollectionResponse: AttachmentCollectionResponse | undefined = {} as AttachmentCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", attachmentCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Attachment>("value", attachmentCollectionResponse.value, serializeAttachment);
        writer.writeAdditionalData(attachmentCollectionResponse.additionalData);
}
