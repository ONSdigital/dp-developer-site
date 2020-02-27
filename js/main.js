import tourInit from './tour/index';

document.addEventListener('DOMContentLoaded', () => {
  // Check if endpoint attribute is present to then load the tour functionality
  if (document.querySelector('[data-tour-endpoint]')) {
    tourInit();
  }
});
