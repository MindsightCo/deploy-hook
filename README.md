# Mindsight Deployment Hook

Reports deployment to Mindsight backend to assist in correlating code behavior with deployed changes.

## Installation

It is strongly recommended that you download and use the latest [binary release](https://github.com/MindsightCo/deploy-hook/releases).

## Usage

This hook should be executed every time you deploy your code into an environment you are tracking with Mindsight.
Running the hook lets Mindsight know how to correlate changes in your data with versions of code that you have deployed.

It is probably most convenient to integrate this tool into your existing deployment automation.

To report a deployment to the Mindsight API, run the command as follows:

```
deploy-hook -commit abc123 -repo "https://github.com/Me/my-app.git"
```

- Replace the argument to `-commit` with the commit SHA1 you are deploying.
- Replace the repository's URL with your app's repo URL.
