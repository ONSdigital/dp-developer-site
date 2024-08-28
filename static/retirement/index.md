---
title: Guide to our approach to the retirement of API endpoints 
---

This guide documents our approach to the retirement of API endpoints and what users can expect when an endpoint is going to be retired.

## Why are endpoints retired?

Endpoints, systems or features can be phased out or retired due to:

- legislative changes
- replacement technology
- lack of users
- costs of upkeep
- other business reasons

"Sunsetting" occurs at the end of the product life cycle and is the period of time when a product is planned to be retired, but is still live and in use. A product or service is only sunset when there is a formal agreement, target date agreed for when a service will be decommissioned and these have been communicated to our users, teams and stakeholders.

## Timeline of the end of life process

1. Formal agreement is put in place internally that an API is due to reach its end of life and should be sunsetted
2. Migration guide and/or blog post detailing sunset and retirement published
3. [Sunset headers](#how-will-users-know-an-endpoint-is-being-sunsetted) are put on API endpoints that are affected
4. End of life date is reached
5. Rolling blackouts can be used to determine if any further users are impacted
6. API endpoint is retired

When an endpoint is retired, 404 error codes will be received by clients still querying the endpoint. In addition to this, during a rolling blackout, the sunset headers will still be issued in order to provide further information to users.

At each stage through this process we will review any feedback collected through our [Feedback Service](https://www.ons.gov.uk/feedback) that may impact the decision to proceed with the the next stage.

## How will users know an endpoint is being sunsetted?

In the response headers to clients, we will issue three additional headers when an endpoint is scheduled to be removed:

- Deprecation (the date the decision was taken to deprecate this endpoint)
- Link (a link where you can find more information)
- Sunset (the date of when this service will cease to return data on its endpoints and instead return blanket 404 status codes)

For example:

```txt
Deprecation: Thu, 7 Mar 2024 09:00:00 GMT
Link: https://developer.ons.gov.uk/retirement
Sunset: Wed, 11 May 2024 23:59:59 GMT
```

## What can users do when their endpoint is being sunsetted?

For any API endpoint that is being sunsetted, we will provide a migration guide to using new services if there is one available.

If this is not sufficient, users can provide feedback via our [Feedback Service](https://www.ons.gov.uk/feedback) which we will review at the end of every stage of the timeline before progressing to the next stage.

## What endpoints are currently reaching end of life?

### Data endpoints for aggregated search pages

Tt has been decided that ONS will retire the data endpoints for our aggregated search pages due to the following reasons:

- out of date technology
- built as a proof of concept
- strategic technology direction is elsewhere
- cost of upkeep is too high with similar services available

The UI for these pages has already been moved to using our Search API and it is the intention of ONS to separate data and UI endpoints by domain in future. Data in JSON format will be accessible via our API found at <https://api.beta.ons.gov.uk/v1/search>

The endpoints to be decommissioned will be:

- `/alladhocs/data`
- `/allmethodologies/data`
- `*/datalist/data`
- `*/publications/data`
- `/publishedrequests/data`
- `*/staticlist/data`
- `/timeseriestool/data`
- `*/topicspecificmethodology/data`
- `/releasecalendar/data`

Items marked with a * can appear at any url path on the ONS website, with the others only appearing at the root level.

These are for the following domains:

- `ons.gov.uk`

To migrate to one of our new services, please read our [migration guide](./aggregatedsearch/).
