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

Your tooling should respect the `Retry-After` header and pause your requests for the specified duration. If this is not respected our algorithms may impose a block to our services for up to 1 hour.

We reserve the right to change our rate limits at any time and without prior notice. Read more in our [fair use policy](https://www.ons.gov.uk/help/fairusepolicy)

## Constructing a good User-Agent for bot usage

The `User-Agent` header informs us of the characteristics of the requesting user agent.

### Naming your bot and versioning

Naming your bot and providing a version makes it easier for us to identify and work with you should there be any problems.

```text
botName/Version1.0.0
```

### Include a URL or contact information in your user-agent

When writing bots, it is best practice to include a means of contacting the bot owner if necessary. The best method is to include a URL to a page detailing the owner of the bot, the purpose of the bot and providing a means of contacting you. Alternatively, you can include a company email address (do not use a personal email) we can use to contact you for further information.

Examples:

```text
botName/Version1.0.0 (organisation-name contact@organisation.com +http://www.organisation-site.com/bot.html)
botName/Version1.0.0 (organisation-name +http://www.organisation-site.com/bot.html)
botName/Version1.0.0 (organisation-name contact@organisation.com)
```

### Do not use personally identifying information in the `User-Agent` header

Do not include your name, username or personal email in the `User-Agent`. Your company name is acceptable if you have one.

Anything you provide will be stored in our logs as per our [terms and conditions](https://www.ons.gov.uk/help/termsandconditions).

### Do not use HTML tags

When providing a URL as your bot information please refrain from using the HTML tags as follows:

```text
botName/Version1.0.0 (organisation-name <a href="http://www.organisation-site.com/bot.html">My Bot</a>)
```

Instead use:

```text
botName/Version1.0.0 (organisation-name +http://www.organisation-site.com/bot.html)
```

### Keep it simple; do not include UTF8 or emojis

It is best practice to keep your `User-Agent` as simple as possible - only use letters, numbers and basic symbols (brackets, hyphens and slashes).
