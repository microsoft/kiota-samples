import {Attachment, Entity, Extension, Message, MultiValueLegacyExtendedProperty, OutlookItem, SingleValueLegacyExtendedProperty} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) : Entity {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    const mappingValueNode = parseNode.getChildNode("@odata.type");
    if (mappingValueNode) {
        const mappingValue = mappingValueNode.getStringValue();
        if (mappingValue) {
            switch (mappingValue) {
                case "#microsoft.graph.attachment":
                    return new Attachment();
                case "#microsoft.graph.extension":
                    return new Extension();
                case "#microsoft.graph.message":
                    return new Message();
                case "#microsoft.graph.multiValueLegacyExtendedProperty":
                    return new MultiValueLegacyExtendedProperty();
                case "#microsoft.graph.outlookItem":
                    return new OutlookItem();
                case "#microsoft.graph.singleValueLegacyExtendedProperty":
                    return new SingleValueLegacyExtendedProperty();
            }
        }
    }
    return new Entity();
}
