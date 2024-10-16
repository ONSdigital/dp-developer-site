---
title: Retirement of v0 API
---

## Background

In line with our [approach to the retirement of API endpoints](../), it has been decided that ONS will retire our 'v0
API' due to the following reasons:

- out of date technology
- built as a proof of concept
- strategic technology direction is elsewhere
- cost of upkeep is too high with similar services available

This application provides the following API endpoints:

- `/dataset`
- `/dataset/{dataset_id}`
- `/dataset/{dataset_id}/timeseries`
- `/dataset/{dataset_id}/timeseries/{timeseries_id}`
- `/dataset/{dataset_id}/timeseries/{timeseries_id}/data`
- `/timeseries`
- `/timeseries/{timeseries_id}/dataset`
- `/search`

These are for the following domains:

- `api.beta.ons.gov.uk`
- `api.ons.gov.uk`

This is not the case for our v1 beta API can be accessed at:

- `api.beta.ons.gov.uk/v1`

Which has, for example, endpoints for:

- `/search`
- `/datasets`

You can also browse the [full Search API Swagger specification](../../search/search/).

If you wish to discuss the retirement of the V0 API with us or would like assistance using this migration guide, please email us at <apiservice@ons.gov.uk>

## Migration guide

Below you can find detailed ways to access the same data from the ONS using our latest services which will benefit from
continued enhancement and investment.

### /dataset

| Version | URI                                                          |
|---------|--------------------------------------------------------------|
| Old     | `{domain}/dataset`                                           |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=dataset` |

Parameters:

| Old parameter | New parameter |
|---------------|---------------|
| start         | offset        |
| limit         | limit         |

You will need to add a new parameter to the request to just return the `dataset` content type:

`content_type=dataset`

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

### /dataset/{dataset\_id}

| Version | URI                             |
|---------|---------------------------------|
| Old     | `{domain}/dataset/{dataset_id}` |
| New     | Deprecated                      |

This endpoint currently returns a 404 for all `dataset_id` provided and so can already be considered deprecated.

In order to get at similar data, you could still use our Search API:

```txt
https://api.beta.ons.gov.uk/v1/search?content_type=dataset&dataset_ids={dataset_id}
```

You also have the option of filtering for multiple dataset ids at once

```txt
https://api.beta.ons.gov.uk/v1/search?content_type=dataset&dataset_ids={dataset_id1,dataset_id2}
```

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

### /dataset/{dataset\_id}/timeseries

| Version | URI                                                                                      |
|---------|------------------------------------------------------------------------------------------|
| Old     | `{domain}/dataset/{dataset_id}/timeseries`                                               |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=timeseries&dataset_ids={dataset_id}` |

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

### /dataset/{dataset\_id}/timeseries/{timeseries\_id} and /dataset/{dataset\_id}/timeseries/{timeseries\_id}/data

| Version | URI                                                                          |
|---------|------------------------------------------------------------------------------|
| Old     | `{domain}/dataset/{dataset_id}/timeseries/{timeseries_id}`                   |
| New     | `https://api.beta.ons.gov.uk/v1/search?content_type=timeseries&cdids={timeseries_id}` |

`/dataset/{dataset_id}/timeseries/{timeseries_id}` is a subset of what is retrieved
via `/dataset/{dataset_id}/timeseries/{timeseries_id}/data`.

After getting the response from our search service, you can get the `uri` parameter of the `item` that matches
your `timeseries_id` to it's `cdids`. Using that `uri` you can then submit that to our data endpoint, for example:

`https://api.beta.ons.gov.uk/v1/data?uri=/economy/nationalaccounts/uksectoraccounts/timeseries/mm23/capstk`

These uris are currently evergreen and so won't change between editions.

### /timeseries

| Version | URI                                                              |
|---------|------------------------------------------------------------------|
| Old     | `{domain}/timeseries`                                            |
| New     | `https://api.beta.ons.gov.uk/v1/search?&content_type=timeseries` |

Parameters:

| Old parameter | New parameter |
|---------------|---------------|
| start         | offset        |
| limit         | limit         |

You will need to add a new parameter to the request to just return the `timeseries` content type:

`content_type=timeseries`

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

### /timeseries/{timeseries\_id}

| Version | URI                                   |
|---------|---------------------------------------|
| Old     | `{domain}/timeseries/{timeseries_id}` |
| New     | Deprecated                            |

There is not a direct replacement for this URI but the URIs for timeseries pages can be derived from search:

`https://api.beta.ons.gov.uk/v1/search?content_type=timeseries&cdids={timeseries_id}`

You can also filter by multiple timeseries IDs:

`https://api.beta.ons.gov.uk/v1/search?content_type=timeseries&cdids={timeseries_id1,timeseries_id2}`

Using the `uri` parameter of the `item` that matches `timeseries_id` to `cdid` you can then submit that to our data
endpoint, for example:

`https://api.beta.ons.gov.uk/v1/data?uri=/economy/nationalaccounts/uksectoraccounts/timeseries/mm23/capstk`

These uris are currently evergreen and so won't change between releases.

### /timeseries/{timeseries\_id}/dataset

| Version | URI                                           |
|---------|-----------------------------------------------|
| Old     | `{domain}/timeseries/{timeseries_id}/dataset` |
| New     | Deprecated                                    |

This endpoint currently returns a 404 for all `timeseries_id` provided and so can already be considered deprecated.

To get a result to what might be expected for this endpoint, you can go to:

`https://api.beta.ons.gov.uk/v1/search?content_type=timeseries&cdids={timeseries_id}`

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

### /search

| Version | URI                                                                      |
|---------|--------------------------------------------------------------------------|
| Old     | `{domain}/search`                                                        |
| New     | `https://api.beta.ons.gov.uk/v1/search?&content_type=timeseries,dataset` |

Parameters:

| Old parameter | New parameter |
|---------------|---------------|
| start         | offset        |
| limit         | limit         |
| q             | q             |

Add new parameters:

`content_type=timeseries,dataset`

See the [search response mapping](#search-response-mapping) for the overall response mapping
and [search item data mapping](#search-item-data-mapping) for the mappings for individual items.

## Search response mapping

| Old JSON attribute | New JSON attribute |
|--------------------|--------------------|
| startIndex         | N/A                |
| itemsPerPage       | N/A                |
| totalItems         | count              |
| items              | items              |

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
| type                          | type               |
| uri                           | uri                |

*Remaining content (with the exception of `searchBoost`) can be retrieved by using our `data` endpoint found at:

`https://api.beta.ons.gov.uk/v1/data`

and using the `uri` from the search response as a query parameter. For example:

`https://api.beta.ons.gov.uk/v1/data?uri=/economy`

The response attributes there directly correspond to the items in the table above.
