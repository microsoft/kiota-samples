import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {EmailAddress} from './emailAddress';

export class Recipient implements Parsable<Recipient> {
    public emailAddress?: EmailAddress | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeObjectValue<EmailAddress>("emailAddress", this.emailAddress);
    };
    public deserializeFields () : Map<string, (item: Recipient, node: ParseNode) => void> {
        return new Map<string, (item: Recipient, node: ParseNode) => void>([
            ["emailAddress", (o, n) => { o.emailAddress = n.getObjectValue<EmailAddress>(EmailAddress); }],
        ]);
    };
}
