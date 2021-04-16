using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages {
    public class AttachmentsResponse {
        public List<Attachment> Value { get; set; }
        public string NextLink { get; set; }
    }
}
