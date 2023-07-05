import {deserializeIntoRecipient} from './deserializeIntoRecipient';
import {Recipient} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createRecipientFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoRecipient;
}
