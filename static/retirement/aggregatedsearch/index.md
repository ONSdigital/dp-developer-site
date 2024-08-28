---
title: Retirement of the data endpoint for aggregated search pages
---

## Background

In line with our [approach to the retirement of API endpoints](../), it has been decided that ONS will retire the data endpoints for our aggregated search pages due to the following reasons:

- out of date technology
- built as a proof of concept
- strategic technology direction is elsewhere
- cost of upkeep is too high with similar services available

The UI for these pages has already been moved to using our Search API and it is the intention of ONS to separate data and UI endpoints by domain in future. Data in JSON format will be accessible via our API found at <https://api.beta.ons.gov.uk/v1/search>. You can also browse the [full API Swagger specification](/search/search).

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

The UI for each page (without `/data`) will continue to remain live where it is still in use.

## Migration guide

Below you can find detailed ways that you can still get the same data from the ONS using our latest services which will benefit from continued enhancement and investment.

### /alladhocs/data

| Version | URI                                                               |
|---------|-------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/alladhocs/data`                           |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=static_adhoc` |

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### /allmethodologies/data

| Version | URI                                                                                                            |
|---------|----------------------------------------------------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/allmethodologies/data`                                                                 |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=static_methodology,static_methodology_download,static_qmi` |

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### */datalist/data

| Version | URI                                                                                                                  |
|---------|----------------------------------------------------------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/datalist/data`                                                                               |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=timeseries,dataset_landing_page,timeseries_dataset,static_adhoc` |

