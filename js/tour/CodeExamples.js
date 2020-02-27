const codeBlock = document.createElement('pre');

const toggleCodeExample = (container, linkContainer, url) => {
  if (container.classList.contains('hidden')) {
    container.classList.remove('hidden');
    linkContainer.innerText = 'Hide JavaScript example';
    buildCodeExample(container, url);
  } else {
    container.classList.add('hidden');
    linkContainer.innerText = 'Show JavaScript example';
    codeBlock.innerHTML = '';
  }
};

const buildCodeExample = (container, url) => {
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
  container.appendChild(codeBlock);
};

export {toggleCodeExample, buildCodeExample};
