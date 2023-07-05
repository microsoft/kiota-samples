import {deserializeIntoOutlookItem} from './deserializeIntoOutlookItem';
import {OutlookItem} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createOutlookItemFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoOutlookItem;
}
