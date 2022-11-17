<<<<<<< HEAD
import {MessageCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageCollectionResponseImpl();
=======
import {MessageCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageCollectionResponse();
>>>>>>> main
}
