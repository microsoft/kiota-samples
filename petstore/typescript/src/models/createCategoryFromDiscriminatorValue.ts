import {Category} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createCategoryFromDiscriminatorValue(parseNode: ParseNode | undefined) : Category {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Category();
}
