import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class MultiValueLegacyExtendedProperty extends Entity implements Parsable<MultiValueLegacyExtendedProperty> {
    public value?: object[] | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues<object>("value", this.value);
    };
    public deserializeFields () : Map<string, (item: MultiValueLegacyExtendedProperty, node: ParseNode) => void> {
        return new Map<string, (item: MultiValueLegacyExtendedProperty, node: ParseNode) => void>([...super.DeserializeFields(),
            ["value", (o, n) => { o.value = n.getCollectionOfPrimitiveValues<object>(); }],
        ]);
    };
}
