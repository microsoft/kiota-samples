import {MessageRuleImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMessageRuleFromDiscriminatorValue(parseNode: ParseNode | undefined) : MessageRuleImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MessageRuleImpl();
}
