import {Order} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createOrderFromDiscriminatorValue(parseNode: ParseNode | undefined) : Order {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Order();
}
