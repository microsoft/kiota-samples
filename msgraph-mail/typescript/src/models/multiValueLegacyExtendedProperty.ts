import type {Entity} from './entity';
import type {Parsable} from '@microsoft/kiota-abstractions';

export interface MultiValueLegacyExtendedProperty extends Entity, Parsable {
    /**
     * A collection of property values.
     */
    value?: string[] | undefined;
}
