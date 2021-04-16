using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Messages {
    public class MultiValueExtendedPropertiesResponse {
        public List<MultiValueLegacyExtendedProperty> Value { get; set; }
        public string NextLink { get; set; }
    }
}
