import {Message, OutlookItem} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createOutlookItemFromDiscriminatorValue(parseNode: ParseNode | undefined) : OutlookItem {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    const mappingValueNode = parseNode.getChildNode("@odata.type");
    if (mappingValueNode) {
        const mappingValue = mappingValueNode.getStringValue();
        if (mappingValue) {
            switch (mappingValue) {
                case "#microsoft.graph.message":
                    return new Message();
                case "#microsoft.graph.outlookItem":
                    return new OutlookItem();
            }
        }
    }
    return new OutlookItem();
}
