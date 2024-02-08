/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	kubeclient "github.com/thakurnishu/ctl-the-kube/kubeClient"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
    nameSpace string
    isWide bool
    clientSet = kubeclient.GetClient()
)

func listDeployments(ns string, ctx context.Context) error { 

    deployList, err := clientSet.AppsV1().Deployments(ns).List(ctx, metaV1.ListOptions{})
    if err != nil {
        return err
    }

    if len(deployList.Items) == 0 {
        fmt.Printf("no deployment resource found in %s namespace\n", ns)
        return nil
    }

    deployTable := tablewriter.NewWriter(os.Stdout)

    if isWide {
        deployTable.SetHeader([]string{"NAME", "REPLICAS", "IMAGE","SELECTOR"})
    } else {
        deployTable.SetHeader([]string{"NAME", "REPLICAS"})
    }
    deployTable.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	deployTable.SetCenterSeparator("")
    deployTable.SetRowLine(false)

    for _, deployment := range deployList.Items {

        var selector string
        var data []string

        noOfReplicas := strconv.Itoa(int(deployment.Status.Replicas))
        matchLabels := deployment.Spec.Selector.MatchLabels
        imageName := deployment.Spec.Template.Spec.Containers[0].Image


        for key, value := range matchLabels {
            selector = fmt.Sprintf("%s=%s", key, value)
        }

        if isWide {
            data = []string{deployment.Name, noOfReplicas, imageName, selector}
        }else {
            data = []string{deployment.Name, noOfReplicas}
        }
        

        deployTable.Append(data)
    }
    deployTable.Render()
    return nil
}

func getDeployment(ns string, name string, ctx context.Context) error {

    deployment, err := clientSet.AppsV1().Deployments(ns).Get(ctx, name, metaV1.GetOptions{})
    if err != nil {
        return err
    }
    fmt.Println(deployment.Name)

    return nil
}

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "list deployment resource",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

        ctx := context.Background()
        if len(args) == 0 {
            if nameSpace == "" {
                nameSpace = "default"
            }
            err := listDeployments(nameSpace, ctx)
            if err != nil {
                fmt.Printf("ERROR: %s\n", err.Error())
            }
        }else {
            if nameSpace == "" {
                nameSpace = "default"
            }
            name := args[0]
            err := getDeployment(nameSpace, name, ctx)
            if err != nil {
                fmt.Printf("ERROR: %s\n", err.Error())
            }
        }
	},
}

func init() {
    GetCmd.AddCommand(deploymentsCmd)

    deploymentsCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "", "from namespace")

    deploymentsCmd.Flags().BoolVarP(&isWide, "wide", "w", false, "display more information")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
