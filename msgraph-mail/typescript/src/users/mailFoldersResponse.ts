import {MailFolder} from './mailFolders/mailFolder';

export class MailFoldersResponse {
    public value?: MailFolder[] | undefined;
    public nextLink?: string | undefined;
}
