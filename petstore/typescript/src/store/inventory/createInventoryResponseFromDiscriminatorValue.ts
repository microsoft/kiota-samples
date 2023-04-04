import {InventoryResponse} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createInventoryResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) : InventoryResponse {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new InventoryResponse();
}
