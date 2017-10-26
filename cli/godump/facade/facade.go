package facade

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/godump"
)

const (
	//Name is applicatin name
	Name = "godump"
	//Version is version for applicatin
	Version = "v0.1.0"
)

//ExitCode is OS exit code enumeration class
type ExitCode int

const (
	//Normal is OS exit code "normal"
	Normal ExitCode = iota
	//Abnormal is OS exit code "abnormal"
	Abnormal
)

//Int convert integer value
func (c ExitCode) Int() int {
	return int(c)
}

//Stringer method
func (c ExitCode) String() string {
	switch c {
	case Normal:
		return "normal end"
	case Abnormal:
		return "abnormal end"
	default:
		return "unknown"
	}
}

var (
	reader io.Reader //input reader (maybe os.Stdin)
	result io.Reader //result string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: Name + " [flags] [binary file]",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if len(args) > 0 {
			file, err2 := os.Open(args[0]) //args[0] is maybe file path
			if err2 != nil {
				return err
			}
			defer file.Close()
			reader = file
		}
		result, err = godump.DumpBytes(reader, name)
		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cui *gocli.UI) (exit ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = Abnormal
		}
	}()

	//execution
	exit = Normal
	reader = cui.Reader() //default reader; maybe os.Stdin
	if err := rootCmd.Execute(); err != nil {
		//cui.OutputErrln(err) //no need to output error
		exit = Abnormal
		return
	}
	cui.WriteFrom(result)
	return
}

func init() {
	result = ioutil.NopCloser(bytes.NewReader(nil))
	rootCmd.Flags().StringP("name", "n", "dumpList", "value name")
}
