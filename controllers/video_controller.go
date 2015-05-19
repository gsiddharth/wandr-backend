package controllers

import (
	"github.com/gorilla/mux"
	"github.com/gsiddharth/wandr/errors"
	"github.com/gsiddharth/wandr/models"
	"github.com/gsiddharth/wandr/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetUserVideos(rw http.ResponseWriter, r *http.Request) {
	if user, err0 := CurrentUser(rw, r); err0 != nil {
		models.NewErrorOutput(errors.Error{Description: "Unauthorized", Code: http.StatusUnauthorized}).Render(rw)
	} else {
		if videos, err1 := models.GetVideosOfUser(user); err1 != nil {
			models.NewErrorOutput(errors.Error{Description: "Error: Videos not found", Code: http.StatusNoContent}).Render(rw)
		} else {
			models.NewOutput(videos, "OK", http.StatusOK).Render(rw)
		}
	}
}

func GetRecommendedVideos(rw http.ResponseWriter, r *http.Request) {
	if user, err00 := CurrentUser(rw, r); err00 != nil {
		models.NewErrorOutput(errors.Error{Description: "Unauthorized", Code: http.StatusUnauthorized}).Render(rw)
	} else {
		city := mux.Vars(r)["city"]
		longitude, _ := strconv.ParseFloat(mux.Vars(r)["longitude"], 64)
		latitude, _ := strconv.ParseFloat(mux.Vars(r)["latitude"], 64)

		location := &models.Location{City: city, Longitude: longitude, Latitude: latitude}

		if videos, err0 := models.GetNearbyVideos(user, location); err0 != nil {
			models.NewErrorOutput(errors.Error{Description: "Error: Videos not found", Code: http.StatusNoContent}).Render(rw)
		} else {
			models.NewOutput(videos, "OK", http.StatusOK).Render(rw)
		}
	}
}

func SaveVideoAndThumbnail(rw http.ResponseWriter, r *http.Request) {
	if user, err00 := CurrentUser(rw, r); err00 != nil {
		models.NewErrorOutput(errors.Error{Description: "Unauthorized", Code: http.StatusUnauthorized})
		return
	} else {

		videoExtension := r.FormValue("video_extension")
		println(videoExtension)
		videoLengthSeconds, _ := strconv.ParseFloat(r.FormValue("video_length"), 10)
		var videoSizeInBytes int64 = 0

		longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 10)
		latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 10)
		city := mux.Vars(r)["city"]
		videoTime := time.Now()
		videoFileName := utils.GetFilePath(utils.BASE_PATH_VIDEOS, utils.VIDEO_FILE_NAME_LENGTH, videoExtension)

		thumbnailExtension := r.FormValue("thumnail_extension")
		thumbnailWidth, _ := strconv.ParseUint(r.FormValue("thumbnail_width"), 10, 64)
		thumbnailHeight, _ := strconv.ParseUint(r.FormValue("thumbnail_height"), 10, 64)
		var thumnailSizeInBytes int64 = 0
		thumbnailFileName := utils.GetFilePath(utils.BASE_PATH_THUMBNAILS, utils.THUMBNAIL_FILE_NAME_LENGTH, thumbnailExtension)

		if formFile, _, err0 := r.FormFile("video_file"); err0 != nil {
			models.NewErrorOutput(errors.Error{Description: "Error: No Video File", Code: http.StatusNoContent}).Render(rw)
			return
		} else {
			defer formFile.Close()

			outFileName := videoFileName

			if osFile, err := os.Create(outFileName); err != nil {
				models.NewErrorOutput(errors.Error{Description: "Error: Video Upload Failed", Code: http.StatusConflict}).Render(rw)
				return
			} else {
				defer osFile.Close()

				if size, err := io.Copy(osFile, formFile); err != nil {
					models.NewErrorOutput(errors.Error{Description: "Error: Video Upload Failed", Code: http.StatusConflict}).Render(rw)
					return
				} else {
					videoSizeInBytes = size
				}

			}
		}

		if formFile, _, err0 := r.FormFile("thumbnail_file"); err0 != nil {
			models.NewErrorOutput(errors.Error{Description: "Error: No Thumbnail file", Code: http.StatusNoContent}).Render(rw)
			return

		} else {
			defer formFile.Close()

			outFileName := thumbnailFileName

			if osFile, err := os.Create(outFileName); err != nil {
				println(err.Error())
				models.NewErrorOutput(errors.Error{Description: "Error: Thumbnail Upload Failed", Code: http.StatusConflict}).Render(rw)
				return

			} else {
				defer osFile.Close()

				if size, err := io.Copy(osFile, formFile); err != nil {
					println(err.Error())
					models.NewErrorOutput(errors.Error{Description: "Error: Thumbnail Upload Failed", Code: http.StatusConflict}).Render(rw)
				} else {
					thumnailSizeInBytes = size
					thumbnails := []models.Thumbnail{models.Thumbnail{Url: thumbnailFileName, Width: uint(thumbnailWidth),
						Height: uint(thumbnailHeight), SizeInBytes: uint(thumnailSizeInBytes)}}

					video := &models.Video{UserID: user.ID, Url: videoFileName, LengthInSeconds: videoLengthSeconds, TimeOfVideo: videoTime,
						SizeInBytes: uint(videoSizeInBytes), Location: models.Location{City: city, Longitude: longitude, Latitude: latitude},
						Thumbnails: thumbnails}

					if _, err0 := models.AddVideo(video); err0 != nil {
						println(err.Error())
						models.NewErrorOutput(errors.Error{Description: "Error: Failed to Save Video", Code: http.StatusConflict}).Render(rw)
						return
					} else {
						models.NewOutput("", "OK", http.StatusOK).Render(rw)
					}
				}

			}
		}

	}
}

func GetThumbnail(rw http.ResponseWriter, r *http.Request) {

}

func GetComments(rw http.ResponseWriter, r *http.Request) {

}

func GetUserVideo(rw http.ResponseWriter, r *http.Request) {

}
