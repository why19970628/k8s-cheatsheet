/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

//import (
//	"errors"
//
//	"github.com/spf13/cobra"
//)
//
//var rootCmd = &cobra.Command{
//	Use:   "kinject",
//	Short: "A brief description of your application",
//	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
//	Args: func(cmd *cobra.Command, args []string) error {
//		//TODO: Environment Legality Control(Must be "key:value")
//		var err error
//		switch cmd.Flag("kind").Value.String() {
//		case "deployment", "deploy", "sts", "statefulset", "pod":
//			err = nil
//		default:
//			return errors.New("requires a valid kind parameter")
//		}
//		switch cmd.Flag("type").Value.String() {
//		case "env", "Env", "Environment", "environment":
//			err = nil
//		default:
//			return errors.New("requires a valid type parameter")
//		}
//		return err
//	},
//	Run: func(cmd *cobra.Command, args []string) {
//		//logic.Entrypoint(cmd)
//	},
//}
//
//func Execute() {
//	cobra.CheckErr(rootCmd.Execute())
//}
//
//func init() {
//	rootCmd.PersistentFlags().StringP("kubeconfig", "c", "~/.kube/config", "Kubeconfig path")
//	rootCmd.PersistentFlags().StringP("environment", "e", "", "Environment Key Value(Key:Value)")
//	rootCmd.PersistentFlags().StringP("namespace", "n", "all", "Kubernetes Namespace That Will Be Affect")
//	rootCmd.PersistentFlags().StringP("mode", "m", "add", "Mode of tool(add,delete).Add will be edit if exist.")
//
//	initDeploymentCommand()
//	rootCmd.AddCommand(deploymentCmd)
//	//deploymentCmd.Flags().StringP("namespace", "n", "all", "Kubernetes Namespace That Will Be Affect")
//	//deploymentCmd.AddCommand(environmentCmd)
//}
