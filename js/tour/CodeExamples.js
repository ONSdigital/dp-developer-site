const codeBlock = document.createElement('pre');

const toggleCodeExample = (details, label, codeContainer, url) => {
  if (!codeContainer.getElementsByTagName('PRE').length) {
    buildCodeExample(codeContainer, url);
  }
  if (details.open) {
    label.innerText = 'Hide JavaScript example';
  } else {
    label.innerText = 'Show JavaScript example';
  }
};

const buildCodeExample = (container, url) => {
  codeBlock.tabIndex = 0;
  codeBlock.innerHTML =
    `<code data-tour-example-code>
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
  container.appendChild(codeBlock);
};

export { toggleCodeExample, buildCodeExample };
