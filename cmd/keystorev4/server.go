/*
 * // Copyright © 2022-2023. duktig666 Limited.
 * // Licensed under the Apache License, Version 2.0 (the "License");
 * // you may not use this file except in compliance with the License.
 * // You may obtain a copy of the License at
 * //
 * //     http://www.apache.org/licenses/LICENSE-2.0
 * //
 * // Unless required by applicable law or agreed to in writing, software
 * // distributed under the License is distributed on an "AS IS" BASIS,
 * // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * // See the License for the specific language governing permissions and
 * // limitations under the License.
 */

// description:
// @author renshiwei
// Date: 2023/6/9

package keystorev4

import (
	"github.com/spf13/cobra"
)

var (
	KeystoreV4Cmd = &cobra.Command{
		Use:   "keystorev4",
		Short: "Ethereum consensus's keystore manager",
		Long:  `Ethereum consensus's keystore manager.`,
		PreRun: func(*cobra.Command, []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			Run(cmd)
		},
	}
)

// Run runs the command.
func Run(cmd *cobra.Command) (string, error) {

}
