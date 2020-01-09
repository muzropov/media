package media

import "fmt"

const URL  = "https://s3.amazonaws.com"
const BUCKET  = "fiesta.storage-resized"
const PROMO_VIDEO_PATH  = "videos/original"

func getPromoVideoUrl(filename string) string {
	return fmt.Sprintf("%s/%s/%s/%s", URL, BUCKET, PROMO_VIDEO_PATH, filename)
}