# windows-jira-notifications

Jira notifications will periodically query Jira and if any Jira issues are found matching the JQL a Windows notification/toast will be created.
Example usage includes monitoring queues for high-priority tickets that are not yet assigned.
 
# Configuration Instructions

The application requires a `config.json` file to run. An example file `config.json.example` is provided in the repository.

```json
{
    "jira_url": "https://your-jira-instance.atlassian.net",
    "token": "your-token",
    "jql": "issuetype = Incident AND status = New AND assignee in (EMPTY)",
    "interval": 5
}
```

- jira_url: This should be the URL of your Jira instance. For example, "https://your-jira-instance.atlassian.net".

- token: This should be your Jira API token. You can generate a new API token from your Atlassian account settings. More info: https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html

- jql: This is the Jira Query Language (JQL) query that the application will use to search for issues. The default query searches for new, unassigned incidents. You can modify this query to fit your needs.

- interval: This is the interval, in minutes, at which the application will check for new issues. The default is 5 minutes.

After filling out the config.json file, save it and run the application.
