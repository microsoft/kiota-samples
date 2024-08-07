// <auto-generated/>
using Microsoft.Kiota.Abstractions.Extensions;
using Microsoft.Kiota.Abstractions.Serialization;
using System.Collections.Generic;
using System.IO;
using System;
namespace Graphdotnetv4.Models
{
    [global::System.CodeDom.Compiler.GeneratedCode("Kiota", "1.16.0")]
    #pragma warning disable CS1591
    public partial class FollowupFlag : IAdditionalDataHolder, IParsable
    #pragma warning restore CS1591
    {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The completedDateTime property</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public global::Graphdotnetv4.Models.DateTimeTimeZone? CompletedDateTime { get; set; }
#nullable restore
#else
        public global::Graphdotnetv4.Models.DateTimeTimeZone CompletedDateTime { get; set; }
#endif
        /// <summary>The dueDateTime property</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public global::Graphdotnetv4.Models.DateTimeTimeZone? DueDateTime { get; set; }
#nullable restore
#else
        public global::Graphdotnetv4.Models.DateTimeTimeZone DueDateTime { get; set; }
#endif
        /// <summary>The flagStatus property</summary>
        public global::Graphdotnetv4.Models.FollowupFlagStatus? FlagStatus { get; set; }
        /// <summary>The startDateTime property</summary>
#if NETSTANDARD2_1_OR_GREATER || NETCOREAPP3_1_OR_GREATER
#nullable enable
        public global::Graphdotnetv4.Models.DateTimeTimeZone? StartDateTime { get; set; }
#nullable restore
#else
        public global::Graphdotnetv4.Models.DateTimeTimeZone StartDateTime { get; set; }
#endif
        /// <summary>
        /// Instantiates a new <see cref="global::Graphdotnetv4.Models.FollowupFlag"/> and sets the default values.
        /// </summary>
        public FollowupFlag()
        {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// </summary>
        /// <returns>A <see cref="global::Graphdotnetv4.Models.FollowupFlag"/></returns>
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        public static global::Graphdotnetv4.Models.FollowupFlag CreateFromDiscriminatorValue(IParseNode parseNode)
        {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new global::Graphdotnetv4.Models.FollowupFlag();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        /// <returns>A IDictionary&lt;string, Action&lt;IParseNode&gt;&gt;</returns>
        public virtual IDictionary<string, Action<IParseNode>> GetFieldDeserializers()
        {
            return new Dictionary<string, Action<IParseNode>>
            {
                { "completedDateTime", n => { CompletedDateTime = n.GetObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>(global::Graphdotnetv4.Models.DateTimeTimeZone.CreateFromDiscriminatorValue); } },
                { "dueDateTime", n => { DueDateTime = n.GetObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>(global::Graphdotnetv4.Models.DateTimeTimeZone.CreateFromDiscriminatorValue); } },
                { "flagStatus", n => { FlagStatus = n.GetEnumValue<global::Graphdotnetv4.Models.FollowupFlagStatus>(); } },
                { "startDateTime", n => { StartDateTime = n.GetObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>(global::Graphdotnetv4.Models.DateTimeTimeZone.CreateFromDiscriminatorValue); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// </summary>
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        public virtual void Serialize(ISerializationWriter writer)
        {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>("completedDateTime", CompletedDateTime);
            writer.WriteObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>("dueDateTime", DueDateTime);
            writer.WriteEnumValue<global::Graphdotnetv4.Models.FollowupFlagStatus>("flagStatus", FlagStatus);
            writer.WriteObjectValue<global::Graphdotnetv4.Models.DateTimeTimeZone>("startDateTime", StartDateTime);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
