import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Extension} from '../../../extension';

export class ExtensionsResponse implements Parsable<ExtensionsResponse> {
    public value?: Extension[] | undefined;
    public nextLink?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfObjectValues<Extension>("value", this.value);
        writer.writeStringValue("nextLink", this.nextLink);
    };
    public deserializeFields () : Map<string, (item: ExtensionsResponse, node: ParseNode) => void> {
        return new Map<string, (item: ExtensionsResponse, node: ParseNode) => void>([
            ["value", (o, n) => { o.value = n.getCollectionOfObjectValues<Extension>(Extension); }],
            ["nextLink", (o, n) => { o.nextLink = n.getStringValue(); }],
        ]);
    };
}
