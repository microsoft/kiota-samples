import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {EmailAddress} from './emailAddress';

export class Recipient implements Parsable<Recipient> {
    public emailAddress?: EmailAddress | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeObjectValue<EmailAddress>("emailAddress", this.emailAddress);
        writer.writeAdditionalData(this.additionalData);
    };
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: Recipient, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: Recipient, node: ParseNode) => void> {
        return new Map<string, (item: Recipient, node: ParseNode) => void>([
            ["emailAddress", (o, n) => { o.emailAddress = n.getObjectValue<EmailAddress>(EmailAddress); }],
        ]);
    };
}
