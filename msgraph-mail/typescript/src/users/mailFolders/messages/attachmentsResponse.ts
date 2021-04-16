import {Attachment} from '../../attachment';

export class AttachmentsResponse {
    public value?: Attachment[] | undefined;
    public nextLink?: string | undefined;
}
