// ignore_for_file: type=lint
import 'package:microsoft_kiota_abstractions/microsoft_kiota_abstractions.dart';

/// auto generated
class Post implements AdditionalDataHolder, Parsable {
    ///  Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    @override
    Map<String, Object?> additionalData;
    ///  The body property
    String? body;
    ///  The id property
    int? id;
    ///  The title property
    String? title;
    ///  The userId property
    int? userId;
    /// Instantiates a new [Post] and sets the default values.
    Post() :  
        additionalData = {};
    /// Creates a new instance of the appropriate class based on discriminator value
    ///  [parseNode] The parse node to use to read the discriminator value and create the object
    static Post createFromDiscriminatorValue(ParseNode parseNode) {
        return Post();
    }
    /// The deserialization information for the current model
    @override
    Map<String, void Function(ParseNode)> getFieldDeserializers() {
        var deserializerMap = <String, void Function(ParseNode)>{};
        deserializerMap['body'] = (node) => body = node.getStringValue();
        deserializerMap['id'] = (node) => id = node.getIntValue();
        deserializerMap['title'] = (node) => title = node.getStringValue();
        deserializerMap['userId'] = (node) => userId = node.getIntValue();
        return deserializerMap;
    }
    /// Serializes information the current object
    ///  [writer] Serialization writer to use to serialize this model
    @override
    void serialize(SerializationWriter writer) {
        writer.writeStringValue('body', body);
        writer.writeIntValue('id', id);
        writer.writeStringValue('title', title);
        writer.writeIntValue('userId', userId);
        writer.writeAdditionalData(additionalData);
    }
}
