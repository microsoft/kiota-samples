import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {SingleValueLegacyExtendedProperty} from '../singleValueLegacyExtendedProperty';

export class SingleValueExtendedPropertiesResponse implements Parsable<SingleValueExtendedPropertiesResponse> {
    public value?: SingleValueLegacyExtendedProperty[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: SingleValueExtendedPropertiesResponse, node: ParseNode) => void> {
        return new Map<string, (item: SingleValueExtendedPropertiesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
