const codeBlock = document.createElement('pre');
const codeContainer = document.querySelector('[data-tour-example-block]');

const toggleCodeExample = (url) => {
  if (codeContainer.classList.contains('hidden')) {
    codeContainer.classList.remove('hidden');
    buildCodeExample(url);
  } else {
    codeContainer.classList.add('hidden');
    codeBlock.innerHTML = '';
  }
};

const buildCodeExample = (url) => {
  codeBlock.innerHTML =
    `<code>
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
    </code>`;
  codeContainer.appendChild(codeBlock);
};

export {toggleCodeExample, buildCodeExample};
