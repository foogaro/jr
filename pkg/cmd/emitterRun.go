//Copyright © 2022 Ugo Landini <ugo.landini@gmail.com>
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ugol/jr/pkg/functions"
	"log"
	"sync"
)

var emitterRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run all or selected configured emitters",
	Long:  `Run all or selected configured emitters`,
	Run: func(cmd *cobra.Command, args []string) {

		err := viper.UnmarshalKey("emitters", &emitters)
		if err != nil {
			log.Println(err)
		}

		var wg sync.WaitGroup
		number := len(emitters)

		if len(args) == 0 {
			wg.Add(number)
			for _, v := range emitters {
				//go v.Run(GlobalCfg)
				v.Run(GlobalCfg)
			}
		} else {
			for _, v := range emitters {
				if functions.Contains(args, v.Name) {
					wg.Add(1)
					go v.Run(GlobalCfg)
					//v.Run(GlobalCfg)
				}
			}
		}
		wg.Wait()

	},
}

func init() {
	emitterCmd.AddCommand(emitterRunCmd)
}