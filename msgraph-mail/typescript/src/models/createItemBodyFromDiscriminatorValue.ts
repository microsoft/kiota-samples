import {ItemBodyImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createItemBodyFromDiscriminatorValue(parseNode: ParseNode | undefined) : ItemBodyImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ItemBodyImpl();
}
