package graphjavav4.utilities.users.item;

import java.util.Objects;
import graphjavav4.utilities.users.inferenceClassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.mailFolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.MessagesRequestBuilder;
import graphjavav4.utilities.users.mailFolders.item.MailFolderRequestBuilder;
import graphjavav4.utilities.users.mailFolders.messages.item.MessageRequestBuilder;
import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.QueryParametersBase;
import java.util.Map;
import java.net.URI;
import java.net.URISyntaxException;
public class UserRequestBuilder {
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new InferenceClassificationRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new MailFoldersRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        return new MessagesRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    private final String pathSegment = "";
    @javax.annotation.Nullable
    public String currentPath;
    @javax.annotation.Nullable
    public HttpCore httpCore;
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/mailFolders/" + id;
        return new MailFolderRequestBuilder() {{ currentPath = parentPath; }};
    }
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/messages/" + id;
        return new MessageRequestBuilder() {{ currentPath = parentPath; }};
    }
}
