const orderByDate = (a, b) => {
  if (a[2] < b[2]) return -1;
  if (a[2] > b[2]) return 1;
  return 0;
};

export {orderByDate};
