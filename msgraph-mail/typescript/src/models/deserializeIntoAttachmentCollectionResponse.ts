import {Attachment} from './attachment';
import {AttachmentCollectionResponse} from './attachmentCollectionResponse';
import {createAttachmentFromDiscriminatorValue} from './createAttachmentFromDiscriminatorValue';
import {serializeAttachment} from './serializeAttachment';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoAttachmentCollectionResponse(attachmentCollectionResponse: AttachmentCollectionResponse | undefined = {} as AttachmentCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { attachmentCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { attachmentCollectionResponse.value = n.getCollectionOfObjectValues<Attachment>(createAttachmentFromDiscriminatorValue); },
    }
}
