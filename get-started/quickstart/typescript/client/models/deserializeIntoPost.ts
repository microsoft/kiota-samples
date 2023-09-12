import { type Post } from './post';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function deserializeIntoPost(post: Post | undefined = {} as Post) : Record<string, (node: ParseNode) => void> {
    return {
        "body": n => { post.body = n.getStringValue(); },
        "id": n => { post.id = n.getNumberValue(); },
        "title": n => { post.title = n.getStringValue(); },
        "userId": n => { post.userId = n.getNumberValue(); },
    }
}
