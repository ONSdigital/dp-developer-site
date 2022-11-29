---
title: Guide to requesting specific observations
---
#### Census 2021 Datasets
To retrieve observations for Census 2021 datasets, the following URL structure is used:

    /datasets/{datasetId}/editions/{edition}/versions/{version}/json

This will retrieve observations for the default dataset which will consist of the area-type of lower tier local authority, and the variable associated with the topic summary e.g. Age.  An example URI is:

    /datasets/TS008/editions/2021/versions/1/json

It is possible to flex the geography and return results for different area-types and restrict the results to specific areas.  The following URI is an example of this, where TS008 results are returned for Wales only:

    /datasets/TS008/editions/2021/versions/1/json?area-type=ctry,W92000004

The area-type parameter values can be discovered by using the population-types api.  The area code to filter the geography which is after the comma in the URI can also be found by using the population-types api.


#### CMD Datasets

Using the ['Explore our data'](../dataset/) services you can request data from a specific version of a dataset. This is done with a GET request providing a single dimension option for each dimension. The URL follows the following structure:

    /datasets/{datasetId}/editions/{edition}/versions/{version}/observations?time={timeLabel}&geography={geographyID}&dimension3={dimension3ID}&dimension4={dimension4ID}...

Example URI;

    /datasets/cpih01/editions/time-series/versions/6/observations?time=Oct-11&geography=K02000001&aggregate=cpih1dim1A0

As well as the requested data, the response provides links to [code lists](../code-list/) for individual items requested, any observation level metadata/information (e.g coefficients of variation) and links to original dataset and complete metadata.

    {
        "dimensions": {
            "aggregate": {
                "option": {
                    "href": "https://api.beta.ons.gov.uk/v1/code-lists/cpih1dim1aggid/codes/cpih1dim1A0",
                    "id": "cpih1dim1A0"
                }
            },
            "geography": {
                "option": {
                    "href": "https://api.beta.ons.gov.uk/v1/code-lists/uk-only/codes/K02000001",
                    "id": "K02000001"
                }
            },
            "time": {
                "option": {
                    "href": "https://api.beta.ons.gov.uk/v1/code-lists/time/codes/Mar-18",
                    "id": "Mar-18"
                }
            }
        },
        "limit": 10000,
        "links": {
            "dataset_metadata": {
                "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6/metadata"
            },
            "self": {
                "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6/observations?time=Mar-18&aggregate=cpih1dim1A0&geography=K02000001"
            },
            "version": {
                "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6",
                "id": "6"
            }
        },
        "observations": [
            {
                "observation": "105.1"
            }
        ],
        "offset": 0,
        "total_observations": 1,
        "unit_of_measure": "Index: 2015=100"
    }

#### Time

The time dimension is treated slightly differently from all other dimensions and you need to use the time label, rather than the ID. This is beacause all IDs for a certain 'type' of time, e.g Months, are currently the same.

#### Wildcard

This endpoint allows a single dimension to be replaced with a wildcard ('*') to return all values for this dimension.

Example URI

    /datasets/cpih01/editions/time-series/versions/6/observations?time=*&aggregate=cpih1dim1A0&geography=K02000001

This is currently limited to 10,000 observations.
