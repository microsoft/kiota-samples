package graphjavav4.utilities.users.inferenceClassification;

import graphjavav4.utilities.users.InferenceClassificationType;
import graphjavav4.utilities.users.EmailAddress;
import graphjavav4.utilities.users.Entity;
public class InferenceClassificationOverride extends Entity {
    @javax.annotation.Nullable
    public InferenceClassificationType classifyAs;
    @javax.annotation.Nullable
    public EmailAddress senderEmailAddress;
}
