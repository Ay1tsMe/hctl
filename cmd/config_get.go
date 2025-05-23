// Copyright 2024 Fabian `xx4h` Sylvester
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io"
	"slices"

	"github.com/spf13/cobra"

	"github.com/xx4h/hctl/pkg"
	o "github.com/xx4h/hctl/pkg/output"
)

const (
	// editorconfig-checker-disable
	configGetExample = `
  # Get all config options
  hctl config get

  # Get a specific config option
  hctl config get hub.url
  `
	// editorconfig-checker-enable
)

func newConfigGetCmd(h *pkg.Hctl, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get [PATH]",
		Short:   "Get configuration parameters",
		Aliases: []string{"g", "ge"},
		Example: configGetExample,
		ValidArgsFunction: func(_ *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return noMoreArgsComp()
			}
			return compListConfig(toComplete, args, h)
		},
		Run: func(_ *cobra.Command, args []string) {
			var header = append([]any{}, "OPTION", "VALUE")
			var clist [][]interface{}
			if len(args) == 0 {
				a, _ := compListConfig("", []string{}, h)
				slices.Sort(a)
				for _, b := range a {
					v, err := h.GetConfigValue(b)
					l := []any{}
					if err == nil {
						l = append(l, b, v)
					}
					clist = append(clist, l)
				}
			} else {
				v, err := h.GetConfigValue(args[0])
				if err != nil {
					o.FprintError(out, err)
				}
				l := append([]any{}, args[0], v)
				clist = append(clist, l)
			}
			o.FprintSuccessListWithHeader(out, header, clist)
		},
	}

	return cmd
}
