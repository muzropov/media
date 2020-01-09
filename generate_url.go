package media

import "fmt"

const URL  = "https://s3.amazonaws.com"
const BUCKET  = "fiesta.storage-resized"
const PROMO_VIDEO_PATH  = "videos/original"
const RESIZED_VIDEO_PATH  = "videos/720p"

func GetPromoVideoUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, PROMO_VIDEO_PATH, filename)
}

func GetResizedVideoUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, RESIZED_VIDEO_PATH, filename)
}