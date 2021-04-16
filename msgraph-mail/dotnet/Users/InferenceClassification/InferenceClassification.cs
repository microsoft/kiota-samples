using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassification : Entity {
        public List<InferenceClassificationOverride> Overrides { get; set; }
    }
}
