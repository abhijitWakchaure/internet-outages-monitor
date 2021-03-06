# Internet Outages Monitor

Simple internet outages monitor in golang with slack notification on reconnection

## Goal

Run the program in background (daemon) and monitor the internet connection status after set interval.
In the event of disconnection the program should notify using a notifier (for now it supports slack only, PR's welcome)

## Prerequisites

This program depends on the nc (or netcat) utility. Most of the OS today have this utility pre-installed, but just to be sure you can check this by running following command:

```bash
nc -dzw1 google.com 443
```

## Installation

### Step 1: Create Slack App

1. Go to [https://api.slack.com/apps](https://api.slack.com/apps) and `Create New App` using `From Scratch`

1. Under `Add features and functionality` choose `Incoming Webhooks` and turn on the `Activate Incoming Webhooks` option

1. Create a new webhook URL for your app by clicking the button `Add New Webhook to Workspace`

1. Grant the permission to post in any channel of your choosing and copy the webhook URL. It should look like this:

    ```text
    https://hooks.slack.com/services/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
    ```

1. Export the webhook URL as `SLACK_WEBHOOK_URL` using command:

    ```bash
    export SLACK_WEBHOOK_URL=<your slack webhook URL>
    ```

1. To test if your slack app is working correctly you can use the following command to send `Hello, World!` message to your channel:

    ```bash
    curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' <your slack webhook URL>
    ```

    You should see the `ok` as output and your slack channel should receive the `Hello, World!` message.

### Step 2: Configure the environment variables

This program uses the following env variables. Set them according to your needs.

| Variable Name | Required | Default Value | Summary |
|--|--|--|--
| SLACK_WEBHOOK_URL | Yes | N/A | The webhook URL for your slack app from `Step 1`
| NC_DOMAIN | No | `google.com` | Domain you want to use to check internet connection
| NC_PORT | No | `443` | Port you want to use for netcat
| SLACK_NOTIFY_ON_REGISTER | No | `true` | If true you will get a message `Slack Notifier registered` everytime the program starts
| TICK_INTERVAL | No | `30s` | Interval between two internet checks. e.g. `15s`, `1m` etc

If you prefer using `.env` files instead of exporting individual variable, here is a sample `.env` file:

```bash
TICK_INTERVAL=30s
NC_DOMAIN=google.com
NC_PORT=443
SLACK_NOTIFY_ON_REGISTER=true
SLACK_WEBHOOK_URL=<your slack webhook URL>
```

To use the `.env` file don't forget to source it using command:

```bash
source .env
```

### Step 3: Run the program

You can run the program by downloading the binary from [latest release](https://github.com/abhijitWakchaure/internet-outages-monitor/releases/latest) or if you have [Go (golang)](https://go.dev/) installed you can directly [download](https://github.com/abhijitWakchaure/internet-outages-monitor/archive/refs/heads/master.zip) the source code and run using following command:

```bash
go run .
```
