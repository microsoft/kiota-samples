using System;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    public class Attachment : Entity {
        public string ContentType { get; set; }
        public bool IsInline { get; set; }
        public string LastModifiedDateTime { get; set; }
        public string Name { get; set; }
        public int Size { get; set; }
    }
}
