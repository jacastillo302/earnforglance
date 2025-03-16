package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	CollectionMediaSettings = "media_settings"
)

// MediaSettings represents media settings
type MediaSettings struct {
	ID                                          primitive.ObjectID `bson:"_id,omitempty"`
	AvatarPictureSize                           int                `bson:"avatar_picture_size"`
	ProductThumbPictureSize                     int                `bson:"product_thumb_picture_size"`
	ProductDetailsPictureSize                   int                `bson:"product_details_picture_size"`
	ProductThumbPictureSizeOnProductDetailsPage int                `bson:"product_thumb_picture_size_on_product_details_page"`
	AssociatedProductPictureSize                int                `bson:"associated_product_picture_size"`
	CategoryThumbPictureSize                    int                `bson:"category_thumb_picture_size"`
	ManufacturerThumbPictureSize                int                `bson:"manufacturer_thumb_picture_size"`
	VendorThumbPictureSize                      int                `bson:"vendor_thumb_picture_size"`
	CartThumbPictureSize                        int                `bson:"cart_thumb_picture_size"`
	OrderThumbPictureSize                       int                `bson:"order_thumb_picture_size"`
	MiniCartThumbPictureSize                    int                `bson:"mini_cart_thumb_picture_size"`
	AutoCompleteSearchThumbPictureSize          int                `bson:"auto_complete_search_thumb_picture_size"`
	ImageSquarePictureSize                      int                `bson:"image_square_picture_size"`
	DefaultPictureZoomEnabled                   bool               `bson:"default_picture_zoom_enabled"`
	AllowSVGUploads                             bool               `bson:"allow_svg_uploads"`
	MaximumImageSize                            int                `bson:"maximum_image_size"`
	DefaultImageQuality                         int                `bson:"default_image_quality"`
	MultipleThumbDirectories                    bool               `bson:"multiple_thumb_directories"`
	ImportProductImagesUsingHash                bool               `bson:"import_product_images_using_hash"`
	AzureCacheControlHeader                     string             `bson:"azure_cache_control_header"`
	UseAbsoluteImagePath                        bool               `bson:"use_absolute_image_path"`
	VideoIframeAllow                            string             `bson:"video_iframe_allow"`
	VideoIframeWidth                            int                `bson:"video_iframe_width"`
	VideoIframeHeight                           int                `bson:"video_iframe_height"`
	ProductDefaultImageID                       int                `bson:"product_default_image_id"`
	AutoOrientImage                             bool               `bson:"auto_orient_image"`
}
