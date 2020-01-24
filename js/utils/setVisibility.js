const setVisibility = (el, isVisible) => {
  if (isVisible) {
    console.log('visible')
    el.classList.remove('hidden');
  } else {
    console.log('hidden')
    el.classList.add('hidden');
  }
};

export {setVisibility};
