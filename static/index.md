---
title: Introduction
---

## Getting data from API.BETA.ONS.GOV.UK

### Observation data
The observation level data is available in JSON through GET requests to the Dataset API. This will allow querying of a single observation through providing one dimension value per dimension but will also allow one of these dimensions to be a 'wildcard' and return all values for this dimension. 

This is currently limited to 10,000 results.

[Requesting specific observation](#requesting-specific-observation)


### Filter API
For more complicated queries the Filter API will be the best approach. This allows any combination of dimensions to be requested and will return the data as CSV and Xlsx.

[Requesting from filter API](#requesting-from-filter-api)

----------

### Requesting specific observation

GET on the dataset providing a single dimension option for each dimension.

For example;

    https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/current/versions/1/observation?geography=E07000224&aggregate=cpi1dim1A0&time=Jan-96

Provides links to codelist for individual items requested, observation value any observation level metadata/information (e.g coefficients of variation) and links to original dataset and complete metadata.

    {
        “geography”:”https://api.ons.gov.uk/codelist/geography/local-authority/E07000224”,
        “aggregate”:”https://api.ons.gov.uk/codelist/coicop/cpi1dim1A0”,
        “time”=”https://api.ons.gov.uk/codelist/time/Jan-96”,
        “observation”:”2.5”,
        “observation_level_metadata”: [
        <key>: string
        ]
        “coefficients of variation”:””,
        “data_marking”:”p”,
        “context”:””,
        “dataset_metadata”:“href”:”https://beta.ons.gov.uk/datasets/cpih01/editions/current/versions/1/”,
        "links": {
            "version": [
            "href": "https://beta.ons.gov.uk/datasets/cpih01/editions/2017/version/1",
            "id": "1"
            }
        ],
            "user_notes": [
                {
                    "title": "Coefficient of variation (CV)",
                    "note": "The CV column the quality of each estimate based on the coefficient of variation (CV) of that estimate. The CV is the ratio of the standard error of an estimate to the estimate itself and is expressed as a percentage. The smaller the coefficient of variation the greater the accuracy of the estimate. The true value is likely to lie within +/- twice the CV. For example, for an estimate of £200 with a CV of 5%, we would expect the true population average to be within the range £180 to £220."
                },
                {
                    "title": "Data Markings",
                    "note":"Estimates with a CV greater than 20% are suppressed from publication on quality grounds, along with those for which there is a risk of disclosure of individual employees or employers. This is shown as an 'x' in the data marking column"
                }
            ]
        }
    }
#### Time
The time dimension is treated slightly differently from all other dimensions in that we require the time label, rather than the IDs (all IDs for a certain 'type' of time, e.g Months, are the same)

#### Wildcard
The API allows a single dimension to be replaced with a wildcard to instead return all values for this dimension (up to a maximum of 10,0000).

For example;

    https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6/observations?time=*&aggregate=cpih1dim1A0&geography=K02000001

### Requesting from filter API

POST to filter API with dataset ID ("dataset_filter_id") and body containing details of dimensions. 

    https://api.beta.ons.gov.uk/v1/filters?submitted=true

Using 'submitted=true' will immediately submit the job to generate the download files

Any dimensions not filtered on will return all available dimension items for that dataset. 
E.G. If you want everything in a dataset for a specific geographic location, you only need to provide the option for that geography, see below;

    {
      "dataset": {
		    "id": "cpih01",
		    "edition": "time-series",
		    "version": "6"
			},
      "dimensions": [
        {
          "name": "Geography",
          "options": [
            "E07000224"
          ]
        }
      ]
    }


Once submitted this filter will run the query and return the following JSON.

    {
      "downloads": {
        "xls": {
          "url": "https://api.beta.ons.gov.uk/datasets/cpih01/editions/current/versions/1/ffe3bc0b6-d6c4-e20-917e-95d7ea8c91dc/downloads/xls",
          "size": "35Kb"
        },
        "json": {
          "url": "https://api.ons.gov.uk/datasets/cpih01/editions/current/versions/1/ffe3bc0b6-d6c4-e20-917e-95d7ea8c91dc/downloads/json",
          "size": "5Kb"
        },
        "csv": {
          "url": "https://api.ons.gov.uk/datasets/cpih01/editions/current/versions/1/ffe3bc0b6-d6c4-e20-917e-95d7ea8c91dc/downloads/csv",
          "size": "40Kb"
        }
      },
      "filter_id": "ffe3bc0b6-d6c4-e20-917e-95d7ea8c91dc",
      "instance_id": "he3bcjb6-d6c4-4e20-917e-95d7ea8c91dc",
      "dimension_list_url": "string",
    “metadata”:“href”:”https://api.ons.gov.uk/datasets/cpih01/editions/current/versions/1/metadata”,
      "links": {
        "version": {
          "href": "https://beta.ons.gov.uk/datasets/cpih01/editions/2017/version/1",
          "id": "1"
        }
      }
    }
