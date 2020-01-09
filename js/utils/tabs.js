let tabs;
let tabPanels;
let tabList;
let activeTab;

const initialiseTabFunctionality = (container) => {
  tabs = container.querySelectorAll('[role="tab"]');
  tabPanels = container.querySelectorAll('[role="tabpanel"]');
  tabList = container.querySelector('[role=tablist]');
  activeTab = container.querySelector('[role=tab][aria-selected=true]');

  tabs.forEach((tab) => {
    tab.addEventListener('click', (e) => {
      e.preventDefault();
      setActiveTab(tab.getAttribute('aria-controls'));
    });
  });

  tabList.addEventListener('keyup', (e) => {
    let previous = [...tabs].indexOf(activeTab) - 1;
    let next = [...tabs].indexOf(activeTab) + 1;

    switch (e.keyCode) {
      case 37:
        e.preventDefault();
        previous = previous >= 0 ? previous : tabs.length - 1;
        setActiveTab(tabs[previous].getAttribute('aria-controls'));
        break;
      case 39:
        e.preventDefault();
        next = next < tabs.length ? next : 0;
        setActiveTab(tabs[next].getAttribute('aria-controls'));
        break;
    }
  });
};

const setActiveTab = (id) => {
  for (const tab of tabs) {
    if (tab.getAttribute('aria-controls') == id) {
      tab.setAttribute('aria-selected', 'true');
      tab.focus();
      activeTab = tab;
    } else {
      tab.setAttribute('aria-selected', 'false');
    }
  }
  for (const tabpanel of tabPanels) {
    if (tabpanel.getAttribute('id') == id) {
      tabpanel.setAttribute('aria-hidden', 'false');
    } else {
      tabpanel.setAttribute('aria-hidden', 'true');
    }
  }
};


export {initialiseTabFunctionality};

