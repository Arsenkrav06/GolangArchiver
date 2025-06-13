package cmd

import (
	"archiver/lib/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using varible-lenght code",
	Run:   pack,
}

var ErrEmptyPath = errors.New("path to file is not specified")

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)

	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
