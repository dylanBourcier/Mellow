export function formatDate(date) {
  return new Date(date)
    .toLocaleString('en-US', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      hour12: true,
    })
    .replace(',', ' ·');
}

export function formatDateShort(date) {
  return new Date(date)
    .toLocaleDateString('en-US', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    })
    .replace(/\//g, '/');
}

//I would like a function that displays the time if the date is today, otherwise displays the date in a short format
export function formatDateTime(date) {
  const today = new Date();
  const messageDate = new Date(date);

  if (
    today.getFullYear() === messageDate.getFullYear() &&
    today.getMonth() === messageDate.getMonth() &&
    today.getDate() === messageDate.getDate()
  ) {
    return messageDate.toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true,
    });
  } else {
    return formatDateShort(date);
  }
}
