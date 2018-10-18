package main

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "368d80faa6b4d52ad65a96e07649713228a723ef"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	// get all pages of results
	var allRepos []*github.Repository
	for {
		// List repos by organization name
		repos, resp, err := client.Repositories.ListByOrg(ctx, "stakater", opt)
		if err != nil {
			fmt.Print(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	var chartRepos []*github.Repository
	for _, repo := range allRepos {
		name := *repo.Name
		// Get all repos containing chart-
		if strings.Contains(name, "chart-") {
			chartRepos = append(chartRepos, repo)
		}
	}
	for _, repo := range chartRepos {
		name := *repo.Name
		cloneUrl := *repo.CloneURL
		// Replace cloneUrl 'https://' with 'git@github' so that we can commit through ssh key without prompting the username
		cloneUrl = strings.Replace(cloneUrl, "https://github.com/", "git@github.com:", -1)
		fmt.Println(name)

		// For initial testing, start with one repo, so use `if` for any one repo
		if name == "chart-prometheus" {
			newRepoName := strings.TrimPrefix(name, "chart-")
			prevURL := "stakater/" + name
			newURL := "stakater-charts/" + newRepoName
			fmt.Println("start")
			runCommandVerbose("git", "clone", cloneUrl)
			changeString := "s+" + prevURL + "+" + newURL + "+g"
			fmt.Println(changeString)
			// Replace the previous Url with New Url in all of the files in the repo
			runCommandVerbose("sed", "-i", changeString, name+"/.git/config")
			runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"/Chart.yaml")
			runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"-templates/Chart.yaml.tmpl")

			runCommandVerbose("git", "-C", name, "commit", "-am", "Change url of repository")
			runCommandVerbose("git", "-C", name, "push", "--force")

			fmt.Println("done")

		}

		// runCommandVerbose("sed", "-i", changeString, name+"/"+"nexus/Chart.yaml")
		// runCommandVerbose("sed", "-i", changeString, name+"/"+"nexus"+"-templates/Chart.yaml.tmpl")
		// runCommandVerbose("sed", "-i", changeString, name+"/"+"nexus-storage"+"/Chart.yaml")
		// runCommandVerbose("sed", "-i", changeString, name+"/"+"nexus-storage"+"-templates/Chart.yaml.tmpl")

		// runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"-db/Chart.yaml")
		// runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"-db-templates/Chart.yaml")

		// runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"-storage/Chart.yaml")
		// runCommandVerbose("sed", "-i", changeString, name+"/"+newRepoName+"-storage-templates/Chart.yaml.tmpl")

		// args := [2]string{"clone", name}
		//runCommandVerbose("git", "clone", name)
		//repoName := strings.TrimPrefix(name, "chart-")
		//fmt.Print(repoName + " ")
		//}
		// name := "chart-prometheus"
		// // args := [2]string{"clone", name}
		// runCommandVerbose("git", "clone", name)
		// fmt.Println()
	}
}

// runCommandVerbose runs the command displaying its output
func runCommandVerbose(name string, args ...string) error {
	e := exec.Command(name, args...)
	err := e.Run()
	if err != nil {
		fmt.Printf("Error: Command failed  %s %s\n", name, strings.Join(args, " "))
	}
	return err
}
