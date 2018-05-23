---
title: Requesting from filter API
---

POST to filter API with dataset ID ("dataset_filter_id") and body containing details of dimensions. 

    /filters?submitted=true

Using 'submitted=true' will immediately submit the job to generate the download files

Any dimensions not filtered on will return all available dimension items for that dataset. 
If you want everything in a dataset for a specific geographic location, you only need to provide the option for that geography, see below;

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

Responce

    {
        "dataset": {
            "id": "cpih01",
            "edition": "time-series",
            "version": 6
        },
        "instance_id": "2426ff0f-ca8b-45f6-8728-e851ea165b85",
        "dimensions": [
            {
                "name": "geography",
                "options": [
                    "K02000001"
                ]
            }
        ],
        "events": {},
        "filter_id": "63c77ab3-b57d-40a8-8397-9830ed98f9fd",
        "published": true,
        "links": {
            "dimensions": {
                "href": "https://api.beta.ons.gov.uk/v1/filters/63c77ab3-b57d-40a8-8397-9830ed98f9fd/dimensions"
            },
            "filter_output": {
                "id": "9dd04dd7-07e9-4ce9-90bb-a4205b743027",
                "href": "https://api.beta.ons.gov.uk/v1/filter-outputs/9dd04dd7-07e9-4ce9-90bb-a4205b743027"
            },
            "filter_blueprint": {},
            "self": {
                "href": "https://api.beta.ons.gov.uk/v1/filters/63c77ab3-b57d-40a8-8397-9830ed98f9fd"
            },
            "version": {
                "id": "6",
                "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6"
            }
        }
    }

This creates a 'filter output' which will contain the files once generated. This is available in the response and follows this format;

    "filter_output": {
                    "id": "9dd04dd7-07e9-4ce9-90bb-a4205b743027",
                    "href": "https://api.beta.ons.gov.uk/v1/filter-outputs/9dd04dd7-07e9-4ce9-90bb-a4205b743027"
                },

A GET on the filter-output will return the following;

    {
    "dataset": {
        "edition": "time-series",
        "id": "cpih01",
        "version": 6
    },
    "dimensions": [
        {
            "name": "geography",
            "options": [
                "K02000001"
            ]
        }
    ],
    "downloads": {
        "csv": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/a5deb5e7-bb91-47bc-bb70-f9da12d54fa1.csv",
            "size": "2061369"
        },
        "xls": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/a5deb5e7-bb91-47bc-bb70-f9da12d54fa1.xlsx",
            "size": "112878"
        }
    },
    "events": {},
    "filter_id": "a5deb5e7-bb91-47bc-bb70-f9da12d54fa1",
    "instance_id": "bc873fb8-0797-469b-84ed-9bb3da80feeb",
    "links": {
        "dimensions": {},
        "filter_blueprint": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/5ce923e2-613e-4dcf-874c-d750d97afe79",
            "id": "5ce923e2-613e-4dcf-874c-d750d97afe79"
        },
        "filter_output": {},
        "self": {
            "href": "https://api.beta.ons.gov.uk/v1/filter-outputs/a5deb5e7-bb91-47bc-bb70-f9da12d54fa1"
        },
        "version": {
            "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6",
            "id": "6"
        }
    },
    "published": true,
    "state": "completed"
    }


The split between 'filter' and 'filter-outputs' is there to allow the filter to be modified and resubmitted. 