---
title: Guide to rate limiting and bot development
---

When creating a bot, scraper or automation to programmatically get data from the ONS site and APIs, please read the following guidance.

## Rate limiting

The ONS dissemination sites and API apply rate limiting to ensure a high-quality service is delivered to all users.

The following rate limits have been applied:

* 120 requests per 10 seconds - all site and API assets

* 200 requests per 1 minute - all site and API assets

* 15 requests per 10 seconds - high demand site assets

The time for which you will be blocked varies and is dynamically changed by the ONS, depending on release schedules and observed demand.

If you exceed these limits, the API will return a `429 Too Many Requests` HTTP status code and a `Retry-After` header containing the number of seconds until you may try your request again.

The ONS expects that you programmatically respect the Retry-After header and back-off your polling accordingly. If this is not respected, our algorithms will automatically impose a total block to our services for anything up to 24 hours.

Read more in our [fair usage policy](TODO)

## Constructing a good user-agent for bot usage

The `user-agent` header informs us of the characteristics of the requesting user agent.

### Naming your bot and versioning

Naming your bot and providing a version makes it easier for us to identify and work with you should there be any problems.

```shell
botName/Version1.0.0
```

### If you are writing a bot include a URL or contact information

When writing bots it is best practice to include a means of contacting the bot owner if necessary. The best method is to include a URL to a page detailing the owner of the bot, the purpose of the bot and providing a means of contacting you. Alternatively, you can include a company email address (do not use a personal email) we can use to contact you for further information.

Examples:

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

### Do not use personally identifying information in the `user-agent` header

Do not include your name, username or personal email in the `user-agent`. Your company name is acceptable if you have one.

Anything you provide will be stored in our logs as per our [data retention policy](https://www.ons.gov.uk/aboutus/transparencyandgovernance/dataprotection/dataprotectionpolicy)

### Do not use HTML tags

When providing a URL as your bot information please refrain from using the HTML tags as follows:

```shell
botName/Version1.0.0 (organisation-name <a href="http://www.organisation-site.com/bot.html>My Bot</a>)
```

Instead use:

```shell
botName/Version1.0.0 (organisation-name +http://www.organisation-site.com/bot.html)
```

### Keep it simple; do not include UTF8 or emojis

It is best practice to keep your `user-agent` as simple as possible - only use letters, numbers and basic symbols (brackets, hyphens and slashes).
