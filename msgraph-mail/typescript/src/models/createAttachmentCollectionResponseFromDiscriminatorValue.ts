<<<<<<< HEAD
import {AttachmentCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createAttachmentCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : AttachmentCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new AttachmentCollectionResponseImpl();
=======
import {AttachmentCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createAttachmentCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : AttachmentCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new AttachmentCollectionResponse();
>>>>>>> main
}
