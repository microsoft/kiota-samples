using System;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationOverride : Entity {
        public InferenceClassificationType ClassifyAs { get; set; }
        public EmailAddress SenderEmailAddress { get; set; }
    }
}
