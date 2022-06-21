import {AttachmentImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createAttachmentFromDiscriminatorValue(parseNode: ParseNode | undefined) : AttachmentImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new AttachmentImpl();
}
