import {DateTimeTimeZone} from './dateTimeTimeZone';
import {FollowupFlagStatus} from './followupFlagStatus';

export class FollowupFlag {
    public completedDateTime?: DateTimeTimeZone | undefined;
    public dueDateTime?: DateTimeTimeZone | undefined;
    public flagStatus?: FollowupFlagStatus | undefined;
    public startDateTime?: DateTimeTimeZone | undefined;
}
