// Converter uses UNIX ffmpeg library to convert .amr files to .mp3 and vice versa.
package converter

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

const (
	AudioSamplingRateMP3  = "22050"
	AudioBitRate          = "12.2k" // in Hz
	NumberOfAudioChannels = "1"
	AudioSamplingRateAMR  = "8000"
)

// ConvertToMP3 converts given .amr file into .mp3 file
// Example: ConvertToMP3("oldAudio", "newAudio")
// Result: newAudio.mp3
// Notice, that filename is passed without extension!
// If only one argument is passed, than result will have the same name
func ConvertToMP3(filenames ...string) error {
	var toFilename string
	var fromFilename string = filenames[0]
	switch len(filenames) {
	case 1:
		toFilename = filenames[0]
		break
	case 0:
		return errors.New("error: no arguements are passed")
	default:
		toFilename = filenames[1]
	}
	// Convert to MP3
	comm := exec.Command("ffmpeg", "-i", fromFilename+".amr", "-ar", AudioSamplingRateMP3, toFilename+".mp3")
	if err := comm.Run(); err != nil {
		return err
	}
	return nil
}

// ConvertToAMR converts given .mp3 file into .amr file
// Example: ConvertToMP3("oldAudio", "newAudio")
// Result: newAudio.amr
// Notice, that filename is passed without extension!
// If only one argument is passed, than result will have the same name
func ConvertToAMR(filenames ...string) error {
	var toFilename string
	var fromFilename string = filenames[0]
	switch len(filenames) {
	case 1:
		toFilename = filenames[0]
		break
	case 0:
		return errors.New("error: no arguements are passed")
	default:
		toFilename = filenames[1]
	}
	// Convert to WAV
	comm := exec.Command("ffmpeg", "-i", fromFilename+".mp3", "-f", "wav", "./"+toFilename+".wav")
	if err := comm.Run(); err != nil {
		return err
	}
	// Convert to AMR
	comm = exec.Command("ffmpeg", "-i", "./"+toFilename+".wav", "-ab", AudioBitRate, "-ac", NumberOfAudioChannels, "-ar", AudioSamplingRateAMR, toFilename+".amr")
	if err := comm.Run(); err != nil {
		return err
	}
	// Delete tmp file
	if err := os.Remove(toFilename + ".wav"); err != nil {
		log.Println(err)
	}
	return nil
}
