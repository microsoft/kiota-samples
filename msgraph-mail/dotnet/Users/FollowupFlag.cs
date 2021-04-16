using System;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    public class FollowupFlag {
        public DateTimeTimeZone CompletedDateTime { get; set; }
        public DateTimeTimeZone DueDateTime { get; set; }
        public FollowupFlagStatus FlagStatus { get; set; }
        public DateTimeTimeZone StartDateTime { get; set; }
    }
}
