import {SingleValueLegacyExtendedPropertyCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createSingleValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : SingleValueLegacyExtendedPropertyCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new SingleValueLegacyExtendedPropertyCollectionResponseImpl();
}
