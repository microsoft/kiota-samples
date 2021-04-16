import {Message} from '../message';

export class MessagesResponse {
    public value?: Message[] | undefined;
    public nextLink?: string | undefined;
}
