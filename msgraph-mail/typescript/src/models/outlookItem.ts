import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createOutlookItemFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoOutlookItem;
}
export function deserializeIntoOutlookItem(outlookItem: OutlookItem | undefined = {} as OutlookItem) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(outlookItem),
        "categories": n => { outlookItem.categories = n.getCollectionOfPrimitiveValues<string>(); },
        "changeKey": n => { outlookItem.changeKey = n.getStringValue(); },
        "createdDateTime": n => { outlookItem.createdDateTime = n.getDateValue(); },
        "lastModifiedDateTime": n => { outlookItem.lastModifiedDateTime = n.getDateValue(); },
    }
}
export interface OutlookItem extends Entity, Parsable {
    /**
     * The categories associated with the item
     */
    categories?: string[] | undefined;
    /**
     * Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     */
    changeKey?: string | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    createdDateTime?: Date | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    lastModifiedDateTime?: Date | undefined;
}
export function serializeOutlookItem(writer: SerializationWriter, outlookItem: OutlookItem | undefined = {} as OutlookItem) : void {
        serializeEntity(writer, outlookItem)
        writer.writeCollectionOfPrimitiveValues<string>("categories", outlookItem.categories);
        writer.writeStringValue("changeKey", outlookItem.changeKey);
        writer.writeDateValue("createdDateTime", outlookItem.createdDateTime);
        writer.writeDateValue("lastModifiedDateTime", outlookItem.lastModifiedDateTime);
}
