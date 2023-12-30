package core

import (
	"crypto/sha1"
	"encoding/hex"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Sam12121/YaraHunter/constants"
	log "github.com/sirupsen/logrus"
)

// CreateRecursiveDir Create directory structure recursively, if they do not exist
// @parameters
// completePath - Complete path of directory which needs to be created
// @returns
// Error - Errors if any. Otherwise, returns nil
func CreateRecursiveDir(completePath string) error {
	if _, err := os.Stat(completePath); os.IsNotExist(err) {
		log.Debugf("Folder does not exist. Creating folder... %s", completePath)
		err = os.MkdirAll(completePath, os.ModePerm)
		if err != nil {
			log.Errorf("createRecursiveDir %q: %s", completePath, err)
		}
		return err
	} else if err != nil {
		log.Errorf("createRecursiveDir %q: %s. Deleting temp dir", completePath, err)
		_ = DeleteTmpDir(completePath)
		return err
	}

	return nil
}

// Create a sanitized string from image name which can used as a filename
// @parameters
// imageName - Name of the container image
// @returns
// string - Sanitized string which can used as part of filename
func getSanitizedString(imageName string) string {
	//nolint:gocritic
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return "error"
	}
	sanitizedName := reg.ReplaceAllString(imageName, "")
	return sanitizedName
}

// GetJSONFilepath Return complete path and filename for json output file
// @parameters
// image - Name of the container image or dir, for which json filename and path will be created
// @returns
// string - Sanitized string which can used as path and filename of json output file
// Error - Errors if path can't be created. Otherwise, returns nil
func GetJSONFilepath(jsonFilename, outputPath string) (string, error) {
	if jsonFilename == "" {
		return "", nil
	}
	outputDir := outputPath
	if outputDir != "" && !PathExists(outputDir) {
		err := CreateRecursiveDir(outputDir)
		if err != nil {
			log.Errorf("GetJsonFilepath: Could not create output dir: %s", err)
			return "", err
		}
	}
	jsonFilePath := filepath.Join(outputDir, jsonFilename)
	log.Infof("Complete json file path and name: %s", jsonFilePath)
	return jsonFilePath, nil
}

// GetTmpDir Create a temporrary directory to extract the conetents of container image
// @parameters
// imageName - Name of the container image
// @returns
// String - Complete path of the based directory where image will be extracted, empty string if error
// Error - Errors if any. Otherwise, returns nil
func GetTmpDir(imageName, tempDirectory string) (string, error) {

	scanID := "df_" + getSanitizedString(imageName)

	tempPath := filepath.Join(tempDirectory, "Toae", constants.TempDirSuffix, scanID)

	// if runtime.GOOS == "windows" {
	//	tempPath = dir + "\temp\Toae\IOCScanning\df_" + scanId
	//}

	completeTempPath := path.Join(tempPath, constants.ExtractedImageFilesDir)

	err := CreateRecursiveDir(completeTempPath)
	if err != nil {
		log.Errorf("getTmpDir: Could not create temp dir%s", err)
		return "", err
	}

	return tempPath, err
}

// DeleteTmpDir Delete the temporary directory
// @parameters
// outputDir - Directory which need to be deleted
// @returns
// Error - Errors if any. Otherwise, returns nil
func DeleteTmpDir(outputDir string) error {
	log.Infof("Deleting temporary dir %s", outputDir)
	// Output dir will be empty string in case of error, don't delete
	if outputDir != "" {
		// deleteFiles(outputDir+"/", "*")
		err := os.RemoveAll(outputDir)
		if err != nil {
			log.Errorf("deleteTmpDir: Could not delete temp dir: %s", err)
			return err
		}
	}
	return nil
}

// DeleteFiles Delete all the files and dirs recursively in specified directory
// @parameters
// path - Directory whose contents need to be deleted
// wildcard - patterns to match the filenames (e.g. '*')
func DeleteFiles(path string, wildCard string) {

	var val string
	files, _ := filepath.Glob(path + wildCard)
	for _, val = range files {
		os.RemoveAll(val)
	}
}

// IsSymLink Check if input is a symLink, not normal file/dir
// path - Pathname which needs to be checked for symbolic link
// @returns
// bool - Return true if input is a symLink
func IsSymLink(path string) bool {
	// can handle symbolic link, but will no follow the link
	fileInfo, err := os.Lstat(path)

	if err != nil {
		// panic(err)
		return false
	}

	// --- check if file is a symlink
	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		return true
	}

	return false
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func LogIfError(text string, err error) {
	if err != nil {
		log.Errorf("%s (%s", text, err.Error())
	}
}

func GetHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}

	return plural
}

func GetEntropy(data string) (entropy float64) {
	if data == "" {
		return 0
	}

	for i := 0; i < 256; i++ {
		px := float64(strings.Count(data, string(byte(i)))) / float64(len(data))
		if px > 0 {
			entropy += -px * math.Log2(px)
		}
	}

	return entropy
}

func GetTimestamp() int64 {
	return time.Now().UTC().UnixNano() / 1000000
}

func GetCurrentTime() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.000") + "Z"
}
