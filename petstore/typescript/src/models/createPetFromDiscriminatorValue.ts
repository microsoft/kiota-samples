import {Pet} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createPetFromDiscriminatorValue(parseNode: ParseNode | undefined) : Pet {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Pet();
}
