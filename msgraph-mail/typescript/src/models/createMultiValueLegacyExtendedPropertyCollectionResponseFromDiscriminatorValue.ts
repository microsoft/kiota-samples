import {MultiValueLegacyExtendedPropertyCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMultiValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MultiValueLegacyExtendedPropertyCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MultiValueLegacyExtendedPropertyCollectionResponseImpl();
}
