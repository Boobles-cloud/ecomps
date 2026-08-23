package httputils

import (
	"errors"
	"io"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"reflect"

	jsonutils "ecomps.boobles.cloud/backend/utils/http_utils/json_utils"
	"github.com/google/uuid"
)

// Creates a file and deserializes the metadata from the given request
// A file upload request should always include a metadata tag and the given file
// If you just want a file path back, please insert a string as generic parameter!
// If you want the metadata, set the name of your wanted strcuct field in [wantedStructField]
func GetMetadataAndFileFromFormValues[T any](r *http.Request, metaDataKey, fileKey, wantedStructField string) (T, error) {
	var s T

	if err := r.ParseMultipartForm(5 << 20); err != nil {
		return s, err
	}

	file, header, err := r.FormFile(fileKey)

	if err != nil {
		return s, err
	}

	filePath, err := saveFile(file, header)

	if err != nil {
		return s, err
	}

	valOfT := reflect.TypeOf(s)

	switch valOfT.Kind() {
	// If its just a string, we check if we can set a value to it and then just return the filepath!
	case reflect.String:

		if reflect.ValueOf(s).Elem().CanSet() {
			reflect.ValueOf(s).Elem().SetString(filePath)
			return s, nil
		}

		return s, errors.New("Failed setting file path to string!")
	// If its a struct, we send the metadata back!
	case reflect.Struct:

		// Get an deserialize the metadata
		metadata := r.FormValue(metaDataKey)
		v, err := jsonutils.JsonDeserilizeBytes[T]([]byte(metadata))

		if err != nil {
			return s, err
		}

		valOfV := reflect.ValueOf(v).Elem()

		for i := range valOfV.NumField() {

			currField := valOfV.Field(i)

			if currField.Type().Name() != wantedStructField {
				continue
			}

			if currField.IsValid() && currField.CanSet() && currField.Kind() == reflect.String {
				currField.SetString(filePath)
				return v, nil
			}
		}

		return s, errors.New("Failed to set filepath to wanted field!")
	default:
		return s, errors.New("Invalid type! Type: " + valOfT.Kind().String() + " is not supported!")
	}
}

// This func saves the file and returns the filepath or an error
func saveFile(file multipart.File, header *multipart.FileHeader) (string, error) {

	defer file.Close()

	dir := os.Getenv("product_picture_path")

	// Just a safety check if the dir exists
	// If not, we migth have a bigger problem
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		return "", err
	}

	id := uuid.NewString()

	fileName := id + "_" + header.Filename

	createdFile, err := os.Create(path.Join(dir, fileName))

	if err != nil {
		return "", err
	}

	defer createdFile.Close()

	if _, err := io.Copy(createdFile, file); err != nil {
		return "", err
	}

	return path.Join(dir, createdFile.Name()), nil
}
