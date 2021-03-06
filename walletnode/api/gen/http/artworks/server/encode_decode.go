// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"unicode/utf8"

	artworks "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks"
	artworksviews "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRegisterResponse returns an encoder for responses returned by the
// artworks register endpoint.
func EncodeRegisterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.RegisterResult)
		enc := encoder(ctx, w)
		body := NewRegisterResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeRegisterRequest returns a decoder for requests sent to the artworks
// register endpoint.
func DecodeRegisterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RegisterRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRegisterRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRegisterPayload(&body)

		return payload, nil
	}
}

// EncodeRegisterError returns an encoder for errors returned by the register
// artworks endpoint.
func EncodeRegisterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeRegisterTaskStateRequest returns a decoder for requests sent to the
// artworks registerTaskState endpoint.
func DecodeRegisterTaskStateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			taskID string
			err    error

			params = mux.Vars(r)
		)
		taskID = params["taskId"]
		if utf8.RuneCountInString(taskID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, true))
		}
		if utf8.RuneCountInString(taskID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRegisterTaskStatePayload(taskID)

		return payload, nil
	}
}

// EncodeRegisterTaskStateError returns an encoder for errors returned by the
// registerTaskState artworks endpoint.
func EncodeRegisterTaskStateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskStateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskStateInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRegisterTaskResponse returns an encoder for responses returned by the
// artworks registerTask endpoint.
func EncodeRegisterTaskResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.Task)
		enc := encoder(ctx, w)
		body := NewRegisterTaskResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRegisterTaskRequest returns a decoder for requests sent to the
// artworks registerTask endpoint.
func DecodeRegisterTaskRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			taskID string
			err    error

			params = mux.Vars(r)
		)
		taskID = params["taskId"]
		if utf8.RuneCountInString(taskID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, true))
		}
		if utf8.RuneCountInString(taskID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRegisterTaskPayload(taskID)

		return payload, nil
	}
}

// EncodeRegisterTaskError returns an encoder for errors returned by the
// registerTask artworks endpoint.
func EncodeRegisterTaskError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRegisterTasksResponse returns an encoder for responses returned by the
// artworks registerTasks endpoint.
func EncodeRegisterTasksResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(artworksviews.TaskCollection)
		enc := encoder(ctx, w)
		body := NewTaskResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeRegisterTasksError returns an encoder for errors returned by the
// registerTasks artworks endpoint.
func EncodeRegisterTasksError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTasksInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUploadImageResponse returns an encoder for responses returned by the
// artworks uploadImage endpoint.
func EncodeUploadImageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.Image)
		enc := encoder(ctx, w)
		body := NewUploadImageResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeUploadImageRequest returns a decoder for requests sent to the artworks
// uploadImage endpoint.
func DecodeUploadImageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload *artworks.UploadImagePayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewArtworksUploadImageDecoder returns a decoder to decode the multipart
// request for the "artworks" service "uploadImage" endpoint.
func NewArtworksUploadImageDecoder(mux goahttp.Muxer, artworksUploadImageDecoderFn ArtworksUploadImageDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**artworks.UploadImagePayload)
			if err := artworksUploadImageDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// EncodeUploadImageError returns an encoder for errors returned by the
// uploadImage artworks endpoint.
func EncodeUploadImageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadImageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadImageInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalArtworksviewsTaskStateViewToTaskStateResponseBody builds a value of
// type *TaskStateResponseBody from a value of type
// *artworksviews.TaskStateView.
func marshalArtworksviewsTaskStateViewToTaskStateResponseBody(v *artworksviews.TaskStateView) *TaskStateResponseBody {
	if v == nil {
		return nil
	}
	res := &TaskStateResponseBody{
		Date:   *v.Date,
		Status: *v.Status,
	}

	return res
}

// marshalArtworksviewsArtworkTicketViewToArtworkTicketResponseBody builds a
// value of type *ArtworkTicketResponseBody from a value of type
// *artworksviews.ArtworkTicketView.
func marshalArtworksviewsArtworkTicketViewToArtworkTicketResponseBody(v *artworksviews.ArtworkTicketView) *ArtworkTicketResponseBody {
	res := &ArtworkTicketResponseBody{
		Name:                     *v.Name,
		Description:              v.Description,
		Keywords:                 v.Keywords,
		SeriesName:               v.SeriesName,
		IssuedCopies:             *v.IssuedCopies,
		YoutubeURL:               v.YoutubeURL,
		ArtistPastelID:           *v.ArtistPastelID,
		ArtistPastelIDPassphrase: *v.ArtistPastelIDPassphrase,
		ArtistName:               *v.ArtistName,
		ArtistWebsiteURL:         v.ArtistWebsiteURL,
		SpendableAddress:         *v.SpendableAddress,
		MaximumFee:               *v.MaximumFee,
	}

	return res
}

// marshalArtworksviewsTaskViewToTaskResponseTiny builds a value of type
// *TaskResponseTiny from a value of type *artworksviews.TaskView.
func marshalArtworksviewsTaskViewToTaskResponseTiny(v *artworksviews.TaskView) *TaskResponseTiny {
	res := &TaskResponseTiny{
		ID:     *v.ID,
		Status: *v.Status,
		Txid:   v.Txid,
	}
	if v.Ticket != nil {
		res.Ticket = marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse(v.Ticket)
	}

	return res
}

// marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse builds a value
// of type *ArtworkTicketResponse from a value of type
// *artworksviews.ArtworkTicketView.
func marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse(v *artworksviews.ArtworkTicketView) *ArtworkTicketResponse {
	res := &ArtworkTicketResponse{
		Name:                     *v.Name,
		Description:              v.Description,
		Keywords:                 v.Keywords,
		SeriesName:               v.SeriesName,
		IssuedCopies:             *v.IssuedCopies,
		YoutubeURL:               v.YoutubeURL,
		ArtistPastelID:           *v.ArtistPastelID,
		ArtistPastelIDPassphrase: *v.ArtistPastelIDPassphrase,
		ArtistName:               *v.ArtistName,
		ArtistWebsiteURL:         v.ArtistWebsiteURL,
		SpendableAddress:         *v.SpendableAddress,
		MaximumFee:               *v.MaximumFee,
	}

	return res
}
