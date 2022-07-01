import {EntityImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) : EntityImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new EntityImpl();
}
