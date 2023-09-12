import { deserializeIntoPost } from './deserializeIntoPost';
import { type Post } from './index';
import { type ParseNode } from '@microsoft/kiota-abstractions';

export function createPostFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoPost;
}
