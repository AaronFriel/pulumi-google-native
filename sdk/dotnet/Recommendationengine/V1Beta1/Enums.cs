// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.Recommendationengine.V1Beta1
{
    /// <summary>
    /// Optional. Online stock state of the catalog item. Default is `IN_STOCK`.
    /// </summary>
    [EnumType]
    public readonly struct GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState : IEquatable<GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState>
    {
        private readonly string _value;

        private GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default item stock status. Should never be used.
        /// </summary>
        public static GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState StockStateUnspecified { get; } = new GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState("STOCK_STATE_UNSPECIFIED");
        /// <summary>
        /// Item in stock.
        /// </summary>
        public static GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState InStock { get; } = new GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState("IN_STOCK");
        /// <summary>
        /// Item out of stock.
        /// </summary>
        public static GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState OutOfStock { get; } = new GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState("OUT_OF_STOCK");
        /// <summary>
        /// Item that is in pre-order state.
        /// </summary>
        public static GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState Preorder { get; } = new GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState("PREORDER");
        /// <summary>
        /// Item that is back-ordered (i.e. temporarily out of stock).
        /// </summary>
        public static GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState Backorder { get; } = new GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState("BACKORDER");

        public static bool operator ==(GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState left, GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState right) => left.Equals(right);
        public static bool operator !=(GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState left, GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState right) => !left.Equals(right);

        public static explicit operator string(GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState other && Equals(other);
        public bool Equals(GoogleCloudRecommendationengineV1beta1ProductCatalogItemStockState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}