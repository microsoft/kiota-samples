import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { createMessageRuleActionsFromDiscriminatorValue, serializeMessageRuleActions, type MessageRuleActions } from './messageRuleActions';
import { createMessageRulePredicatesFromDiscriminatorValue, serializeMessageRulePredicates, type MessageRulePredicates } from './messageRulePredicates';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageRuleFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessageRule;
}
export function deserializeIntoMessageRule(messageRule: MessageRule | undefined = {} as MessageRule) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(messageRule),
        "actions": n => { messageRule.actions = n.getObjectValue<MessageRuleActions>(createMessageRuleActionsFromDiscriminatorValue); },
        "conditions": n => { messageRule.conditions = n.getObjectValue<MessageRulePredicates>(createMessageRulePredicatesFromDiscriminatorValue); },
        "displayName": n => { messageRule.displayName = n.getStringValue(); },
        "exceptions": n => { messageRule.exceptions = n.getObjectValue<MessageRulePredicates>(createMessageRulePredicatesFromDiscriminatorValue); },
        "hasError": n => { messageRule.hasError = n.getBooleanValue(); },
        "isEnabled": n => { messageRule.isEnabled = n.getBooleanValue(); },
        "isReadOnly": n => { messageRule.isReadOnly = n.getBooleanValue(); },
        "sequence": n => { messageRule.sequence = n.getNumberValue(); },
    }
}
export interface MessageRule extends Entity, Parsable {
    /**
     * The actions property
     */
    actions?: MessageRuleActions | undefined;
    /**
     * The conditions property
     */
    conditions?: MessageRulePredicates | undefined;
    /**
     * The display name of the rule.
     */
    displayName?: string | undefined;
    /**
     * The exceptions property
     */
    exceptions?: MessageRulePredicates | undefined;
    /**
     * Indicates whether the rule is in an error condition. Read-only.
     */
    hasError?: boolean | undefined;
    /**
     * Indicates whether the rule is enabled to be applied to messages.
     */
    isEnabled?: boolean | undefined;
    /**
     * Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     */
    isReadOnly?: boolean | undefined;
    /**
     * Indicates the order in which the rule is executed, among other rules.
     */
    sequence?: number | undefined;
}
export function serializeMessageRule(writer: SerializationWriter, messageRule: MessageRule | undefined = {} as MessageRule) : void {
        serializeEntity(writer, messageRule)
        writer.writeObjectValue<MessageRuleActions>("actions", messageRule.actions, );
        writer.writeObjectValue<MessageRulePredicates>("conditions", messageRule.conditions, );
        writer.writeStringValue("displayName", messageRule.displayName);
        writer.writeObjectValue<MessageRulePredicates>("exceptions", messageRule.exceptions, );
        writer.writeBooleanValue("hasError", messageRule.hasError);
        writer.writeBooleanValue("isEnabled", messageRule.isEnabled);
        writer.writeBooleanValue("isReadOnly", messageRule.isReadOnly);
        writer.writeNumberValue("sequence", messageRule.sequence);
}
