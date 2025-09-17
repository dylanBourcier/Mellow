export default function previewMessage(message, numberOfChars = 100) {
  return message.length > 100
    ? message.substring(0, numberOfChars) + '...'
    : message;
}
