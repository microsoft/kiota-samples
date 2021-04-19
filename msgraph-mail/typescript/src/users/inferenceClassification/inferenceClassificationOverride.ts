import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {EmailAddress} from '../emailAddress';
import {Entity} from '../entity';
import {InferenceClassificationType} from '../inferenceClassificationType';

export class InferenceClassificationOverride extends Entity implements Parsable<InferenceClassificationOverride> {
    public classifyAs?: InferenceClassificationType | undefined;
    public senderEmailAddress?: EmailAddress | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeObjectValue<InferenceClassificationType>("classifyAs", this.classifyAs);
        writer.writeObjectValue<EmailAddress>("senderEmailAddress", this.senderEmailAddress);
    };
    public deserializeFields () : Map<string, (item: InferenceClassificationOverride, node: ParseNode) => void> {
        return new Map<string, (item: InferenceClassificationOverride, node: ParseNode) => void>([...super.deserializeFields(),
            ["classifyAs", (o, n) => { o.classifyAs = n.getObjectValue<InferenceClassificationType>(InferenceClassificationType); }],
            ["senderEmailAddress", (o, n) => { o.senderEmailAddress = n.getObjectValue<EmailAddress>(EmailAddress); }],
        ]);
    };
}
