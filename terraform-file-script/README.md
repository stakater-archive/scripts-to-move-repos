# Create / Update repos in an Organization

Get the repos name that you want to create/update, add them in the `script.sh` file in the loop. Change the Topics & Description to your specifications, currently this is for stakater-charts repos, but you need to change it to your repos.

## Creating Repos without Branch Protection

Run the script using the `repo-before-pushing.tf` in the script for creating a new repo, it will create new repo files in your organization repo. Push the new files so that these new repos will be created without Branch Protection, so that you can move the  previous code to the new repo. After creation, Run the other golang code to copy all the commits and code and release history to the new repo.

## Updating Repo with Branch Protection

After the code has been copied, again run this script, using the `repo-after-pushing.tf` to enable branch protection on the repos. Push the changed files and run the pipeline.