import type {Entity} from './entity';
import type {Parsable} from '@microsoft/kiota-abstractions';

export interface SingleValueLegacyExtendedProperty extends Entity, Parsable {
    /**
     * A property value.
     */
    value?: string | undefined;
}
