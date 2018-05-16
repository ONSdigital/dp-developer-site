---
title: Requesting from filter API
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
                    "id": "cpih1dim1A0",
                    "href": "http://localhost:22400/code-lists/cpih1dim1aggid/codes/cpih1dim1A0"
                }
            },
            "geography": {
                "option": {
                    "id": "K02000001",
                    "href": "http://localhost:22400/code-lists/uk-only/codes/K02000001"
                }
            },
            "time": {
                "option": {
                    "id": "Oct-11",
                    "href": "http://localhost:22400/code-lists/time/codes/Oct-11"
                }
            }
        },
        "limit": 10000,
        "links": {
            "dataset_metadata": {
                "href": "http://localhost:22000/datasets/cpih01/editions/time-series/versions/6/metadata"
            },
            "self": {
                "href": "http://localhost:22000/datasets/cpih01/editions/time-series/versions/6/observations?time=Oct-11&aggregate=cpih1dim1A0&geography=K02000001"
            },
            "version": {
                "id": "6",
                "href": "http://localhost:22000/datasets/cpih01/editions/time-series/versions/6"
            }
        },
        "offset": 0,
        "total_observations": 1,
        "observations": [
            {
                "observation": "94.5"
            }
        ]
    }

  
#### Time
The time dimension is treated slightly differently from all other dimensions in that we require the time label, rather than the IDs (all IDs for a certain 'type' of time, e.g Months, are the same)

#### Wildcard
The API allows a single dimension to be replaced with a wildcard to instead return all values for this dimension (up to a maximum of 10,0000).

Example URI

    /datasets/cpih01/editions/time-series/versions/6/observations?time=*&aggregate=cpih1dim1A0&geography=K02000001
