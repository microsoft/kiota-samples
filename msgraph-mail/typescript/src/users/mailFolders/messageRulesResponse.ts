import {MessageRule} from './messageRule';

export class MessageRulesResponse {
    public value?: MessageRule[] | undefined;
    public nextLink?: string | undefined;
}
