// For reference

let latest_version_url = "";

    window.onload = () => {
        document.getElementById("tab-results").addEventListener('click', function(){
            document.getElementById("tryitout-list-datasets-result-result").classList.remove("hidden");
            document.getElementById("tryitout-list-datasets-json-result").classList.add("hidden");
            document.getElementById("tryitout-list-datasets-json-result").classList.remove("btn--secondary-active");
            document.getElementById("tab-results").classList.add("btn--secondary-active");
            document.getElementById("tab-json").classList.remove("btn--secondary-active");
        });
        document.getElementById("tab-json").addEventListener('click', function(){
            document.getElementById("tryitout-list-datasets-result-result").classList.add("hidden");
            document.getElementById("tryitout-list-datasets-json-result").classList.remove("hidden");
            document.getElementById("tab-results").classList.remove("btn--secondary-active");
            document.getElementById("tab-json").classList.add("btn--secondary-active");
        });

        document.getElementById("tryitout-get-dataset").addEventListener('click', function(){
            document.getElementById("tryitout-get-dataset").setAttribute("disabled", "");
            document.getElementById("tryitout-get-dataset-results").setAttribute("class", "hidden");
            document.getElementById("tryitout-get-dataset").setAttribute("class", "btn btn--primary btn--primary-disabled");

            fetch("https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6")
                .then((r) => {
                    return r.json();
                })
                .then((r) => {
                    document.getElementById("tryitout-get-dataset-results").removeAttribute("class");
                    document.getElementById("tryitout-get-dataset").removeAttribute("disabled");
                    document.getElementById("tryitout-get-dataset").setAttribute("class", "btn btn--primary");

                    document.getElementById("tryitout-get-dataset-json").innerText = JSON.stringify(r, null, 2);

                    let table = document.getElementById("tryitout-get-version-table");
                    table.innerHTML = "";
                    table.setAttribute("class", "table");

                    ["release_date", "edition", ].forEach(field => {
                        let row = document.createElement("tr");
                        let fieldTitle = document.createElement("td");
                        fieldTitle.innerText = field;
                        row.appendChild(fieldTitle);
                        let fieldValue = document.createElement("td");
                        fieldValue.innerText = r[field];
                        row.appendChild(fieldValue);
                        table.appendChild(row);
                    });
                })
                .catch((e) => {
                    console.log(e);
                });
        });
    };