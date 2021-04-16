using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages {
    public class SingleValueExtendedPropertiesResponse {
        public List<SingleValueLegacyExtendedProperty> Value { get; set; }
        public string NextLink { get; set; }
    }
}
