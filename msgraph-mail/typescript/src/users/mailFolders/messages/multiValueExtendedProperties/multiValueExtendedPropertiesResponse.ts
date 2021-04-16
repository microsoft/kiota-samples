import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {MultiValueLegacyExtendedProperty} from '../../../multiValueLegacyExtendedProperty';

export class MultiValueExtendedPropertiesResponse implements Parsable<MultiValueExtendedPropertiesResponse> {
    public value?: MultiValueLegacyExtendedProperty[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: MultiValueExtendedPropertiesResponse, node: ParseNode) => void> {
        return new Map<string, (item: MultiValueExtendedPropertiesResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(MultiValueLegacyExtendedProperty); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
