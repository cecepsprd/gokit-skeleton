/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/cecepsprd/gokit-skeleton/cmd/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start gokit-skeleton",
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			// starting http and grpc server
			if err := server.RunServer(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		}()

		// Wait All services to end
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
