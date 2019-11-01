---
title: Requesting from filter API
---

The ['Filter a dataset'](../filter/) service allows any combination of dimension options to be requested. This is aimed at more complicated queries than the ['Explore our dataset'](../dataset/) service and requires a POST containing dataset ID, edition, version and details of dimensions. 

Any dimensions not filtered on will return all available dimension items for that dataset. 
If you want everything in a dataset for a specific geographic location, you only need to provide the option for that geography, for example;

```
{
    "dataset": {
        "id": "cpih01",
        "edition": "time-series",
        "version": 6
    },
    "dimensions": [
        {
            "name": "geography",
            "options": [
                "K02000001"
            ]
        }
    ]
}
```

This example body can be sent via post to create a filter, with the response giving you a unique ID. As shown in this example under `filter_id`:

```
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
    "events": {},
    "filter_id": "d9645c21-0baa-4a58-834e-feb1919b14bb",
    "instance_id": "bc873fb8-0797-469b-84ed-9bb3da80feeb",
    "links": {
        "dimensions": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/d9645c21-0baa-4a58-834e-feb1919b14bb/dimensions"
        },
        "filter_blueprint": {},
        "filter_output": {},
        "self": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/d9645c21-0baa-4a58-834e-feb1919b14bb"
        },
        "version": {
            "href": "https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6",
            "id": "6"
        }
    },
    "published": true
}
```

Futher PUT requests can be made to `/filter/{filter-ID}` to update the dimensions required and once the request is complete it can be submitted by adding `?state=submitted`.

For example;
```
/filters/d9645c21-0baa-4a58-834e-feb1919b14bb?submitted=true
```

**Note**: `?submitted=true` can be added to the initial POST request and the job will start immediately.
```
/filters?submitted=true
```

Once submitted, you will get the following response:
```
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
```

This creates a 'filter output' which will contain the files once generated. This is available in the response and follows this format;
```
"filter_output": {
    "id": "9dd04dd7-07e9-4ce9-90bb-a4205b743027",
    "href": "https://api.beta.ons.gov.uk/v1/filter-outputs/9dd04dd7-07e9-4ce9-90bb-a4205b743027"
}
```

A GET on the filter-output will return the following;
```
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
```

**Note**: The split between 'filter' and 'filter-outputs' is there to allow the filter to be modified and resubmitted. 