import {FollowupFlagImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createFollowupFlagFromDiscriminatorValue(parseNode: ParseNode | undefined) : FollowupFlagImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new FollowupFlagImpl();
}
