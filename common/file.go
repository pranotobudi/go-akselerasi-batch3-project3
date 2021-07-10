package common

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadFile(userID uint, file multipart.File, filename string) (string, error) {
	cld, err := cloudinary.NewFromParams("domtvglu8", "523282955746113", "jA9Q7DdnReBJEj9Co1jyEPphxVk")
	// Upload the my_picture.jpg image and set the public_id to "my_image".
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	var ctx = context.Background()

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: "authorID-" + fmt.Sprintf("%d", userID) + "-" + filename})
	// Get details about the image with public id "my_image" and log the secure URL.
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	log.Println(resp.SecureURL)

	asset, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "authorID-" + fmt.Sprintf("%d", userID) + "-" + filename})
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	// Print some basic information about the asset.
	log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
	return resp.SecureURL, nil
}
