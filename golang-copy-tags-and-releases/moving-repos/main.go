package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/stakater/moving-repos/helper"

	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("Starting server")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "99f8e420af637569c6b132ceaddf5eeb75372bf9"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fmt.Println("Initialized client successfully")

	sourceRepoName := "fabric8-pipeline-library"
	destinationRepoName := "stakater-pipeline-library"

	repo, _, err := client.Repositories.Get(ctx, "stakater", sourceRepoName)
	if err != nil {
		fmt.Println("Can not get repo: ", err)
	}

	name := *repo.Name
	cloneURL := *repo.SSHURL
	fmt.Println("Repo Name: ", name)
	fmt.Println("Repo SSH URL: ", cloneURL)

	prevURL := "stakater/" + name
	newURL := "stakater/" + destinationRepoName

	helper.RunCommandVerbose("git", "clone", cloneURL)
	changeString := "s+" + prevURL + "+" + newURL + "+g"
	fmt.Println(changeString)
	helper.RunCommandVerbose("sed", "-i", changeString, name+"/.git/config")
	helper.RunCommandVerbose("git", "-C", name, "push", "--force")

	fmt.Println("done")

	repos, _, _ := client.Repositories.ListTags(ctx, "stakater", name, &github.ListOptions{})
	res, _ := json.Marshal(repos)
	fmt.Printf("%v", string(res))

	for _, element := range repos {
		SHAValueForTag := *element.Commit.SHA
		tagName := *element.Name

		fmt.Println("Adding Tag: ", tagName, " to repo: ", destinationRepoName, " from repo: ", name)

		helper.RunCommandVerbose("git", "-C", name, "checkout", "-b", SHAValueForTag, SHAValueForTag)
		helper.RunCommandVerbose("git", "-C", name, "push", "--tags")

		fmt.Println("Pushed Tag: ", tagName, " to repo: ", destinationRepoName, " from repo: ", name)

		fmt.Println("tag name: ", tagName)
		release, _, _ := client.Repositories.GetReleaseByTag(ctx, "stakater", sourceRepoName, tagName)
		fmt.Println("Release: ", release.Name)

		repositoryRelease := &github.RepositoryRelease{
			Name:            release.Name,
			TagName:         release.TagName,
			TargetCommitish: release.TargetCommitish,
			Body:            release.Body,
			Draft:           release.Draft,
			Prerelease:      release.Prerelease,
		}

		newRelease, _, _ := client.Repositories.CreateRelease(ctx, "stakater", destinationRepoName, repositoryRelease)
		fmt.Println("New release created", newRelease.Name)
	}
}
