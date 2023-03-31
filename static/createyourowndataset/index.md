---
title: Guide to Creating your own dataset - Census 2021
---

#### Create a Custom Dataset
To create a custom dataset for Census 2021 data, there are several endpoints available to build up the query.  All the endpoints below are GET requests.

Firstly, it should be determined which population the query is to be based on.  The list of available populations can be accessed using:
     
     /population-types  
     
Only those populations returned which are of type microdata can be used for a create your own dataset query.

Once the population has been established, an area-type to run the query for should be selected.  A list of available area-types can be found on the following URL, replacing {population-type-name} with the name of the selected population:

    /population-types/{population-type-name}/area-types

It is possible to filter a custom dataset query to specific areas.  To access a list of available areas for an area-type, then the following URL pattern should be used, replacing {population-type-name} with the selected population and {area-type-id} with the id of selected area-type.

    /population-types/{population-type-name}/area-types/{area-type-id}/areas

If a "q" parameter is added to the `/areas` request, then a search is carried out on the supplied text in the parameter, e.g. `/population-types/UR/area-types/ltla/areas?q=swa` will return only lower tier local authorities with the text "swa" in the label.

The next step is to determine which variables are available to add to the query.  These can be accessed using the following URL pattern which returns all base variables for a population:

    /population-types/{population-type-name}/dimensions

If a "q" parameter is added to the `/dimensions` request, a search is carried out on the supplied text in the parameter, e.g. `/population-types/UR/dimensions?q=age` will return only variables with the text "age" in the label.

If a different level of categorisation is required for a query to be run against, these levels of categrisation can be found at the following URL pattern, replacing {population-type-name} with the name of the population and the {dimension-id} with the id returned in the `/dimensions` response:

    /population-types/{population-type-name}/dimensions/{dimension-id}/categorisations

Once all the above have been established, the `/census-observations` endpoint can be used to create a custom dataset query.  For example, the URL below returns information about general health and the highest level of qualification within the usual residents population at area-type level of country for England only:

    /population-types/UR/census-observations?area-type=ctry,E92000001&dimensions=health_in_general,highest_qualification

The area-type is supplied in the area-type parameter, and any areas to be filtered should be added in commas after the area-type-id.  Any variables should be supplied in the dimensions parameter in a comma separated list.  These are the dimension-ids returned in the `/dimensions` and `/categorisations` endpoints above.
