import type {Entity} from './entity';
import type {MessageRuleActions} from './messageRuleActions';
import type {MessageRulePredicates} from './messageRulePredicates';
import type {Parsable} from '@microsoft/kiota-abstractions';

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
