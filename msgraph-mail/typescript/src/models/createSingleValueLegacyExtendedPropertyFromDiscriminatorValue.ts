import {SingleValueLegacyExtendedPropertyImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createSingleValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode: ParseNode | undefined) : SingleValueLegacyExtendedPropertyImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new SingleValueLegacyExtendedPropertyImpl();
}
