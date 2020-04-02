// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"github.com/fanux/sealos/cert"
	"github.com/spf13/cobra"
	"github.com/wonderivan/logger"
	"net"
)

type Flag struct {
	DNS []string
	IP []string
	NodeName string
	NodeIP string
}

var config *Flag

// certCmd represents the cert command
var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "generate certs",
	Long: `you can specify expire time`,
	Run: func(cmd *cobra.Command, args []string) {
		var ips []net.IP
		for _,ip := range config.IP {
			netip:=net.ParseIP(ip).To4()
			if netip == nil {
				logger.Warn("invalid altname : %s",ip)
				continue
			}
			ips = append(ips, netip)
		}
		certConfig := &cert.SealosCertMetaData{
			APIServer: cert.AltNames{
				DNSNames: config.DNS,
				IPs: ips,
			},
			NodeName:  config.NodeName,
			NodeIP:    config.NodeIP,
		}
		cert.GenerateAll(certConfig)
	},
}

func init() {
	config = &Flag{}
	rootCmd.AddCommand(certCmd)

	cleanCmd.Flags().StringSliceVar(&config.DNS, "alt-names", []string{}, "like sealyun.com or 10.103.97.2")
	cleanCmd.Flags().StringVar(&config.NodeName, "node-name", "", "like master0")
	cleanCmd.Flags().StringVar(&config.NodeIP, "serviceCIRD", "", "like 10.103.97.2/24")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// certCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// certCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
