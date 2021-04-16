using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    public class OutlookItem : Entity {
        public List<object> Categories { get; set; }
        public string ChangeKey { get; set; }
        public string CreatedDateTime { get; set; }
        public string LastModifiedDateTime { get; set; }
    }
}
