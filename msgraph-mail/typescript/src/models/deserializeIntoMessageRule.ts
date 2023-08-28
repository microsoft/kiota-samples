import {createMessageRuleActionsFromDiscriminatorValue} from './createMessageRuleActionsFromDiscriminatorValue';
import {createMessageRulePredicatesFromDiscriminatorValue} from './createMessageRulePredicatesFromDiscriminatorValue';
import {deserializeIntoEntity} from './deserializeIntoEntity';
import type {MessageRule} from './messageRule';
import type {MessageRuleActions} from './messageRuleActions';
import type {MessageRulePredicates} from './messageRulePredicates';
import {serializeMessageRuleActions} from './serializeMessageRuleActions';
import {serializeMessageRulePredicates} from './serializeMessageRulePredicates';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

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
