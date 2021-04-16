import {Entity} from '../entity';
import {MessageRuleActions} from './messageRuleActions';
import {MessageRulePredicates} from './messageRulePredicates';

export class MessageRule extends Entity {
    public actions?: MessageRuleActions | undefined;
    public conditions?: MessageRulePredicates | undefined;
    public displayName?: string | undefined;
    public exceptions?: MessageRulePredicates | undefined;
    public hasError?: boolean | undefined;
    public isEnabled?: boolean | undefined;
    public isReadOnly?: boolean | undefined;
    public sequence?: number | undefined;
}
