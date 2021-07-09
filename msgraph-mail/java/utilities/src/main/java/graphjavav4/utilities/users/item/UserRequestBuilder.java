package graphjavav4.utilities.users.item;

import com.microsoft.kiota.HttpCore;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestInfo;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.inferenceClassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.mailFolders.item.MailFolderRequestBuilder;
import graphjavav4.utilities.users.mailFolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.messages.MessagesRequestBuilder;
import java.io.InputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.function.Function;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}  */
public class UserRequestBuilder {
    /** Current path for the request  */
    @javax.annotation.Nullable
    public String currentPath;
    /** Core service to use to execute the requests  */
    @javax.annotation.Nullable
    public HttpCore httpCore;
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        return new InferenceClassificationRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        return new MailFoldersRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment;
        final HttpCore parentCore = httpCore;
        return new MessagesRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @return a void
     */
    public UserRequestBuilder() {
        this.pathSegment = "";
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.mailFolders collection
     * @param id Unique identifier of the item
     * @return a MailFolderRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/mailFolders/" + id;
        final HttpCore parentCore = httpCore;
        return new MailFolderRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.messages collection
     * @param id Unique identifier of the item
     * @return a MessageRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        final String parentPath = (currentPath == null ? "" : currentPath) + pathSegment + "/messages/" + id;
        final HttpCore parentCore = httpCore;
        return new MessageRequestBuilder() {{ currentPath = parentPath; httpCore = parentCore; }};
    }
}
