import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class SingleValueLegacyExtendedProperty extends Entity implements Parsable<SingleValueLegacyExtendedProperty> {
    public value?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeStringValue("value", this.value);
    };
    public deserializeFields () : Map<string, (item: SingleValueLegacyExtendedProperty, node: ParseNode) => void> {
        return new Map<string, (item: SingleValueLegacyExtendedProperty, node: ParseNode) => void>([...super.deserializeFields(),
            ["value", (o, n) => { o.value = n.getStringValue(); }],
        ]);
    };
}
