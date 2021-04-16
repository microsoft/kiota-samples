using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Graphdotnetv4.Users.InferenceClassification;
using Graphdotnetv4.Users.MailFolders;
using Graphdotnetv4.Users.MailFolders.Messages;
namespace Graphdotnetv4.Users.Item {
    public class UserRequestBuilder {
        public InferenceClassificationRequestBuilder InferenceClassification { get =>
            new InferenceClassificationRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MailFoldersRequestBuilder MailFolders { get =>
            new MailFoldersRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MessagesRequestBuilder Messages { get =>
            new MessagesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        private string PathSegment { get; } = "";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
    }
}
