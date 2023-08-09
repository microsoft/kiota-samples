import type {Post} from './post';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializePost(writer: SerializationWriter, post: Post | undefined = {} as Post) : void {
        writer.writeStringValue("body", post.body);
        writer.writeNumberValue("id", post.id);
        writer.writeStringValue("title", post.title);
        writer.writeNumberValue("userId", post.userId);
        writer.writeAdditionalData(post.additionalData);
}
