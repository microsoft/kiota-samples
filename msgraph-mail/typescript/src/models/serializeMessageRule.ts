import type {MessageRule} from './messageRule';
import type {MessageRuleActions} from './messageRuleActions';
import type {MessageRulePredicates} from './messageRulePredicates';
import {serializeEntity} from './serializeEntity';
import {serializeMessageRuleActions} from './serializeMessageRuleActions';
import {serializeMessageRulePredicates} from './serializeMessageRulePredicates';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMessageRule(writer: SerializationWriter, messageRule: MessageRule | undefined = {} as MessageRule) : void {
        serializeEntity(writer, messageRule)
        writer.writeObjectValue<MessageRuleActions>("actions", messageRule.actions, serializeMessageRuleActions);
        writer.writeObjectValue<MessageRulePredicates>("conditions", messageRule.conditions, serializeMessageRulePredicates);
        writer.writeStringValue("displayName", messageRule.displayName);
        writer.writeObjectValue<MessageRulePredicates>("exceptions", messageRule.exceptions, serializeMessageRulePredicates);
        writer.writeBooleanValue("hasError", messageRule.hasError);
        writer.writeBooleanValue("isEnabled", messageRule.isEnabled);
        writer.writeBooleanValue("isReadOnly", messageRule.isReadOnly);
        writer.writeNumberValue("sequence", messageRule.sequence);
}
