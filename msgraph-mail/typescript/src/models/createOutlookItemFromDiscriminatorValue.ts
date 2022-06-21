import {OutlookItemImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createOutlookItemFromDiscriminatorValue(parseNode: ParseNode | undefined) : OutlookItemImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new OutlookItemImpl();
}
