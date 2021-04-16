import {Entity} from './entity';

export class OutlookItem extends Entity {
    public categories?: object[] | undefined;
    public changeKey?: string | undefined;
    public createdDateTime?: string | undefined;
    public lastModifiedDateTime?: string | undefined;
}
