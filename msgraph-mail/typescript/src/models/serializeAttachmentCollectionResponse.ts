// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { type Attachment } from './attachment';
import { type AttachmentCollectionResponse } from './attachmentCollectionResponse';
import { serializeAttachment } from './serializeAttachment';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function serializeAttachmentCollectionResponse(writer: SerializationWriter, attachmentCollectionResponse: AttachmentCollectionResponse | undefined = {} as AttachmentCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", attachmentCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<Attachment>("value", attachmentCollectionResponse.value, serializeAttachment);
        writer.writeAdditionalData(attachmentCollectionResponse.additionalData);
}
// tslint:enable
// eslint-enable
