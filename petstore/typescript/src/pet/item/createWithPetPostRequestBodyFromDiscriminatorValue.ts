import {WithPetPostRequestBody} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createWithPetPostRequestBodyFromDiscriminatorValue(parseNode: ParseNode | undefined) : WithPetPostRequestBody {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new WithPetPostRequestBody();
}
