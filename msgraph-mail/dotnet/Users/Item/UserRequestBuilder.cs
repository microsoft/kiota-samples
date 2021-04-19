using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.InferenceClassification;
using Graphdotnetv4.Users.MailFolders;
using Graphdotnetv4.Users.MailFolders.Messages;
namespace Graphdotnetv4.Users.Item {
    public class UserRequestBuilder {
        public InferenceClassificationRequestBuilder InferenceClassification { get =>
            new InferenceClassificationRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public MailFoldersRequestBuilder MailFolders { get =>
            new MailFoldersRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public MessagesRequestBuilder Messages { get =>
            new MessagesRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        private string PathSegment { get; } = "";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
    }
}
