// Elements
const codeContainer = document.getElementById("codeExample");
let codeBlock = document.createElement('pre')
const jsExample = document.getElementById('jsExampleLink')

const endpoint = `https://api.beta.ons.gov.uk/v1${jsExample.dataset.endpoint}`

const toggleCodeExample = () => {
    if (codeContainer.classList.contains('hidden')) {
        codeContainer.classList.remove('hidden')
        buildCodeExample(endpoint)
    } else {
        codeContainer.classList.add('hidden')
        codeBlock.innerHTML = ''
    }
}

const buildCodeExample = (url) => {
    codeBlock.innerHTML = `
        <code class="embed-code__code">
            fetch("${url}")
            .then((result) => {
                return result.json();
            })
            .then((result) => {
                console.log(result)
            })
            .catch(function(error) {
                console.log(error);
            });
        </code>
    `
    codeContainer.appendChild(codeBlock)
}

// Event listeners
jsExample.addEventListener('click', () => {
    toggleCodeExample()
})

export { toggleCodeExample, buildCodeExample }