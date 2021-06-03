import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class SingleValueLegacyExtendedProperty extends Entity implements Parsable<SingleValueLegacyExtendedProperty> {
    /** A property value.  */
    public value?: string | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: SingleValueLegacyExtendedProperty, node: ParseNode) => void>
     */
    public getFieldDeserializers () : Map<string, (item: SingleValueLegacyExtendedProperty, node: ParseNode) => void> {
        return new Map<string, (item: SingleValueLegacyExtendedProperty, node: ParseNode) => void>([...super.getFieldDeserializers(),
            ["value", (o, n) => { o.value = n.getStringValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeStringValue("value", this.value);
    };
}
