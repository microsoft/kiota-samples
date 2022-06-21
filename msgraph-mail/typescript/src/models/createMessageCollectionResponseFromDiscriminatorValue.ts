import {MessageCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageCollectionResponseImpl();
}
