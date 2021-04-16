using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders;
namespace Graphdotnetv4.Users {
    public class MailFoldersResponse {
        public List<MailFolder> Value { get; set; }
        public string NextLink { get; set; }
    }
}
