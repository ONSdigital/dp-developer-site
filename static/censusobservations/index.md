---
title: Guide to requesting Census 2021 observations
---

#### Census 2021 Datasets - Topic Summaries and Multivariate Datasets
To retrieve observations for Census 2021 datasets, the following URL structure is used:

    /datasets/{datasetId}/editions/{edition}/versions/{version}/json

This will retrieve observations for the default dataset which will consist of the area-type of lower tier local authority, and the variable associated with the topic summary e.g. Age.  An example URI is:

    /datasets/TS008/editions/2021/versions/1/json

It is possible to flex the geography and return results for different area-types and restrict the results to specific areas.  The following URI is an example of this, where TS008 results are returned for Wales only:

    /datasets/TS008/editions/2021/versions/1/json?area-type=ctry,W92000004

The area-type parameter values can be discovered by using the population-types api.  The area code to filter the geography which is after the comma in the URI can also be found by using the population-types api.
