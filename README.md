# Hello Slack

A brief demo in golang/GAE/Slack.

The key `application` in app.yaml needs to be modified and then this app will
deploy to *YOUR* Google cloud account and allow you to respond with a hello
world style response to anyone saying "hello slack".

In case it wasn't obvious, this is meant to be expanded upon.

Add this (and maybe its offspring) to your Slack integrations by going to Slack
-> Integrations -> All Services -> Outgoing Webhooks and clicking "Add Outgoing
WebHooks Integrations".  Trigger words should be "hello slack" and the URL
should be the base URL for your GAE application.  Mine is
`https://strudelline-helloslack.appspot.com/`.

The token doesn't matter though you *will* want it to matter if you decide to
store proprietary data in whatever you make out of this -- the token is your
method of verifying that the Slack account contacting your server is *your*
Slack account.

Note also:

It appears that the number of webhooks and slash commands doesn't matter for
your 5 free integrations -- they count as 1, no matter how many you've
configured.
