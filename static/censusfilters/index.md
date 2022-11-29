---
title: Guide to filtering a Census 2021 dataset
---

If you require download files for Census 2021 data to be generated via the API then the filter service should be used.  Otherwise, the ['Observations'](../observations/) service may be more appropriate for instant access to data.

Using the filter service for Census 2021 data is similar to the ['CMD Filter'](../filters/) filter journey, but more fields are required to begin the filter process.

An example POST request is as follows, this example will create a filter against dataset TS008 for the region of North East.

```
{
    "dataset": {
        "id": "TS008",
        "edition": "2021",
        "version": 1
    },
    "population_type":"UR",
    "dimensions": [
        {
            "name": "rgn",
            "is_area_type":true,
            "options": [
                "E12000001"
            ]
        },
        {
            "name": "sex"
        }
    ]
}
```

The response will provide you with a unique ID. As shown in this example under `filter_id`:

```
{
    "filter_id": "2deaa7ef-8dcc-40a0-b739-96a919b35cba",
    "instance_id": "1e2ce748-b2a2-4eb0-aa52-7b361f3f5899",
    "dataset": {
        "id": "TS008",
        "edition": "2021",
        "version": 1,
        "lowest_geography": "",
        "release_date": "2022-11-10T00:00:00.000Z",
        "title": "Sex"
    },
    "published": true,
    "type": "flexible",
    "population_type": "UR",
    "links": {
        "version": {
            "href": "https://api.beta.ons.gov.uk/v1/datasets/TS008/editions/2021/version/1",
            "id": "1"
        },
        "self": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/2deaa7ef-8dcc-40a0-b739-96a919b35cba"
        },
        "dimensions": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/2deaa7ef-8dcc-40a0-b739-96a919b35cba/dimensions"
        }
    }
}
```

This filter record can be further amended if required, by using the POST and DELETE /options endpoints to filter the record to specific configurations required.

To submit the filter job to begin the file generation process, the POST /submit endpoint should be used, e.g.

```
/filters/2deaa7ef-8dcc-40a0-b739-96a919b35cba/submit
```

Once submitted, you will get the following response:

```
{
    "instance_id": "1e2ce748-b2a2-4eb0-aa52-7b361f3f5899",
    "filter_output_id": "f70121c4-9464-4e48-b790-a776623f0b8f",
    "dataset": {
        "id": "TS008",
        "edition": "2021",
        "version": 1,
        "lowest_geography": "",
        "release_date": "2022-11-10T00:00:00.000Z",
        "title": "Sex"
    },
    "links": {
        "version": {
            "href": "https://api.beta.ons.gov.uk/v1/datasets/TS008/editions/2021/version/1",
            "id": "1"
        },
        "self": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/2deaa7ef-8dcc-40a0-b739-96a919b35cba"
        },
        "dimensions": {
            "href": "https://api.beta.ons.gov.uk/v1/filters/2deaa7ef-8dcc-40a0-b739-96a919b35cba/dimensions"
        }
    },
    "population_type": "UR"
}
```

This creates a 'filter output' which will contain the files once generated. This is available in the response and follows this format;

```
  "filter_output_id": "f70121c4-9464-4e48-b790-a776623f0b8f"
```

A GET on the filter-output, e.g. https://api.beta.ons.gov.uk/v1/filter-outputs/f70121c4-9464-4e48-b790-a776623f0b8f will return the following;

```
{
    "dataset": {
        "edition": "2021",
        "id": "TS008",
        "lowest_geography": "",
        "release_date": "2022-11-10T00:00:00.000Z",
        "title": "Sex",
        "version": 1
    },
    "dimensions": [
        {
            "id": "rgn",
            "is_area_type": true,
            "label": "Regions",
            "name": "rgn",
            "options": [
                "E12000001"
            ]
        },
        {
            "id": "sex",
            "label": "Sex (2 categories)",
            "name": "sex",
            "options": []
        }
    ],
    "downloads": {
        "csv": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/f70121c4-9464-4e48-b790-a776623f0b8f.csv",
            "private": "",
            "public": "https://static.ons.gov.uk/datasets/f70121c4-9464-4e48-b790-a776623f0b8f/TS008-2021-1-filtered-2022-11-29T10:55:37Z.csv",
            "size": "150",
            "skipped": false
        },
        "csvw": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/f70121c4-9464-4e48-b790-a776623f0b8f.csv-metadata.json",
            "private": "",
            "public": "https://static.ons.gov.uk/datasets/f70121c4-9464-4e48-b790-a776623f0b8f/TS008-2021-1-2022-11-29T10:55:37Z.csvw",
            "size": "1987",
            "skipped": false
        },
        "txt": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/f70121c4-9464-4e48-b790-a776623f0b8f.txt",
            "private": "",
            "public": "https://static.ons.gov.uk/datasets/f70121c4-9464-4e48-b790-a776623f0b8f/TS008-2021-1-2022-11-29T10:55:38Z.txt",
            "size": "3036",
            "skipped": false
        },
        "xls": {
            "href": "https://download.beta.ons.gov.uk/downloads/filter-outputs/f70121c4-9464-4e48-b790-a776623f0b8f.xlsx",
            "private": "",
            "public": "https://static.ons.gov.uk/datasets/f70121c4-9464-4e48-b790-a776623f0b8f/TS008-2021-1-filtered-2022-11-29T10:55:37Z.xlsx",
            "size": "8257",
            "skipped": false
        }
    },
    "events": [],
    "filter_id": "2deaa7ef-8dcc-40a0-b739-96a919b35cba",
    "id": "f70121c4-9464-4e48-b790-a776623f0b8f",
    "instance_id": "",
    "links": {
        "filter_blueprint": {
            "href": "https://api.beta.ons.gov.uk/v1/filters",
            "id": "2deaa7ef-8dcc-40a0-b739-96a919b35cba"
        },
        "self": {
            "href": "https://api.beta.ons.gov.uk/v1/filter-outputs",
            "id": "f70121c4-9464-4e48-b790-a776623f0b8f"
        },
        "version": {
            "href": "https://api.beta.ons.gov.uk/v1/datasets/TS008/editions/2021/version/1",
            "id": "1"
        }
    },
    "population_type": "UR",
    "published": true,
    "type": "flexible"
}
```

**Note**: The split between 'filter' and 'filter-outputs' is there to allow the filter to be modified and resubmitted.
