// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks views
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// RegisterResult is the viewed result type that is projected based on a view.
type RegisterResult struct {
	// Type to project
	Projected *RegisterResultView
	// View to render
	View string
}

// Task is the viewed result type that is projected based on a view.
type Task struct {
	// Type to project
	Projected *TaskView
	// View to render
	View string
}

// TaskCollection is the viewed result type that is projected based on a view.
type TaskCollection struct {
	// Type to project
	Projected TaskCollectionView
	// View to render
	View string
}

// Image is the viewed result type that is projected based on a view.
type Image struct {
	// Type to project
	Projected *ImageView
	// View to render
	View string
}

// RegisterResultView is a type that runs validations on a projected type.
type RegisterResultView struct {
	// Task ID of the registration process
	TaskID *int
}

// TaskView is a type that runs validations on a projected type.
type TaskView struct {
	// JOb ID of the registration process
	ID *int
	// Status of the registration process
	Status *string
	// List of states from the very beginning of the process
	States []*TaskStateView
	// txid
	Txid   *string
	Ticket *ArtworkTicketView
}

// TaskStateView is a type that runs validations on a projected type.
type TaskStateView struct {
	// Date of the status creation
	Date *string
	// Status of the registration process
	Status *string
}

// ArtworkTicketView is a type that runs validations on a projected type.
type ArtworkTicketView struct {
	// Name of the artwork
	Name *string
	// Description of the artwork
	Description *string
	// Keywords
	Keywords *string
	// Series name
	SeriesName *string
	// Number of copies issued
	IssuedCopies *int
	// Uploaded image ID
	ImageID *int
	// Artwork creation video youtube URL
	YoutubeURL *string
	// Artist's PastelID
	ArtistPastelID *string
	// Name of the artist
	ArtistName *string
	// Artist website URL
	ArtistWebsiteURL *string
	// Spendable address
	SpendableAddress *string
	NetworkFee       *float32
}

// TaskCollectionView is a type that runs validations on a projected type.
type TaskCollectionView []*TaskView

// ImageView is a type that runs validations on a projected type.
type ImageView struct {
	// Uploaded image ID
	ImageID *int
	// Image expiration
	ExpiresIn *string
}

var (
	// RegisterResultMap is a map of attribute names in result type RegisterResult
	// indexed by view name.
	RegisterResultMap = map[string][]string{
		"default": []string{
			"task_id",
		},
	}
	// TaskMap is a map of attribute names in result type Task indexed by view name.
	TaskMap = map[string][]string{
		"tiny": []string{
			"id",
			"status",
			"txid",
		},
		"default": []string{
			"id",
			"status",
			"states",
			"txid",
			"ticket",
		},
	}
	// TaskCollectionMap is a map of attribute names in result type TaskCollection
	// indexed by view name.
	TaskCollectionMap = map[string][]string{
		"tiny": []string{
			"id",
			"status",
			"txid",
		},
		"default": []string{
			"id",
			"status",
			"states",
			"txid",
			"ticket",
		},
	}
	// ImageMap is a map of attribute names in result type Image indexed by view
	// name.
	ImageMap = map[string][]string{
		"default": []string{
			"image_id",
			"expires_in",
		},
	}
)

// ValidateRegisterResult runs the validations defined on the viewed result
// type RegisterResult.
func ValidateRegisterResult(result *RegisterResult) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateRegisterResultView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateTask runs the validations defined on the viewed result type Task.
func ValidateTask(result *Task) (err error) {
	switch result.View {
	case "tiny":
		err = ValidateTaskViewTiny(result.Projected)
	case "default", "":
		err = ValidateTaskView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"tiny", "default"})
	}
	return
}

// ValidateTaskCollection runs the validations defined on the viewed result
// type TaskCollection.
func ValidateTaskCollection(result TaskCollection) (err error) {
	switch result.View {
	case "tiny":
		err = ValidateTaskCollectionViewTiny(result.Projected)
	case "default", "":
		err = ValidateTaskCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"tiny", "default"})
	}
	return
}

// ValidateImage runs the validations defined on the viewed result type Image.
func ValidateImage(result *Image) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateImageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateRegisterResultView runs the validations defined on
// RegisterResultView using the "default" view.
func ValidateRegisterResultView(result *RegisterResultView) (err error) {
	if result.TaskID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("task_id", "result"))
	}
	return
}

// ValidateTaskViewTiny runs the validations defined on TaskView using the
// "tiny" view.
func ValidateTaskViewTiny(result *TaskView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "Registration Started" || *result.Status == "Artwork Accepted" || *result.Status == "Waiting Activation" || *result.Status == "Activated" || *result.Status == "Error" || *result.Status == "Finish") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"Registration Started", "Artwork Accepted", "Waiting Activation", "Activated", "Error", "Finish"}))
		}
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) < 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, true))
		}
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) > 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, false))
		}
	}
	return
}

