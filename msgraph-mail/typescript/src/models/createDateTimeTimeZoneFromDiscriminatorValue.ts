import {DateTimeTimeZoneImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createDateTimeTimeZoneFromDiscriminatorValue(parseNode: ParseNode | undefined) : DateTimeTimeZoneImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new DateTimeTimeZoneImpl();
}
