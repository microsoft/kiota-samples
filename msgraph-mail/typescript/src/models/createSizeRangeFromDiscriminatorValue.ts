import {SizeRangeImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createSizeRangeFromDiscriminatorValue(parseNode: ParseNode | undefined) : SizeRangeImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new SizeRangeImpl();
}
