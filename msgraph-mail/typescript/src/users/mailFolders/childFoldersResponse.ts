import {MailFolder} from './mailFolder';

export class ChildFoldersResponse {
    public value?: MailFolder[] | undefined;
    public nextLink?: string | undefined;
}
