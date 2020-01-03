window.feedbackOrigin = "https://beta.ons.gov.uk"

window.onload = () => {
    const toggleCodeExample = () => {
        var x = document.getElementById("codeExample");
        if (x.style.display === "none") {
            x.style.display = "block";
        } else {
            x.style.display = "none";
        }
    }
    
    const jsExample = document.getElementById('jsExampleLink')
    
    jsExample.addEventListener('click', toggleCodeExample)
    
    console.log(toggleCodeExample)
    console.log(document.getElementById('codeExample'))
}
