import {Entity} from './entity';

export class Attachment extends Entity {
    public contentType?: string | undefined;
    public isInline?: boolean | undefined;
    public lastModifiedDateTime?: string | undefined;
    public name?: string | undefined;
    public size?: number | undefined;
}
