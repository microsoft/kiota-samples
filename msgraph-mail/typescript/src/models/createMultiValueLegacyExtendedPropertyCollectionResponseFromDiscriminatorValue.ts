<<<<<<< HEAD
import {MultiValueLegacyExtendedPropertyCollectionResponseImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMultiValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MultiValueLegacyExtendedPropertyCollectionResponseImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MultiValueLegacyExtendedPropertyCollectionResponseImpl();
=======
import {MultiValueLegacyExtendedPropertyCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createMultiValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : MultiValueLegacyExtendedPropertyCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new MultiValueLegacyExtendedPropertyCollectionResponse();
>>>>>>> main
}
