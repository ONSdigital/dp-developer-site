---
title: Requesting data directly
---

GET on the dataset providing a single dimension option for each dimension. The URL follows the following structure:    

    /datasets/{datasetId}/editions/{edition}/versions/{version}/observations?time={timeLabel}&geography={geographyID}&dimension3={dimension3ID}&dimension4={dimension4ID}...

Example URI;

    /datasets/cpih01/editions/time-series/versions/6/observations?time=Oct-11&geography=K02000001&aggregate=cpih1dim1A0



As well as the requested data, the response provides links to code-lists for individual items requested, any observation level metadata/information (e.g coefficients of variation) and links to original dataset and complete metadata.

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
Theis endpoint allows a single dimension to be replaced with a wildcard (a '*') to instead return all values for this dimension.

Example URI

    /datasets/cpih01/editions/time-series/versions/6/observations?time=*&aggregate=cpih1dim1A0&geography=K02000001

This is currently limited to 10,000 observations.