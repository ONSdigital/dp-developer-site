// For reference

window.onload = () => {
    document.getElementById("tab-results").addEventListener('click', function () {
        document.getElementById("tryitout-list-datasets-result-result")
            .classList.remove("hidden");
        document.getElementById("tryitout-list-datasets-json-result").classList
            .add("hidden");
        document.getElementById("tryitout-list-datasets-json-result").classList
            .remove("btn--secondary-active");
        document.getElementById("tab-results").classList.add(
            "btn--secondary-active");
        document.getElementById("tab-json").classList.remove(
            "btn--secondary-active");
    });
    document.getElementById("tab-json").addEventListener('click', function () {
        document.getElementById("tryitout-list-datasets-result-result")
            .classList.add("hidden");
        document.getElementById("tryitout-list-datasets-json-result").classList
            .remove("hidden");
        document.getElementById("tab-results").classList.remove(
            "btn--secondary-active");
        document.getElementById("tab-json").classList.add(
            "btn--secondary-active");
    });

    document.getElementById("tryitout-list-datasets").addEventListener('click',
        function () {
            document.getElementById("tryitout-list-datasets").setAttribute(
                "disabled", "");
            document.getElementById("tryitout-list-datasets-results").setAttribute(
                "class", "hidden");
            document.getElementById("tryitout-list-datasets").setAttribute("class",
                "btn btn--primary btn--primary-disabled"
                );

            fetch("https://api.beta.ons.gov.uk/v1/datasets")
                .then((r) => {
                    return r.json();
                })
                .then((r) => {
                    document.getElementById("tryitout-list-datasets-results")
                        .removeAttribute("class");
                    document.getElementById("tryitout-list-datasets")
                        .removeAttribute("disabled");
                    document.getElementById("tryitout-list-datasets")
                        .setAttribute("class", "hidden");

                    document.getElementById("tryitout-list-datasets-json")
                        .innerHTML = "<code>" + JSON.stringify(r, null, 2); +
                    "</code>"
                    let table = document.getElementById(
                        "tryitout-list-datasets-table");
                    table.innerHTML = "";
                    table.setAttribute("class", "table");
                    r.items.forEach(dataset => {
                        let row = document.createElement("tr");
                        let datasetID = document.createElement("td");
                        datasetID.innerText = dataset.id;
                        row.appendChild(datasetID);
                        let datasetName = document.createElement("td");
                        datasetName.innerText = dataset.title;
                        row.appendChild(datasetName);
                        table.appendChild(row);
                    });
                })
                .catch((e) => {
                    console.log(e);
                });
        });
};