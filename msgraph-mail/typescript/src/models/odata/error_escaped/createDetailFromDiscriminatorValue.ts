import {Detail} from './detail';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createDetailFromDiscriminatorValue(parseNode: ParseNode | undefined) : Detail {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new Detail();
}