// ValidateTaskView runs the validations defined on TaskView using the
// "default" view.
func ValidateTaskView(result *TaskView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Ticket == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ticket", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "Registration Started" || *result.Status == "Artwork Accepted" || *result.Status == "Waiting Activation" || *result.Status == "Activated" || *result.Status == "Error" || *result.Status == "Finish") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"Registration Started", "Artwork Accepted", "Waiting Activation", "Activated", "Error", "Finish"}))
		}
	}
	for _, e := range result.States {
		if e != nil {
			if err2 := ValidateTaskStateView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) < 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, true))
		}
	}
	if result.Txid != nil {
		if utf8.RuneCountInString(*result.Txid) > 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.txid", *result.Txid, utf8.RuneCountInString(*result.Txid), 64, false))
		}
	}
	if result.Ticket != nil {
		if err2 := ValidateArtworkTicketView(result.Ticket); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTaskStateView runs the validations defined on TaskStateView.
func ValidateTaskStateView(result *TaskStateView) (err error) {
	if result.Date == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "Registration Started" || *result.Status == "Artwork Accepted" || *result.Status == "Waiting Activation" || *result.Status == "Activated" || *result.Status == "Error" || *result.Status == "Finish") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"Registration Started", "Artwork Accepted", "Waiting Activation", "Activated", "Error", "Finish"}))
		}
	}
	return
}

// ValidateArtworkTicketView runs the validations defined on ArtworkTicketView.
func ValidateArtworkTicketView(result *ArtworkTicketView) (err error) {
	if result.ArtistName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artist_name", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.IssuedCopies == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("issued_copies", "result"))
	}
	if result.ImageID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_id", "result"))
	}
	if result.ArtistPastelID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artist_pastelid", "result"))
	}
	if result.SpendableAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("spendable_address", "result"))
	}
	if result.NetworkFee == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("network_fee", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 256, false))
		}
	}
	if result.Description != nil {
		if utf8.RuneCountInString(*result.Description) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.description", *result.Description, utf8.RuneCountInString(*result.Description), 1024, false))
		}
	}
	if result.Keywords != nil {
		if utf8.RuneCountInString(*result.Keywords) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.keywords", *result.Keywords, utf8.RuneCountInString(*result.Keywords), 256, false))
		}
	}
	if result.SeriesName != nil {
		if utf8.RuneCountInString(*result.SeriesName) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.series_name", *result.SeriesName, utf8.RuneCountInString(*result.SeriesName), 256, false))
		}
	}
	if result.IssuedCopies != nil {
		if *result.IssuedCopies < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.issued_copies", *result.IssuedCopies, 1, true))
		}
	}
	if result.IssuedCopies != nil {
		if *result.IssuedCopies > 1000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.issued_copies", *result.IssuedCopies, 1000, false))
		}
	}
	if result.ImageID != nil {
		if *result.ImageID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.image_id", *result.ImageID, 1, true))
		}
	}
	if result.YoutubeURL != nil {
		if utf8.RuneCountInString(*result.YoutubeURL) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.youtube_url", *result.YoutubeURL, utf8.RuneCountInString(*result.YoutubeURL), 128, false))
		}
	}
	if result.ArtistPastelID != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.artist_pastelid", *result.ArtistPastelID, "^[a-zA-Z0-9]+$"))
	}
	if result.ArtistPastelID != nil {
		if utf8.RuneCountInString(*result.ArtistPastelID) < 86 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.artist_pastelid", *result.ArtistPastelID, utf8.RuneCountInString(*result.ArtistPastelID), 86, true))
		}
	}
	if result.ArtistPastelID != nil {
		if utf8.RuneCountInString(*result.ArtistPastelID) > 86 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.artist_pastelid", *result.ArtistPastelID, utf8.RuneCountInString(*result.ArtistPastelID), 86, false))
		}
	}
	if result.ArtistName != nil {
		if utf8.RuneCountInString(*result.ArtistName) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.artist_name", *result.ArtistName, utf8.RuneCountInString(*result.ArtistName), 256, false))
		}
	}
	if result.ArtistWebsiteURL != nil {
		if utf8.RuneCountInString(*result.ArtistWebsiteURL) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.artist_website_url", *result.ArtistWebsiteURL, utf8.RuneCountInString(*result.ArtistWebsiteURL), 256, false))
		}
	}
	if result.SpendableAddress != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.spendable_address", *result.SpendableAddress, "^[a-zA-Z0-9]+$"))
	}
	if result.SpendableAddress != nil {
		if utf8.RuneCountInString(*result.SpendableAddress) < 35 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.spendable_address", *result.SpendableAddress, utf8.RuneCountInString(*result.SpendableAddress), 35, true))
		}
	}
	if result.SpendableAddress != nil {
		if utf8.RuneCountInString(*result.SpendableAddress) > 35 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.spendable_address", *result.SpendableAddress, utf8.RuneCountInString(*result.SpendableAddress), 35, false))
		}
	}
	if result.NetworkFee != nil {
		if *result.NetworkFee < 1e-05 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("result.network_fee", *result.NetworkFee, 1e-05, true))
		}
	}
	return
}

// ValidateTaskCollectionViewTiny runs the validations defined on
// TaskCollectionView using the "tiny" view.
func ValidateTaskCollectionViewTiny(result TaskCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateTaskViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTaskCollectionView runs the validations defined on
// TaskCollectionView using the "default" view.
func ValidateTaskCollectionView(result TaskCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateTaskView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateImageView runs the validations defined on ImageView using the
// "default" view.
func ValidateImageView(result *ImageView) (err error) {
	if result.ImageID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_id", "result"))
	}
	if result.ExpiresIn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expires_in", "result"))
	}
	if result.ExpiresIn != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.expires_in", *result.ExpiresIn, goa.FormatDateTime))
	}
	return
}
