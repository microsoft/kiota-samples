import {ToDoItemCollectionResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createToDoItemCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : ToDoItemCollectionResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ToDoItemCollectionResponse();
}
