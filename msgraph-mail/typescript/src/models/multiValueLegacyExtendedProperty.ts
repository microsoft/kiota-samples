import {Entity} from './entity';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface MultiValueLegacyExtendedProperty extends Entity, Parsable {
    /**
     * A collection of property values.
     */
    value?: string[] | undefined;
}
