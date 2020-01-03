// For reference

function orderByDate(a, b) {
    if (a[2] < b[2])
        return -1;
    if (a[2] > b[2])
        return 1;
    return 0;
}

window.onload = () => {
    document
        .getElementById("tab-results")
        .addEventListener('click', function () {
            document
                .getElementById("tryitout-get-observation-result-result")
                .classList
                .remove("hidden");
            document
                .getElementById("tryitout-get-observation-json-result")
                .classList
                .add("hidden");
            document
                .getElementById("tryitout-get-observation-json-result")
                .classList
                .remove("btn--secondary-active");
            document
                .getElementById("tab-results")
                .classList
                .add("btn--secondary-active");
            document
                .getElementById("tab-json")
                .classList
                .remove("btn--secondary-active");
        });
    document
        .getElementById("tab-json")
        .addEventListener('click', function () {
            document
                .getElementById("tryitout-get-observation-result-result")
                .classList
                .add("hidden");
            document
                .getElementById("tryitout-get-observation-json-result")
                .classList
                .remove("hidden");
            document
                .getElementById("tab-results")
                .classList
                .remove("btn--secondary-active");
            document
                .getElementById("tab-json")
                .classList
                .add("btn--secondary-active");
        });

    document
        .getElementById("tryitout-get-observation")
        .addEventListener('click', function () {
            fetch("https://api.beta.ons.gov.uk/v1/datasets/cpih01")
                .then((r) => {
                    return r.json();
                })
                .then((r) => {
                    document
                        .getElementById("title")
                        .innerText = r.title;
                })
                .catch((e) => {
                    console.log(e);
                });

            fetch("https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6/observations?time=*&geography=K02000001&aggregate=cpih1dim1A0")
                .then((r) => {
                    return r.json();
                })
                .then((r) => {
                    document
                        .getElementById("tryitout-get-observation-results")
                        .removeAttribute("class");
                    document
                        .getElementById("tryitout-get-observation")
                        .removeAttribute("disabled");
                    document
                        .getElementById("tryitout-get-observation")
                        .setAttribute("class", "btn btn--primary");

                    document
                        .getElementById("tryitout-get-observation-json")
                        .innerText = JSON.stringify(r, null, 2);

                    //create an array to story data
                    let timeseries = [];

                    //build a map of the dimensions
                    r
                        .observations
                        .map(function (data) {
                            //create a new field from the date to allow it to  be ordered
                            sort = new Date("1-" + data.dimensions.time.label.replace("-", " "));
                            //build an array with the 3 values we need and convert the value to a number
                            let chartdata = [
                                data.dimensions.time.label,
                                parseFloat(data.observation),
                                sort
                            ];
                            //add these arrays to the timeseries array
                            timeseries.push(chartdata);
                        });

                    //sort the array on the 'sort' field created earlier
                    timeseries = timeseries.sort(orderByDate);

                    //create the chart
                    Highcharts.chart('chart', {
                        series: [
                            {
                                //add the data to the chart
                                data: timeseries
                            }
                        ],
                        navigation: {
                            buttonOptions: {
                                enabled: false
                            }
                        },
                        title: {
                            text: null
                        },
                        yAxis: {
                            labels: {
                                enabled: false
                            },
                            title: {
                                enabled: false
                            },
                            // min: 0
                        },
                        legend: {
                            enabled: false
                        },
                        xAxis: {
                            type: "category",
                            crosshair: true
                        }
                    });
                })
                .catch((e) => {
                    console.log(e);
                });
        });

};