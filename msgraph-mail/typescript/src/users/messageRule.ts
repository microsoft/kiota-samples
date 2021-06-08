import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';
import {MessageRuleActions} from './mailFolders/messageRuleActions';
import {MessageRulePredicates} from './mailFolders/messageRulePredicates';

export class MessageRule extends Entity implements Parsable {
    public actions?: MessageRuleActions | undefined;
    public conditions?: MessageRulePredicates | undefined;
    /** The display name of the rule.  */
    public displayName?: string | undefined;
    public exceptions?: MessageRulePredicates | undefined;
    /** Indicates whether the rule is in an error condition. Read-only.  */
    public hasError?: boolean | undefined;
    /** Indicates whether the rule is enabled to be applied to messages.  */
    public isEnabled?: boolean | undefined;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.  */
    public isReadOnly?: boolean | undefined;
    /** Indicates the order in which the rule is executed, among other rules.  */
    public sequence?: number | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["actions", (o, n) => { (o as unknown as MessageRule).actions = n.getObjectValue<MessageRuleActions>(MessageRuleActions); }],
            ["conditions", (o, n) => { (o as unknown as MessageRule).conditions = n.getObjectValue<MessageRulePredicates>(MessageRulePredicates); }],
            ["displayName", (o, n) => { (o as unknown as MessageRule).displayName = n.getStringValue(); }],
            ["exceptions", (o, n) => { (o as unknown as MessageRule).exceptions = n.getObjectValue<MessageRulePredicates>(MessageRulePredicates); }],
            ["hasError", (o, n) => { (o as unknown as MessageRule).hasError = n.getBooleanValue(); }],
            ["isEnabled", (o, n) => { (o as unknown as MessageRule).isEnabled = n.getBooleanValue(); }],
            ["isReadOnly", (o, n) => { (o as unknown as MessageRule).isReadOnly = n.getBooleanValue(); }],
            ["sequence", (o, n) => { (o as unknown as MessageRule).sequence = n.getNumberValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
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
}
