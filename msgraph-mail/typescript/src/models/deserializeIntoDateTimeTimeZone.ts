import type {DateTimeTimeZone} from './dateTimeTimeZone';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoDateTimeTimeZone(dateTimeTimeZone: DateTimeTimeZone | undefined = {} as DateTimeTimeZone) : Record<string, (node: ParseNode) => void> {
    return {
        "dateTime": n => { dateTimeTimeZone.dateTime = n.getStringValue(); },
        "timeZone": n => { dateTimeTimeZone.timeZone = n.getStringValue(); },
    }
}
