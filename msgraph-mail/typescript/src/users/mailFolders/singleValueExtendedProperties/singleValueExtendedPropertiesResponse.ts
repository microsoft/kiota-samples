import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {SingleValueLegacyExtendedProperty} from '../../singleValueLegacyExtendedProperty';

export class SingleValueExtendedPropertiesResponse implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    public nextLink?: string | undefined;
    public value?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["@odata.nextLink", (o, n) => { (o as unknown as SingleValueExtendedPropertiesResponse).nextLink = n.getStringValue(); }],
            ["value", (o, n) => { (o as unknown as SingleValueExtendedPropertiesResponse).value = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("@odata.nextLink", this.nextLink);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("value", this.value);
        writer.writeAdditionalData(this.additionalData);
    };
}
