import {MessageRuleCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageRuleCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageRuleCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageRuleCollectionResponseImpl();
}
