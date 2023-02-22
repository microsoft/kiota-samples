import {ToDoItem} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createToDoItemFromDiscriminatorValue(parseNode: ParseNode | undefined) : ToDoItem {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new ToDoItem();
}
