import {MessageRulePredicatesImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageRulePredicatesFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageRulePredicatesImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageRulePredicatesImpl();
}
