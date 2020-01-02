let latest_version_url = "";

        window.onload = () => {
            document.getElementById("tab-results").addEventListener('click', function () {
                document.getElementById("tryitout-get-observation-result-result").classList.remove("hidden");
                document.getElementById("tryitout-get-observation-json-result").classList.add("hidden");
                document.getElementById("tryitout-get-observation-json-result").classList.remove("btn--secondary-active");
                document.getElementById("tab-results").classList.add("btn--secondary-active");
                document.getElementById("tab-json").classList.remove("btn--secondary-active");
            });
            document.getElementById("tab-json").addEventListener('click', function () {
                document.getElementById("tryitout-get-observation-result-result").classList.add("hidden");
                document.getElementById("tryitout-get-observation-json-result").classList.remove("hidden");
                document.getElementById("tab-results").classList.remove("btn--secondary-active");
                document.getElementById("tab-json").classList.add("btn--secondary-active");
            });

            document.getElementById("tryitout-get-observation").addEventListener('click', function () {
                fetch("https://api.beta.ons.gov.uk/v1/datasets/cpih01/editions/time-series/versions/6/observations?time=Aug-16&geography=K02000001&aggregate=cpih1dim1A0")
                    .then((r) => {
                        return r.json();
                    })
                    .then((r) => {
                        document.getElementById("tryitout-get-observation-results").removeAttribute("class");
                        document.getElementById("tryitout-get-observation").removeAttribute("disabled");
                        document.getElementById("tryitout-get-observation").setAttribute("class", "btn btn--primary");

                        document.getElementById("tryitout-get-observation-json").innerText = JSON.stringify(r, null, 2);

                        document.getElementById("tryitout-get-observation-value").innerText = r.observations[0].observation;
                    })
                    .catch((e) => {
                        console.log(e);
                    });
            });

        };