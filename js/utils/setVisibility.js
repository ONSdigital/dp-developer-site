const setVisibility = (el, isVisible) => {
  if (isVisible) {
    el.classList.remove('hidden');
  } else {
    el.classList.add('hidden');
  }
};

export {setVisibility};
