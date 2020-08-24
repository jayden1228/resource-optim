package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"resource-optim/internal/pkg/env"
	"resource-optim/internal/pkg/logger"
	"resource-optim/internal/pkg/optim"
	"resource-optim/internal/pkg/path"
	"strings"

	"resource-optim/config"

	"github.com/gogf/gf/os/gfile"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/xfrr/goffmpeg/transcoder"
)

var rootCmd = &cobra.Command{
	Use:   "resource-optim",
	Short: "resource-optim is a command line tool for resource-optim sdk",
	Long:  "resource-optim is a command line tool for resource-optim sdk",
	Run:   OptimCmd,
}

func OptimCmd(_ *cobra.Command, args []string) {
	if err := env.CheckToolRequired(); err != nil {
		logger.LogE(err)
	}
	inputRootPath, err := prompt("input dir", "input dir cannot be empty")
	if err != nil {
		logger.LogE(err)
		return
	}

	outputRootPath, err := promptDefault("output dir", "output dir cannot be empty", defaultOutputDir)
	if err != nil {
		logger.LogE(err)
		return
	}

	optimType, err := promptOptimType(optim.GetOptimTypes())

	trans := new(transcoder.Transcoder)

	inputRootPath = path.HandleHomedirOrPwd(inputRootPath)
	if !gfile.Exists(inputRootPath) {
		logger.LogE("input dir not exist")
		return
	}

	outputRootPath = path.HandleHomedirOrPwd(outputRootPath)
	if outputRootPath == emptyPath {
		outputRootPath = gfile.Join(gfile.Dir(inputRootPath), defaultOutputDir)
	}

	err = BatchOptim(trans, inputRootPath, outputRootPath, optimType)
	if err != nil {
		logger.LogE(err.Error())
		return
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}

const (
	emptyPath        = ""
	defaultOutputDir = "output"
)

func BatchOptim(trans *transcoder.Transcoder, inputRootPath, outputRootPath, optimType string) error {
	err := filepath.Walk(inputRootPath, func(fPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", fPath, err)
			return err
		}
		// path是目录不处理
		if gfile.IsDir(fPath) {
			return nil
		}
		// 计算path
		outPath := strings.Replace(fPath, inputRootPath, outputRootPath, 1)
		outDir := gfile.Dir(outPath)
		if !gfile.Exists(outDir) {
			if err := gfile.Mkdir(outDir); err != nil {
				return err
			}
		}

		// 根据选定文件方式处理
		err = optim.TypeOptimMap[optimType](trans, fPath, outPath)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

func prompt(label string, alert string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: emptyValidate(alert),
	}
	return prompt.Run()
}

func promptDefault(label string, alert string, defaultValue string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: emptyValidate(alert),
		Default:  defaultValue,
	}
	return prompt.Run()
}

func promptOptimType(template []string) (string, error) {
	prompt := promptui.Select{
		Label: "select optim file type",
		Items: template,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return result, err
	}
	return result, nil
}

func emptyValidate(alert string) promptui.ValidateFunc {
	return func(input string) error {
		if len(input) < 1 {
			return errors.New(alert)
		}
		return nil
	}
}
