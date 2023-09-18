---
title: Bots
---

Hello world 3

When creating a bot, scrapper or automation to programmatically get data from the ONS site and APIs please read the following guidance.

## Rate limiting

The ONS API applies rate limiting to ensure a high quality service is delivered to all users and to protect client applications from unexpected loops.

The following rate limits have been implemented:

* 120 requests per 10 seconds - All site and API assets

* 200 requests per 1 minute - All site and API assets

* 15 requests per 10 seconds - High demand site assets

The time for which you will be blocked varies and is dynamically changed by ONS depending on release schedules and observed demand.

If you exceed these limits the API will return a 429 Too Many Requests HTTP status code and a Retry-After header containing the number of seconds until you may try your request again.

ONS expect that you programmatically respect the Retry-After header and back off your polling accordingly. If this is not respected our algorithms will automatically impose a total block to our services for anything up to 24hrs.

## Constructing a good user-agent for Bot usage

The `user-agent` header informs us of the characteristics of the requesting user agent.

### Naming your Bot and versioning

Naming your bot and providing a version makes it easier for us to identify and work with you should there be any problems.

```shell
botName/Version1.0.0
```

### Don't leak personally identifying information in the user agent

Do not include your name, username or personal email in the user-agent. It is not necessary to provide us such information and not the intention of the `user-agent` header.

Your company name is acceptable if you have one.

Anything you provide will be stored in our logs as per our [data retention policy](https://www.ons.gov.uk/aboutus/transparencyandgovernance/dataprotection/dataprotectionpolicy)

### If you're writing a bot, make sure to include some kind of url/contact info

This can be an company email address (not a personal one) or a url to a page you host to define your bots intentions or contact information:

All the information

```shell
botName/Version1.0.0 (organisation-name contact@organisation.com +http://www.organisation-site.com/bot.html)
```

Just a URL

```shell
botName/Version1.0.0 (organisation-name +http://www.organisation-site.com/bot.html)
```

Just contact information

```shell
botName/Version1.0.0 (organisation-name contact@organisation.com)
```

### Don't use HTML tags

Using the example of providing a URL to you bot information please refrain from using the HTML tags as the following;

```shell
botName/Version1.0.0 (organisation-name <a href="http://www.organisation-site.com/bot.html>My Bot</a>)
```

Instead use;

```shell
botName/Version1.0.0 (organisation-name +http://www.organisation-site.com/bot.html)
```

### Keep it simple no UTF8 or emojis

We have no problems processing UTF8 or even emojis, but its not necessary, it's a better practice to keep your `user-agent`` as simple as possible - only use letters, numbers and basic symbols (brackets, hypens, slashes etc).
