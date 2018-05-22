---
title: Introduction
---

## Getting data from ONS Beta API

All URIs are relative to `https://api.beta.ons.gov.uk/v1`

### Observation data
The observation level data is provided in [JSON](https://www.w3schools.com/js/js_json_intro.asp) through GET requests to the Dataset API. This will allow querying of a single observation/value through providing one option per dimension but will also allow one of these dimensions to be a 'wildcard' and return all values for this dimension. 

This is currently limited to 10,000 observations.

[Requesting specific observation](observations)


### Filter API
For more complicated queries the Filter API will be the best approach. This allows any combination of dimensions to be requested and will return the data as CSV and Xlsx.

[Requesting from filter API](filters)