For topic filtered versions of this page, for example /economy/datalist/data, you will need to [use topic IDs from our Topic API](#converting-topics-to-our-new-format).

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### */publications/data

| Version | URI                                                                                                            |
|---------|----------------------------------------------------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/*/publications/data`                                                                   |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=article,article_download,bulletin,compendium_landing_page` |

For topic filtered versions of this page, for example /economy/publications/data, you will need to [use topic IDs from our Topic API](#converting-topics-to-our-new-format).

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### /publishedrequests/data

| Version | URI                                                             |
|---------|-----------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/publishedrequests/data`                 |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=static_foi` |

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### */staticlist/data

| Version | URI                                                                                                 |
|---------|-----------------------------------------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/staticlist/data`                                                            |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=static_page,static_landing_page,static_article` |

For topic filtered versions of this page, for example /economy/staticlist/data, you will need to [use topic IDs from our Topic API](#converting-topics-to-our-new-format).

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### /timeseriestool/data

| Version | URI                                                             |
|---------|-----------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/timeseriestool/data`                    |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=timeseries` |

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### */topicspecificmethodology/data

| Version | URI                                                                                                            |
|---------|----------------------------------------------------------------------------------------------------------------|
| Old     | `https://www.ons.gov.uk/topicspecificmethodology/data`                                                         |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=static_methodology,static_methodology_download,static_qmi` |

For topic filtered versions of this page, for example /economy/topicspecificmethodology/data, you will need to [use topic IDs from our Topic API](#converting-topics-to-our-new-format).

See the [search response mapping](#search-response-mapping) for the overall response mapping and [search item data mapping](#search-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

### /releasecalendar/data

| Version | URI                                              |
|---------|--------------------------------------------------|
| Old     | `https://www.ons.gov.uk/releasecalendar/data`    |
| New     | `https://api.beta.ons.gov.uk/v1/search/releases` |

See the [release response mapping](#release-response-mapping) for the overall response mapping and [release item data mapping](#release-item-data-mapping) for the mappings for individual items. All of these endpoints offer the same parameters so please see [parameter mapping](#parameter-mapping) for how to work with them.

## Converting topics to our new format

Previously on our site we have used our URL structure to dictate the topics associated with content, but we are aiming to move away from this. In our Search Service, we offer content tagged by topic using unique IDs. These can be obtained from our Topic API:

`https://api.beta.ons.gov.uk/v1/topics`

## Parameter mapping

Parameters:

| Old parameter | New parameter                 |
|---------------|-------------------------------|
| page          | offset                        |
| size          | limit                         |
| fromDateDay   | fromDate (format: YYYY-MM-DD) |
| fromDateMonth | fromDate (format: YYYY-MM-DD) |
| fromDateYear  | fromDate (format: YYYY-MM-DD) |
| toDateDay     | toDate (format: YYYY-MM-DD)   |
| toDateMonth   | toDate (format: YYYY-MM-DD)   |
| toDateYear    | toDate (format: YYYY-MM-DD)   |
| topics        | topics                        |

See details about [converting topics to our new format](#converting-topics-to-our-new-format)

## Search response mapping

| Old JSON attribute      | New JSON attribute |
|-------------------------|--------------------|
| type                    | N/A                |
| listType                | N/A                |
| uri                     | N/A                |
| result.numberOfResults  | count              |
| result.took             | took               |
| result.results*         | items              |
| suggestions             | N/A                |
| docCounts               | N/A                |
| paginator.numberOfPages | N/A                |
| paginator.currentPage   | N/A                |
| paginator.start         | N/A                |
| paginator.end           | N/A                |
| paginator.pages         | N/A                |
| sortBy                  | N/A                |

*for mapping individual items please see the [search item data mapping](#search-item-data-mapping).

## Search item data mapping

| Old JSON attribute            | New JSON attribute |
|-------------------------------|--------------------|
| description.cdid              | cdid               |
| description.contact.email     | N/A*               |
| description.contact.name      | N/A*               |
| description.contact.telephone | N/A*               |
| description.datasetId         | dataset_id         |
| description.datasetUri        | N/A*               |
| description.date              | N/A*               |
| description.edition           | edition            |
| description.keyNote           | N/A*               |
| description.keywords          | keywords           |
| description.metaDescription   | meta_description   |
| description.nationalStatistic | N/A*               |
| description.nextRelease       | N/A*               |
| description.number            | N/A*               |
| description.preUnit           | N/A*               |
| description.releaseDate       | release_date       |
| description.sampleSize        | N/A*               |
| description.source            | N/A*               |
| description.title             | title              |
| description.summary           | summary            |
| description.unit              | N/A*               |
| searchBoost                   | N/A                |
| _type                         | type               |
| uri                           | uri                |

*Remaining content (with the exception of `searchBoost`) can be retrieved by using our `data` endpoint found at:

`https://api.beta.ons.gov.uk/v1/data`

and using the `uri` from the search response as a query parameter. For example:

`https://api.beta.ons.gov.uk/v1/data?uri=/economy`

The response attributes there directly correspond to the items in the table above.

## Release response mapping

| Old JSON attribute      | New JSON attribute |
|-------------------------|--------------------|
| type                    | N/A                |
| listType                | N/A                |
| uri                     | N/A                |
| result.numberOfResults  | breakdown.total    |
| result.took             | took               |
| result.results          | releases*          |
| suggestions             | N/A                |
| docCounts               | N/A                |
| paginator.numberOfPages | N/A                |
| paginator.currentPage   | N/A                |
| paginator.start         | N/A                |
| paginator.end           | N/A                |
| paginator.pages         | N/A                |
| sortBy                  | N/A                |

*for mapping individual items please see the [release item data mapping](#release-item-data-mapping).

## Release item data mapping

| Old JSON attribute             | New JSON attribute              |
|--------------------------------|---------------------------------|
| description.cancelled          | description.cancelled           |
| description.cancellationNotice | description.cancellation_notice |
| description.cdid               | cdid                            |
| description.contact.email      | N/A*                            |
| description.contact.name       | N/A*                            |
| description.contact.telephone  | N/A*                            |
| description.edition            | edition                         |
| description.finalised          | description.finalised           |
| description.metaDescription    | meta_description                |
| description.nationalStatistic  | N/A*                            |
| description.nextRelease        | N/A*                            |
| description.preUnit            | N/A*                            |
| description.published          | description.published           |
| description.releaseDate        | release_date                    |
| description.source             | N/A*                            |
| description.title              | title                           |
| description.summary            | summary                         |
| searchBoost                    | N/A                             |
| _type                          | type                            |
| uri                            | uri                             |

*Remaining content (with the exception of `searchBoost`) can be retrieved by using our `data` endpoint found at:

`https://api.beta.ons.gov.uk/v1/data`

and using the `uri` from the search response as a query parameter. For example:

`https://api.beta.ons.gov.uk/v1/data?uri=/economy`

The response attributes there directly correspond to the items in the table above.
