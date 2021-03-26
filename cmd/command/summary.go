package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// timeCmd represents the time command
var summaryTimeCmd = &cobra.Command{
	Use:   "summary sourcepath dst",
	Short: "summary tool",
	Long:  `根据当前目录自动生成summary文件`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("源路径不能为空")
			return
		}
		path := args[0]
		if path == "." {
			currentPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatal(err)
			}
			path = currentPath
		}
		data := genSummary(path)
		if data ==""{
			fmt.Println("源目录下没有可用信息")
			return
		}
		var dstpath string
		if len(args) <= 1 {
			dstpath = path + "/" + "SUMMARY.md"
		} else {
			dstpath = args[1]
		}
		saveSummary(dstpath, data)
	},
}

func init() {
	rootCmd.AddCommand(summaryTimeCmd)

}

func genSummary(dirpath string) string {
	if !strings.HasSuffix(dirpath, "/") {
		dirpath += "/"
	}
	allDir := make(map[string]bool)
	var result string
	filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}

		allPath := path[len(dirpath):]
		paths := strings.Split(allPath, "/")
		if len(paths) <= 1 {
			return nil
		}

		for i := 0; i < len(paths); i++ {
			if !strings.HasSuffix(paths[i], ".md") {
				if !allDir[paths[i]] {
					allDir[paths[i]] = true
					var kongge string
					var jing string = "##"
					for j := 0; j < i; j++ {
						kongge += "  "
						jing += "##"
					}

					result += fmt.Sprintf("%s* %s [%s](%s) \r\n", kongge, jing, paths[i], "./"+strings.Join(paths[0:i+1], "/"))
				} else {
					continue
				}
			} else {
				var kongge string
				for j := 0; j < i; j++ {
					kongge += "  "
				}
				//	文件暂时不展示
				//result += fmt.Sprintf("%s* [%s](%s) \r\n", kongge, strings.Split(paths[i], ".")[0], "./"+strings.Join(paths[0:i+1], "/"))
			}

		}

		return nil
	})
	return result
}

func saveSummary(dstpath, data string) {
	file, err := os.Create(dstpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
