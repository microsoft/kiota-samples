using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    public class MultiValueLegacyExtendedProperty : Entity {
        public List<object> Value { get; set; }
    }
}
