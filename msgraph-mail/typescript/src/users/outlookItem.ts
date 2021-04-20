import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class OutlookItem extends Entity implements Parsable<OutlookItem> {
    /** The categories associated with the item  */
    public categories?: string[] | undefined;
    /** Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.  */
    public changeKey?: string | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    public createdDateTime?: Date | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    public lastModifiedDateTime?: Date | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues<string>("categories", this.categories);
        writer.writeStringValue("changeKey", this.changeKey);
        writer.writeDateValue("createdDateTime", this.createdDateTime);
        writer.writeDateValue("lastModifiedDateTime", this.lastModifiedDateTime);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: OutlookItem, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: OutlookItem, node: ParseNode) => void> {
        return new Map<string, (item: OutlookItem, node: ParseNode) => void>([...super.deserializeFields(),
            ["categories", (o, n) => { o.categories = n.getCollectionOfPrimitiveValues<string>(); }],
            ["changeKey", (o, n) => { o.changeKey = n.getStringValue(); }],
            ["createdDateTime", (o, n) => { o.createdDateTime = n.getDateValue(); }],
            ["lastModifiedDateTime", (o, n) => { o.lastModifiedDateTime = n.getDateValue(); }],
        ]);
    };
}
