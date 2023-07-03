import {deserializeIntoAttachmentCollectionResponse} from './deserializeIntoAttachmentCollectionResponse';
import {AttachmentCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createAttachmentCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoAttachmentCollectionResponse;
}
