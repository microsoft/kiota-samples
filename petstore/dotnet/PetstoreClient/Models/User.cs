using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace KiotaSamples.PetStoreSdk.Models {
    public class User : IParsable {
        /// <summary>The email property</summary>
        public string Email { get; set; }
        /// <summary>The firstName property</summary>
        public string FirstName { get; set; }
        /// <summary>The id property</summary>
        public long? Id { get; set; }
        /// <summary>The lastName property</summary>
        public string LastName { get; set; }
        /// <summary>The password property</summary>
        public string Password { get; set; }
        /// <summary>The phone property</summary>
        public string Phone { get; set; }
        /// <summary>The username property</summary>
        public string Username { get; set; }
        /// <summary>User Status</summary>
        public int? UserStatus { get; set; }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        /// </summary>
        public static User CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new User();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"email", n => { Email = n.GetStringValue(); } },
                {"firstName", n => { FirstName = n.GetStringValue(); } },
                {"id", n => { Id = n.GetLongValue(); } },
                {"lastName", n => { LastName = n.GetStringValue(); } },
                {"password", n => { Password = n.GetStringValue(); } },
                {"phone", n => { Phone = n.GetStringValue(); } },
                {"username", n => { Username = n.GetStringValue(); } },
                {"userStatus", n => { UserStatus = n.GetIntValue(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteStringValue("email", Email);
            writer.WriteStringValue("firstName", FirstName);
            writer.WriteLongValue("id", Id);
            writer.WriteStringValue("lastName", LastName);
            writer.WriteStringValue("password", Password);
            writer.WriteStringValue("phone", Phone);
            writer.WriteStringValue("username", Username);
            writer.WriteIntValue("userStatus", UserStatus);
        }
    }
}
