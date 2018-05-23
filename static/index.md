---
title: Introduction
---

The Office for National Statistics API makes datasets and other data available programmatically using HTTP. It allows you to filter datasets and directly access specific data points.

The API is available at     `https://api.beta.ons.gov.uk/v1`

The API is open and unrestricted - no API keys are required, so you can start using it immediately.

## Getting data from ONS Beta API

### Requesting data directly
Observation level data is provided in [JSON](https://www.w3schools.com/js/js_json_intro.asp) through GET requests to the [Dataset API](https://developer.beta.ons.gov.uk/dataset/). This allows querying of a single observation/value through providing one option per dimension but will also allow one of these dimensions to be a 'wildcard' and return all values for this dimension. 

[Guide to requesting specific observation](/observations)


### Filtering a dataset
For more complicated queries the [Filter API](https://developer.beta.ons.gov.uk/filter/) will be the best approach. This allows any combination of dimensions to be requested and will return the data as CSV and Xlsx.

[Guide to requesting from filter API](/filters)