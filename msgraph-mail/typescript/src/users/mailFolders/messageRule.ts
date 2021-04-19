import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from '../entity';
import {MessageRuleActions} from './messageRuleActions';
import {MessageRulePredicates} from './messageRulePredicates';

export class MessageRule extends Entity implements Parsable<MessageRule> {
    public actions?: MessageRuleActions | undefined;
    public conditions?: MessageRulePredicates | undefined;
    public displayName?: string | undefined;
    public exceptions?: MessageRulePredicates | undefined;
    public hasError?: boolean | undefined;
    public isEnabled?: boolean | undefined;
    public isReadOnly?: boolean | undefined;
    public sequence?: number | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeObjectValue<MessageRuleActions>("actions", this.actions);
        writer.writeObjectValue<MessageRulePredicates>("conditions", this.conditions);
        writer.writeStringValue("displayName", this.displayName);
        writer.writeObjectValue<MessageRulePredicates>("exceptions", this.exceptions);
        writer.writeBooleanValue("hasError", this.hasError);
        writer.writeBooleanValue("isEnabled", this.isEnabled);
        writer.writeBooleanValue("isReadOnly", this.isReadOnly);
        writer.writeNumberValue("sequence", this.sequence);
    };
    public deserializeFields () : Map<string, (item: MessageRule, node: ParseNode) => void> {
        return new Map<string, (item: MessageRule, node: ParseNode) => void>([...super.deserializeFields(),
            ["actions", (o, n) => { o.actions = n.getObjectValue<MessageRuleActions>(MessageRuleActions); }],
            ["conditions", (o, n) => { o.conditions = n.getObjectValue<MessageRulePredicates>(MessageRulePredicates); }],
            ["displayName", (o, n) => { o.displayName = n.getStringValue(); }],
            ["exceptions", (o, n) => { o.exceptions = n.getObjectValue<MessageRulePredicates>(MessageRulePredicates); }],
            ["hasError", (o, n) => { o.hasError = n.getBooleanValue(); }],
            ["isEnabled", (o, n) => { o.isEnabled = n.getBooleanValue(); }],
            ["isReadOnly", (o, n) => { o.isReadOnly = n.getBooleanValue(); }],
            ["sequence", (o, n) => { o.sequence = n.getNumberValue(); }],
        ]);
    };
}
