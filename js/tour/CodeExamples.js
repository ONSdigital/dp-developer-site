const codeBlock = document.createElement('pre');

const toggleCodeExample = (container, url) => {
  if (container.classList.contains('hidden')) {
    container.classList.remove('hidden');
    buildCodeExample(container, url);
  } else {
    container.classList.add('hidden');
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
