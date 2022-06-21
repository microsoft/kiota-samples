import {MultiValueLegacyExtendedPropertyImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMultiValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode: ParseNode | undefined) : MultiValueLegacyExtendedPropertyImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MultiValueLegacyExtendedPropertyImpl();
}
