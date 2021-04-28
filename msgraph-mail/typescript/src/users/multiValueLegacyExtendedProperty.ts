import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class MultiValueLegacyExtendedProperty extends Entity implements Parsable<MultiValueLegacyExtendedProperty> {
    /** A collection of property values.  */
    public value?: Array[] | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: MultiValueLegacyExtendedProperty, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: MultiValueLegacyExtendedProperty, node: ParseNode) => void> {
        return new Map<string, (item: MultiValueLegacyExtendedProperty, node: ParseNode) => void>([...super.deserializeFields(),
            ["value", (o, n) => { o.value = n.getCollectionOfPrimitiveValues<array>(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues<array>("value", this.value);
    };
}
