import type {DateTimeTimeZone} from './dateTimeTimeZone';
import type {FollowupFlag} from './followupFlag';
import {FollowupFlagStatus} from './followupFlagStatus';
import {serializeDateTimeTimeZone} from './serializeDateTimeTimeZone';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeFollowupFlag(writer: SerializationWriter, followupFlag: FollowupFlag | undefined = {} as FollowupFlag) : void {
        writer.writeObjectValue<DateTimeTimeZone>("completedDateTime", followupFlag.completedDateTime, serializeDateTimeTimeZone);
        writer.writeObjectValue<DateTimeTimeZone>("dueDateTime", followupFlag.dueDateTime, serializeDateTimeTimeZone);
        writer.writeEnumValue<FollowupFlagStatus>("flagStatus", followupFlag.flagStatus);
        writer.writeObjectValue<DateTimeTimeZone>("startDateTime", followupFlag.startDateTime, serializeDateTimeTimeZone);
        writer.writeAdditionalData(followupFlag.additionalData);
}
