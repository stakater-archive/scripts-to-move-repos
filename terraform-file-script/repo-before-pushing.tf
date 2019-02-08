module "@{MODULE-NAME}" {
  source         = "github.com/stakater/terraform-module-github.git//modules/repository?ref=1.0.9"
  name           = "@{REPO-NAME}"
  topics = ["@{MODULE-NAME}", "chart", "stakater"]
  require_status_checks = true
  enable_branch_protection = false
  enforce_admins = false
  description = "Helm chart for @{REPO-NAME}"
  team_id  = "${github_team.developers.id}"
  webhooks = [
    {
      url = "https://gitwebhookproxy.tools.stackator.com/github-webhook/",
      events = "push,pull_request"
      secret = "dummysecret"
    }
  ]
}