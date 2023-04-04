import {Tag} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createTagFromDiscriminatorValue(parseNode: ParseNode | undefined) : Tag {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Tag();
}
