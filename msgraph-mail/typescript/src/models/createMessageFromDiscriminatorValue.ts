import {MessageImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageImpl();
}
