---
title: Introduction
---

The Office for National Statistics API makes datasets and other data available programmatically using HTTP. It allows you to filter datasets and directly access specific data points.

<p style="font-size: 21px">The API is available at <code>https://api.beta.ons.gov.uk/v1</code></p>

The API is open and unrestricted - no API keys are required, so you can start using it immediately.

<div class="markdown-warning-box--container margin-bottom-sm--3 margin-bottom-md--4">
<span aria-label="Warning text" class="markdown-warning-box--icon">!</span>
<div class="markdown-warning-box--text padding-left--4">
<p>This API is currently in Beta and still being developed. Please be aware that as a result of this there may occasionally be breaking changes as we enhance functionality and respond to feedback.</p>
</div>
</div>

<div>
    <a class = "btn btn--primary btn--bold margin-bottom-md--2" href="tour/getting-started/">Take a tour of the API</a>
</div>

## Getting data from ONS Beta API

### Requesting data

Observation level data is provided in [JSON](https://www.w3schools.com/js/js_json_intro.asp) through GET requests to the ['Explore our data'](dataset/) service. This allows querying of a single observation/value by providing one option per dimension, but will also allow one of these dimensions to be a 'wildcard' and return all values for this dimension.

[Guide to requesting specific observation](observations/)

### Filtering a dataset

For more complicated queries the ['Filter a dataset'](filter/) service will be the best approach. This allows any combination of dimensions to be requested using POST and will return the data as CSV and XLSX.

[Guide to filtering a dataset](filters/)

### Rate limiting rules

Rate limiting has been implemented to contain the number of requests a client can make to our API. This is to protect the service against excessive use and denial of service attacks. 

The following rate limits have been implemented on the API: 

* 120 requests per 10 seconds

* 200 requests per 1 minute

They have been applied to following domains: 

* api.ons.gov.uk

* api.beta.ons.gov.uk

If requests are made that reach these limits the IP address of the client will be blocked for 1 minute. The client will receive a JSON response with a 429 error code, the Retry-After header shows the seconds remaining until a retry can be made.  


## Terminology

Descriptions for some terms used within this API.

### Dataset

A grouping of data (editions) with shared dimensions, for example Sex, Age and Geography, and all published history of this group of data. The options in these dimensions can change over time leading to separate editions.
For example: Population Estimates for UK, England and Wales, Scotland and Northern Ireland

### Edition

A subset of the dataset representing a specific time period. For some datasets this edition can contain all time periods (all historical data). The latest version of this is displayed by default.
For example: Population Estimates for UK, England and Wales, Scotland and Northern Ireland Mid-2016

### Version

A specific instance of the edition at a point in time. New versions can be published as a result of corrections, revisions or new data becoming available.

Are these definitions helpful? Are there other terms we need to explain? [Let us know](mailto:customise.my.data@ons.gov.uk)
